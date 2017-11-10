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

<<<<<<< HEAD
func main() {
  config := nsq.NewConfig()
  w, _ := nsq.NewProducer("18.221.119.174:4150", config)
<<<<<<< HEAD
  num, _ := strconv.Atoi(os.Args[1])
  for i:=0 ; i<num;i++{
  err := w.Publish("write_test", []byte("test"))

  if err != nil {
=======
=======
func test(msg_count int,topic string){
  //create publisher for the topic
  config := nsq.NewConfig()
  w, _ := nsq.NewProducer("18.221.119.174:4150", config)
  
  //publish n messages 
>>>>>>> 5c5e0bf7332fdf64e0c13611308b6c40637413f7
  for i:=0 ; i<msg_count ; i++ {
    i_str := strconv.Itoa(i)
    fmt.Println("message is publishing, the content is %+v \n", "test"+i_str)
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
  //usage go run pub.go <numof publisher> <numof messages for each publisher> <topic raw name>
  if len(os.Args) < 4 {
    fmt.Println("Please input publisher number&number&topic as well")
  }

  pub_num, _ := strconv.Atoi(os.Args[1])
  num, _ := strconv.Atoi(os.Args[2])
  topic := os.Args[3]
  
<<<<<<< HEAD
  go test(num,"write_test")
>>>>>>> d07900374c46567cf578c83398f79d8c83e5c850
=======
  //create many publisher 
  for i := 0 ; i< pub_num ; i++ {
    fmt.Printf("Right now is publisher %d \n",i)
    pub_str := strconv.Itoa(i)
    fmt.Printf("The topic is %+v \n",topic+pub_str)
    go test(num,topic+pub_str+"m#ephemeral") 
  }

  for i := 0 ; i < 1024 ; i++ {
    time.Sleep(20 * time.Microsecond) 
  }
>>>>>>> 5c5e0bf7332fdf64e0c13611308b6c40637413f7
}
