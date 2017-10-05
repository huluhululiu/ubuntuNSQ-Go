package main,mq

import (
	"encoding/binary"
	//"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)
type Nsq struct {
	pub       *nsq.Producer
	msgCount  int
	msgSize   int
	topic     string
	topic_raw string
}

func NewNsq(msgCount int, msgSize int, topic_raw string) *Nsq {
	//topic := "0#ephemeral"
	topic := topic_raw
	topic += "m#ephemeral"

	pub, _ := nsq.NewProducer("18.221.119.174:4150", nsq.NewConfig())
	return &Nsq{
		pub:       pub,
		msgCount:  msgCount,
		msgSize:   msgSize,
		topic:     topic,
		topic_raw: topic_raw,
	}
}

func (n *Nsq) Teardown() {
	n.pub.Stop()
}

func (n *Nsq) Send(message []byte) {
	
	message = append(message, n.topic_raw...)
	message = append(message, "\n"...)
	b := make([]byte, n.msgSize-len(message))
	message = append(message, b...)
	n.pub.PublishAsync(n.topic, message, nil)
}
func newTest(msgCount, msgSize int, topic string) {

	nsq := mq.NewNsq(msgCount, msgSize, topic)

	start := time.Now().UnixNano()
	b := make([]byte, 24)
	id := make([]byte, 5)
	//b:=[]byte{}
	//time.Sleep(5000*time.Millisecond)
	for i := 0; i < msgCount; i++ {
		if i == 1 {
			time.Sleep(5 * time.Second)
		}
		binary.PutVarint(b, time.Now().UnixNano())
		binary.PutVarint(id, int64(i))
		//b=append(b, strconv.FormatInt(int64(i), 10)...)
		copy(b[19:23], id[:])

		nsq.Send(b)
		time.Sleep(4096 * time.Microsecond)
	}

	stop := time.Now().UnixNano()
	ms := float32(stop-start) / 1000000
	log.Printf("Sent %d messages in %f ms\n", msgCount, ms)
	log.Printf("Sent %f per second\n", 1000*float32(msgCount)/ms)

	nsq.Teardown()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	num, _ := strconv.Atoi(os.Args[1])
	topic, _ := strconv.Atoi(os.Args[2])
	for i := 0; i < num; i++ {
		go newTest(13000, 512, strconv.Itoa(topic+i))
	}
	for {
		time.Sleep(20 * time.Second)
	}

}
