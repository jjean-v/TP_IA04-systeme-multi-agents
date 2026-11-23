// Agent LLM minimal : boucle agentique + tool calling, sans framework.
// Fonctionne avec toute API compatible OpenAI (Groq, Ollama, ...).
//
//	export LLM_URL=https://api.groq.com/openai/v1/chat/completions
//	export LLM_KEY=gsk_...          # inutile pour Ollama
//	export LLM_MODEL=...            # nom du modèle chez le fournisseur
//	go run . "Combien de lignes fait main.go ? Multiplie par 7."
//
// Ollama : LLM_URL=http://localhost:11434/v1/chat/completions
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// ---------- Types de l'API (sous-ensemble du format OpenAI) ----------

type Message struct {
	Role       string     `json:"role"` // system, user, assistant, tool
	Content    string     `json:"content"`
	ToolCalls  []ToolCall `json:"tool_calls,omitempty"`
	ToolCallID string     `json:"tool_call_id,omitempty"`
}

type ToolCall struct {
	ID       string `json:"id"`
	Type     string `json:"type"` // "function"
	Function struct {
		Name      string `json:"name"`
		Arguments string `json:"arguments"` // JSON encodé dans une chaîne !
	} `json:"function"`
}

type ToolDef struct {
	Type     string `json:"type"` // "function"
	Function struct {
		Name        string         `json:"name"`
		Description string         `json:"description"`
		Parameters  map[string]any `json:"parameters"` // JSON Schema
	} `json:"function"`
}

type ChatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Tools       []ToolDef `json:"tools,omitempty"`
	Temperature float64   `json:"temperature"`
}

type ChatResponse struct {
	Choices []struct {
		Message      Message `json:"message"`
		FinishReason string  `json:"finish_reason"` // "stop", "tool_calls", ...
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
	} `json:"usage"`
}

// ---------- Appel REST (stateless) ----------

func chat(ctx context.Context, req ChatRequest) (Message, error) {
	// log.Println(os.Getenv("LLM_URL"))
	body, _ := json.Marshal(req)
	httpReq, err := http.NewRequestWithContext(ctx, "POST", os.Getenv("LLM_URL"), bytes.NewReader(body))
	if err != nil {
		return Message{}, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	// log.Println(os.Getenv("LLM_KEY"))
	if key := os.Getenv("GROQ_API_KEY"); key != "" {
		httpReq.Header.Set("Authorization", "Bearer "+key)
	}

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return Message{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		msg, _ := io.ReadAll(io.LimitReader(resp.Body, 2048))
		return Message{}, fmt.Errorf("HTTP %d : %s", resp.StatusCode, msg)
	}

	var cr ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&cr); err != nil {
		return Message{}, err
	}
	if len(cr.Choices) == 0 {
		return Message{}, errors.New("réponse vide")
	}
	return cr.Choices[0].Message, nil
}

// ---------- Outils ----------

type Tool struct {
	Def ToolDef
	Run func(args map[string]any) (string, error)
}

func newTool(name, desc string, params map[string]any, run func(map[string]any) (string, error)) Tool {
	var d ToolDef
	d.Type = "function"
	d.Function.Name = name
	d.Function.Description = desc
	d.Function.Parameters = params
	return Tool{Def: d, Run: run}
}

var tools = map[string]Tool{
	"count_lines": newTool(
		"count_lines",
		"Compte le nombre de lignes d'un fichier du répertoire courant.",
		map[string]any{
			"type": "object",
			"properties": map[string]any{
				"path": map[string]any{"type": "string", "description": "chemin du fichier"},
			},
			"required": []string{"path"},
		},
		func(args map[string]any) (string, error) {
			path, _ := args["path"].(string)
			if strings.Contains(path, "..") || strings.HasPrefix(path, "/") {
				return "", errors.New("accès refusé") // moindre privilège
			}
			data, err := os.ReadFile(path)
			if err != nil {
				return "", err
			}
			return strconv.Itoa(strings.Count(string(data), "\n")), nil
		}),
	"multiply": newTool(
		"multiply",
		"Multiplie deux nombres.",
		map[string]any{
			"type": "object",
			"properties": map[string]any{
				"a": map[string]any{"type": "number"},
				"b": map[string]any{"type": "number"},
			},
			"required": []string{"a", "b"},
		},
		func(args map[string]any) (string, error) {
			a, _ := args["a"].(float64)
			b, _ := args["b"].(float64)
			return strconv.FormatFloat(a*b, 'f', -1, 64), nil
		}),
}

func execute(tc ToolCall) string {
	t, ok := tools[tc.Function.Name]
	if !ok {
		return "erreur : outil inconnu " + tc.Function.Name
	}
	var args map[string]any
	if err := json.Unmarshal([]byte(tc.Function.Arguments), &args); err != nil {
		return "erreur : arguments JSON invalides"
	}
	out, err := t.Run(args)
	if err != nil {
		return "erreur : " + err.Error() // l'erreur est renvoyée au modèle, pas fatale
	}
	return out
}

// ---------- Boucle agentique (le « harnais ») ----------

const maxSteps = 10

func runAgent(ctx context.Context, task string) (string, error) {
	var defs []ToolDef
	for _, t := range tools {
		defs = append(defs, t.Def)
	}

	history := []Message{
		{Role: "system", Content: "Tu es un agent. Utilise les outils si nécessaire, puis réponds brièvement."},
		{Role: "user", Content: task},
	}

	for step := 0; step < maxSteps; step++ {
		msg, err := chat(ctx, ChatRequest{
			Model:    os.Getenv("LLM_MODEL"),
			Messages: history,
			Tools:    defs,
		})
		if err != nil {
			return "", err
		}
		history = append(history, msg)

		if len(msg.ToolCalls) == 0 { // pas d'action demandée : réponse finale
			return msg.Content, nil
		}
		for _, tc := range msg.ToolCalls {
			res := execute(tc)
			fmt.Printf("  [outil] %s(%s) -> %s\n", tc.Function.Name, tc.Function.Arguments, res)
			history = append(history, Message{Role: "tool", ToolCallID: tc.ID, Content: res})
		}
	}
	return "", errors.New("nombre maximal d'étapes atteint")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage : go run . \"tâche\"")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	answer, err := runAgent(ctx, strings.Join(os.Args[1:], " "))
	if err != nil {
		fmt.Println("échec :", err)
		os.Exit(1)
	}
	fmt.Println(answer)
}
