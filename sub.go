package main

import (
  "log"
  "sync"
  "time"
  "os"
  "runtime"
  "strconv"
  "github.com/bitly/go-nsq"
)
func test(msg_count int,topic string){
  config := nsq.NewConfig()
  w, _ := nsq.NewProducer("18.221.119.174:4150", config)
  for i:=0 ; i<msg_count ; i++ {
    i_str := strconv.Itoa(i)
    fmt.Println("message is publishing", i)
    err := w.Publish(topic, []byte("test"+i_str))
    if err != nil {
      log.Panic("Could not connect")
    }
    time.Sleep(1024 * time.Microsecond)
  }
  w.Stop()
}

func main() {

  wg := &sync.WaitGroup{}
  wg.Add(1)
  topic := os.Args[1]
  config := nsq.NewConfig()
  q, _ := nsq.NewConsumer(topic, "ch", config)
  q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
      log.Printf("Got a message: %v", message)
      wg.Done()
      return nil
  }))
  err := q.ConnectToNSQD("18.221.119.174:4150")
  if err != nil {
      log.Panic("Could not connect")
  }
  wg.Wait()

}
