package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"tp3web"
)

const MODEL_IA_GPT_OSS_20B = "openai/gpt-oss-20b"
const MODEL_IA_GPT_OSS_SAFEGUARD_20B = "openai/gpt-oss-safeguard-20b"
const MODEL_IA_ALIBABA_QWEN_3_8_27B = "qwen/qwen3.8-27b"
const MODEL_IA_SDAIA_ALLAM_2_7B = "allam-2-7b"

func main() {

	for {
		fmt.Print("Question: ")
		reader := bufio.NewReader(os.Stdin)
		question, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		request := tp3web.PrepareRequestWithContext(question, MODEL_IA_GPT_OSS_20B)
		result := tp3web.ChatWitModels(request, false)

		for _, element := range result.Choices {
			fmt.Println(element.Message.Content)
			reasoning := tp3web.Reasoning(element.Message.Reasoning, MODEL_IA_GPT_OSS_SAFEGUARD_20B)
			//tp3web.StoreConversation("reasoning: " + reasoning.Model)
			for _, reason := range reasoning.Choices {
				fmt.Println("\n\n Reasoning:" + string(reason.Message.Content) + "\n\n")
				tp3web.StoreConversation("reasoning: " + reason.Message.Content)
			}

		}
	}
}

//fmt.Println(generateContext())
