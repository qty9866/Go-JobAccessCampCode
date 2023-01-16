package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {

	pi := PersonalInformation{
		Id:     23,
		Name:   "Acsiya",
		Sex:    "男",
		Tall:   1.81,
		Weight: 81,
		Age:    24,
	}
	data, err := json.Marshal(&pi)
	if err != nil {
		log.Fatal("json marshal pi error", err)
	}
	fmt.Println(string(data))
	fmt.Println("==============")
	protoData, err := proto.Marshal(&pi)
	if err != nil {
		log.Fatal("proto marshal pi error", err)
	}
	fmt.Println(string(protoData))
	toString := base64.StdEncoding.EncodeToString(protoData)
	fmt.Println("==============")
	fmt.Println(toString)

}
