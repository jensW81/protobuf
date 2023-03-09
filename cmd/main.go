package main

import (
	"golang.source-fellows.com/seminar/protobuf-hello-word"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

func main() {
	hello := &protobuf_hello_world.HelloWorld{Message: "Hello World"}
	b, err := proto.Marshal(hello)
	if err != nil {
		log.Fatalf("could not encode hell %v", err)
	}
	f, err := os.OpenFile("db.blob", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("could not open db %v", err)
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		log.Fatalf("could not write to %v", err)
	}

}
