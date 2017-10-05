package main

import (
  "log"
  "fmt"
  "os"
  "strconv"
  "time"
  "github.com/nsqio/go-nsq"
)


func test(msg_count,topic string){
  config := nsq.NewConfig()
  w, _ := nsq.NewProducer("18.221.119.174:4150", config)
  for i:=0 ; i<msg_count ; i++ {
    i_str, _ := strconv.Itoa(i)
    err := w.Publish(topic+i_str, []byte("test"+i_str))
    if err != nil {
      log.Panic("Could not connect")
    }
    time.Sleep(1024 * time.Microsecond)
  }
}

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Please input number as well")
  }

  num, _ := strconv.Atoi(os.Args[1])
  
  go test(num,"test")

  w.Stop()
}
