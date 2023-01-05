package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type frClient struct {
	handring Interface
}

func (f frClient) register() {
	pi, _ := f.handring.ReadPersonalInformation()
	data, _ := json.Marshal(pi)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:8080/register", "application/json", r)
	if err != nil {
		log.Println("WARNING: Register failed!", err)
		return
	}
	if resp.Body != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("返回:", string(data))
	}
}
func (f frClient) update() {
	pi, _ := f.handring.ReadPersonalInformation()
	data, _ := json.Marshal(pi)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://localhost:8080/personalinfo", "application/json", r)
	if err != nil {
		log.Println("WARNING: Register failed!", err)
		return
	}
	if resp.Body != nil {
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("返回:", string(data))
	}
}

func main() {
	frCli := &frClient{handring: &fakeInterface{
		name:       fmt.Sprintf("Hud:%d", time.Now().UnixNano()),
		sex:        "男",
		baseWeight: 81.0,
		baseTall:   1.81,
		baseAge:    24,
	}}
	frCli.register()
	for {
		frCli.update()
		time.Sleep(1 * time.Second)
	}
}
