package runner

import (
  "time"
  "phul-poc/definition"
  "phul-poc/builder"
)

// Init the runner
func Init(ch chan definition.Message) {
  pl := builder.List()
  go runLoop (pl, ch)
}

// runLoop polls for updates in initial version
func runLoop(pl []definition.Provider, ch chan definition.Message) {
  for true {
    for _, p := range pl {
      res := p.GetMessages()
      for _, r := range res {
        ch <- r
      }
      time.Sleep(1*time.Second)
    }
    time.Sleep(10*time.Second)
  }
}
