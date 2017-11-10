package main

import (
  "log"
  "fmt"
  "os"
  "strconv"
  "encoding/json" 
  //"github.com/nsqio/go-nsq"
)
type Mine struct{
    name string //`json:"a"`
    number float64
}

func main(){
	if len(os.Args) < 1 {
  fmt.Println("Please input topic &message as well")
  }
    s := Mine{}
    v := os.Args[1]//`{"a":"1"}`
    json.Unmarshal([]byte(v), &s)
    fmt.Println(s)
}