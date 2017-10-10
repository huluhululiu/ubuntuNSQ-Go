package main

import (
  "log"
  "fmt"
  "os"
  "strconv"
  "time"
  "github.com/nsqio/go-nsq"
)


func test(msg_count int,topic string){
  config := nsq.NewConfig()
  w, _ := nsq.NewProducer("18.221.119.174:4150", config)
  for i:=0 ; i<msg_count ; i++ {
    i_str := strconv.Itoa(i)
    fmt.Println("message is publishing, the content is %+v \n", "test"+i_str)
    err := w.Publish(topic, []byte("test"+i_str))
    if err != nil {
      log.Panic("Could not connect")
    }
    time.Sleep(1024 * time.Microsecond)
  }
  w.Stop()
}

func main() {
  if len(os.Args) < 4 {
    fmt.Println("Please input publisher number&number&topic as well")
  }

  pub_num, _ := strconv.Atoi(os.Args[1])
  num, _ := strconv.Atoi(os.Args[2])
  topic := os.Args[3]
  
  for i := 0 ; i< pub_num ; i++ {
    fmt.Printf("Right now is publisher %d \n",i)
    pub_str := strconv.Itoa(i)
    fmt.Printf("The topic is %+v \n",topic+pub_str)
    go test(num,topic+pub_str+"m#ephemeral") 
  }

  for i := 0 ; i < 1024 ; i++ {
    time.Sleep(20 * time.Microsecond) 
  }
}
