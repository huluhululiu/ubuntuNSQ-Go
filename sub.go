package main

import (
  "log"
<<<<<<< HEAD
  "sync"
  "time"
  "os"
  "runtime"
  "strconv"
  "github.com/bitly/go-nsq"
=======
  "fmt"
  "time"
  "os"
  "strconv"
  "github.com/nsqio/go-nsq"
>>>>>>> 5c5e0bf7332fdf64e0c13611308b6c40637413f7
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

<<<<<<< HEAD
func main() {

  wg := &sync.WaitGroup{}
  wg.Add(1)
  topic := os.Args[1]
  config := nsq.NewConfig()
  q, _ := nsq.NewConsumer(topic, "ch", config)
  q.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
      log.Printf("Got a message: %v", message)
      wg.Done()
=======
//for further implementation
func HandleMsg(message []byte) {
  //only to print out
  fmt.Printf("The message is %+v \n",message)
}

func test(Channel string) {
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
>>>>>>> 5c5e0bf7332fdf64e0c13611308b6c40637413f7
      return nil
  }))
  
  //connect
  err := sub.ConnectToNSQD("18.221.119.174:4150")
  if err != nil {
      log.Panic("Could not connect")
  }
}

func main() {
  //usage go run sub.go <number of topic to subscribe> <topic raw name>
  if len(os.Args) < 3 {
    fmt.Println("Please input souscriber_number&ctopic(channel) as well")
  }
  
  sub_num, _ := strconv.Atoi(os.Args[1])
  topic := os.Args[2]
  
  //subscribe to each topic
  for i := 0 ; i< sub_num ; i++ {
    fmt.Printf("Right now is sublisher %d \n",i)
    sub_str := strconv.Itoa(i)
    fmt.Printf("The topic is %+v \n",topic+sub_str)
    go test(topic+sub_str) 
  }
  
  for {
    time.Sleep(20 * time.Second)
  }
}
