package speech_to_text

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func SpeechToText(apiToken string, filePath string) string {
	apiUrl := "https://stt.api.cloud.yandex.net/speech/v1/stt:recognize?topic=general&lang=ru-RU"

	var file *os.File
	file, _ = os.Open(filePath)

	req, _ := http.NewRequest("POST", apiUrl, file)
	req.Header.Set("Authorization", "Api-Key "+apiToken)
	res, _ := http.DefaultClient.Do(req)
	body, _ := io.ReadAll(res.Body)
	res.Body.Close()

	type Body struct {
		Result string `json:"result"`
	}
	var message Body
	json.Unmarshal(body, &message)

	return message.Result
}
