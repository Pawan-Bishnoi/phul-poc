package definition

type Message interface {

}

// Response iface that every provider's response should adhere to
type Response interface {
  Body() []byte
  Status() string
  StatusCode() int
  Result() interface{}
}

// Provider iface - all providers should be of this form
type Provider interface {
  GetName() (string)
  GetUpdates() (Response, error)
  GetMessages() ([]Message)
  SendMessage(string) (Response, error)
}
