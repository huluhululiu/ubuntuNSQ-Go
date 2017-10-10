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

func test(numofmsg int, msg_size int, Channel string)
{
  //assign topic&channel, config consumer
  channel := Channel + "m#ephemeral"
  topic := channel
  config := nsq.NewConfig()
  config.MaxInFlight = 1000000
  config.OutputBufferSize = -1
  sub, _ := nsq.NewConsumer(topic, channel, config)
  
  //design handler's handle function
  sub.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
      HandleMsg(message.Body)
      return nil
  }))
  
  //connect
  err := sub.ConnectToNSQD("18.221.119.174:4150")
  if err != nil {
      log.Panic("Could not connect")
  }
   
}

func main() {
  wg := &sync.WaitGroup{}
  wg.Add(1)
  topic := os.Args[1]
  
  wg.Wait()

}
