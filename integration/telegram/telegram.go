package telegram

import (
  "github.com/go-resty/resty"
  "phul-poc/definition"
  "fmt"
  "encoding/json"
)

const name = "TELEGRAM"
const botkey =  "insert-bot-key-here"

type Instance struct {
  ApiKey string
}

func (tg *Instance) GetName() string {
  return name
}

func (tg *Instance) GetUpdates() (definition.Response, error) {
  client := resty.New()

  resp, err := client.R().EnableTrace().Get("https://api.telegram.org/" + botkey + "/getUpdates")

  fmt.Println("resp: ", resp)
  fmt.Println("err: ", err)
  return resp, err
}

func (tg *Instance) GetMessages() ([]definition.Message) {
  updates, _ := tg.GetUpdates()

  if updates == nil {
    return nil
  }

  var m []TelResult
  results := updates.Result()
  b, _ := json.Marshal(results)
  json.Unmarshal(b, &m)
  for res := range m {
    fmt.Println("Message: ", res)
  }
  return nil
}

func (tg *Instance) SendMessage(msg string) (definition.Response, error){
  return nil, nil
}
