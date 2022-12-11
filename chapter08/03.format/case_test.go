package main

import (
	"Learning-JobAccess-Camp/pkg/apis"
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestMarshal(t *testing.T) {
	ps := &apis.PersonalInformation{
		Name:   "Hud",
		Sex:    "男",
		Tall:   181,
		Weight: 82,
		Age:    24,
	}
	fmt.Printf("%+v\n", ps)
	data, err := json.Marshal(ps)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	var pi apis.PersonalInformation
	err = json.Unmarshal(data, &pi)
	if err != nil {
		return
	}
	fmt.Printf("%+v\n", pi)
}
