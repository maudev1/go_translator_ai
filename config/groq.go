package config

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type Message struct {
	Role    string
	Content string
}

var url = "https://api.groq.com/openai/v1/chat/completions"
var llModel = "llama-3.1-8b-instant"

func Chat(content string, groqToken string) string {

	payload := strings.NewReader("{\n\t\"model\": \"" + llModel + "\",\n\t\"messages\": [\n\t\t{\n\t\t\t\"role\": \"system\",\n\t\t\t\"content\":\"Traduza tudo que estiver entre <<  >>, em portugues, retorne somente a tradução sem explicações.\"\n\t\t\t\n\t\t},\n\t\t{\n\t\t\t\"role\": \"user\",\n\t\t\t\"content\": \"<<" + content + ">>\"  \n\t\t}\n\t]\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "insomnia/11.4.0")
	req.Header.Add("Authorization", "Bearer "+groqToken)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)

}
