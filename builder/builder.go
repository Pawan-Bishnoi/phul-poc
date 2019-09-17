package builder

import (
	"phul-poc/definition"
	"phul-poc/integration/telegram"
	"sync"
)

var once sync.Once
var telegramInstance *telegram.Instance

// List - returns list of provider instances
func List() []definition.Provider {
	pl := make([]definition.Provider, 0)
	once.Do(func() {
		telegramInstance = &telegram.Instance{ApiKey: "key"}
	})

	return append(pl, telegramInstance)
}
