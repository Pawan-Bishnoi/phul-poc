package telegram

/*
Sample response
{
  "ok":true,
  "result":
  [{
    "update_id":109399140,
    "message":
    {
      "message_id":8,
      "from":
      {
        "id":388648082,
        "is_bot":false,
        "first_name":"Pawan",
        "last_name":"Bishnoi",
        "language_code":"en"
      },
      "chat":
      {
        "id":388648082,
        "first_name":"Pawan",
        "last_name":"Bishnoi",
        "type":"private"
      },
      "date":1568651093,
      "text":"ting"
    }
   }]
}
*/

type TelResponse struct {
  Ok bool              `json:"ok"`
  Result []TelResult   `json:"result"`
}

// Telegram result is the response of GetUpdates call
type TelResult struct {
  UpdateId string      `json:"update_id"`
  Message TelMessage   `json:"message"`
}

// Telegram user struct as per docs
type TelUser struct {
  Id int               `json:"id"`
  IsBot bool           `json:"is_bot"`
  FirstName string     `json:"first_name "`
  Username string      `json:"username"`
}

// Telegram chat struct as per docs
type TelChat struct {
  Id int               `json:"id"`
  FirstName string     `json:"first_name"`
  LastName string      `json:"last_name"`
  Type string          `json:"type"`
}

// Telegram message struct as per docs
type TelMessage struct {
  Message_id int    `json:"message_id"`
  From TelUser      `json:"from"`
  Chat TelChat      `json:"chat"`
  Date int          `json:"date"`
  Text string       `json:"text"`
  Entities []TelEntity    `json:"entities"`
}

// Telegram entity struct as per docs
type TelEntity struct {
  Offset int     `json:"offset"`
  Length int     `json:"length"`
  Type string    `json:"type"`
}

