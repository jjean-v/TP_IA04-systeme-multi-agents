package tp3web

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

var GROQ_KEY = os.Getenv("GROQ_API_KEY")

func SendRequestGet(url string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+GROQ_KEY)
	resp, errRequest := http.DefaultClient.Do(req)

	if errRequest != nil {
		fmt.Println("Error")
		fmt.Println(errRequest)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return body

}

func SendRequestPost(url string, data []byte) []byte {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+GROQ_KEY)

	resp, errRequest := http.DefaultClient.Do(req)

	if errRequest != nil {
		fmt.Println("Error")
		fmt.Println(errRequest)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return body

}
