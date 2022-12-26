package main

import (
	"fmt"
	"net/http"
)

func main() {

	//http.ListenAndServe(":8088", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
	//	qp := request.URL.Query()
	//	data, _ := json.Marshal(qp)
	//	writer.Write([]byte(`This is hello from Hud` + string(data)))
	//}))

	// 读取body写body
	/*http.ListenAndServe(":8080", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Body == nil {
			writer.Write([]byte("empty body"))
			return
		}
		//data, _ := ioutil.ReadAll(request.Body)
		//defer request.Body.Close()
		writer.Write([]byte("index"))
	}))*/

	mux := http.NewServeMux()
	mux.Handle("/hello", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello"))
	}))
	mux.Handle("/rank", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("rank"))
	}))
	mux.Handle("/history/Hud", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//todo call rank function
		writer.Write([]byte("Hud's History"))
	}))
	// ....

	mux.Handle("/history", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		qp := request.URL.Query()
		name := qp.Get("name")
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(fmt.Sprintf("%s这里的一切只关于%s", request.Method, name)))
	}))

	http.ListenAndServe(":8080", mux)
}
