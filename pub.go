package main

import (
  "log"
  "github.com/nsqio/go-nsq"
)

func IsValidTopicName(name string) bool {
	return isValidName(name)
}

func IsValidChannelName(name string) bool {
	return isValidName(name)
}

func main() {
  config := nsq.NewConfig()
  w, _ := nsq.NewProducer("18.221.119.174:4150", config)
  num, _ := strconv.Atoi(os.Args[1])
  for i:=0 ; i<num;i++{
  err := w.Publish("write_test", []byte("test"))

  if err != nil {
      log.Panic("Could not connect")
  }
  }

  w.Stop()
}
