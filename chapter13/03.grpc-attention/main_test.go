package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestDecode(t *testing.T) {
	jsonData := ` {"id":23,"name":"Acsiya","sex":"男","tall":1.81,"weight":81,"age":24}`
	protoDataBase64 := `CBcSBkFjc2l5YRoD55S3JRSu5z8tAACiQjAY`

	pi1 := &PersonalInformation{}
	json.Unmarshal([]byte(jsonData), pi1)
	fmt.Println("解析json")
	fmt.Printf("%+v\n", *pi1)

	pi2 := &PersonalInformation{}
	protoData, _ := base64.StdEncoding.DecodeString(protoDataBase64)
	proto.Unmarshal(protoData, pi2)
	fmt.Println("解析protobuf")
	fmt.Printf("%+v", *pi2)
}
