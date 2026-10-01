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
		tp3web.ChatWitModels(question, MODEL_IA_GPT_OSS_20B)
	}

	//fmt.Println(generateContext())

}
