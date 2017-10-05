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
<<<<<<< HEAD
  num, _ := strconv.Atoi(os.Args[1])
  for i:=0 ; i<num;i++{
  err := w.Publish("write_test", []byte("test"))

  if err != nil {
=======
  for i:=0 ; i<msg_count ; i++ {
    i_str = strconv.Itoa(i)
    err := w.Publish(topic, []byte("test"+i_str))
    if err != nil {
>>>>>>> d07900374c46567cf578c83398f79d8c83e5c850
      log.Panic("Could not connect")
  }
<<<<<<< HEAD
  }

  w.Stop()
=======
  w.Stop()
}

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Please input number as well")
  }

  num, _ := strconv.Atoi(os.Args[1])
  
  go test(num,"write_test")
>>>>>>> d07900374c46567cf578c83398f79d8c83e5c850
}
