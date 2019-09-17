package telegram

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	"phul-poc/definition"
)

const name = "TELEGRAM"
const botkey = "insert-bot-key-here"

type Instance struct {
	ApiKey string
}

func (tg *Instance) GetName() string {
	return name
}

func (tg *Instance) GetUpdates() (definition.Response, error) {
	client := resty.New()

	resp, err := client.R().EnableTrace().Get("https://api.telegram.org/" + botkey + "/getUpdates")
	return resp, err
}

func (tg *Instance) GetMessages() []definition.Message {
	arr := make([]definition.Message, 0)
	updates, _ := tg.GetUpdates()
	if updates == nil {
		return nil
	}

	var r TelResponse
	json.Unmarshal(updates.Body(), &r)
	for _, res := range r.Result {
		arr = append(arr, res.Message)
	}
	fmt.Println(arr)
	return arr
}

func (tg *Instance) SendMessage(msg string) (definition.Response, error) {
	return nil, nil
}
