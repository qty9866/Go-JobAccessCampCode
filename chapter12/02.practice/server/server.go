package main

import (
	"Learning-JobAccess-Camp/chapter12/02.practice/frinterface"
	"Learning-JobAccess-Camp/pkg/apis"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	var rankServer frinterface.ServeInterface = NewFatRateRank()

	mux := http.NewServeMux()
	mux.Handle("/register", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// 不区分大小写的判断是否是post
		if !strings.EqualFold(request.Method, "post") {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		// 如果不给内容
		if request.Body == nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		defer request.Body.Close()

		// 从body中读取内容
		payload, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("无法解析数据:%s", err)))
			return
		}
		var pi *apis.PersonalInformation
		err = json.Unmarshal(payload, &pi)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("Json Mashall数据失败:%s", err)))
			return
		}

		if err = rankServer.RegisterPersonalInformation(pi); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("注册失败:%s", err)))
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Register succeed"))
	}))
	mux.Handle("/personalinfo", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.EqualFold(request.Method, "post") {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		// 如果不给内容
		if request.Body == nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		defer request.Body.Close()

		// 从body中读取内容
		payload, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("无法解析数据:%s", err)))
			return
		}
		var pi *apis.PersonalInformation
		err = json.Unmarshal(payload, &pi)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(fmt.Sprintf("Json Mashall数据失败:%s", err)))
			return
		}

		if fr, err := rankServer.UpdatePersonalInformation(pi); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("更新失败:%s", err)))
			return
		} else {
			writer.WriteHeader(http.StatusOK)
			data, _ := json.Marshal(fr)
			writer.Write([]byte("update information succeed~"))
			writer.Write(data)
		}
	}))
	mux.Handle("/rank", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//todo call rank function
		if !strings.EqualFold(request.Method, "get") {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		name := request.URL.Query().Get("name")
		if name == "" {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("name参数不能为空"))
			return
		}
		if fr, err := rankServer.GetFatRate(name); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("获取排行数据失败:%s", err)))
			return
		} else {
			writer.WriteHeader(http.StatusOK)
			data, _ := json.Marshal(fr)
			writer.Write([]byte("GET rank succeed~"))
			writer.Write(data)
		}
	}))
	mux.Handle("/ranktop", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.EqualFold(request.Method, "get") {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if frTop, err := rankServer.GetTop(); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(fmt.Sprintf("获取排行数据失败:%s", err)))
			return
		} else {
			writer.WriteHeader(http.StatusOK)
			data, _ := json.Marshal(frTop)
			writer.Write([]byte("GET rank succeed~"))
			writer.Write(data)
		}
	}))

	http.ListenAndServe(":8080", mux)

}
