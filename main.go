package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type DB struct {
	counter int
}
type CounterOperation struct {
	OperationType string
}

var DBInst = DB{counter: 0}

func main() {
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
func enableCors(responseWriter *http.ResponseWriter) {
	(*responseWriter).Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	(*responseWriter).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*responseWriter).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*responseWriter).Header().Set("Access-Control-Allow-Credentials", "true")
}
func counter(responseWriter http.ResponseWriter, request *http.Request) {
	//	_, err := fmt.Fprintf(w, "hello world")
	enableCors(&responseWriter)
	if request.Method == "GET" {

		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusOK)
		respData := fmt.Sprintf("{\"counter\":\"%d\"}", DBInst.counter)
		responseWriter.Write([]byte(respData))

	} else if request.Method == "POST" {

		b, _ := ioutil.ReadAll(request.Body)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(request.Body)

		var counterOp CounterOperation
		_ = json.Unmarshal(b, &counterOp)

		if counterOp.OperationType == "Increment" {
			DBInst.counter = DBInst.counter + 1

		} else if counterOp.OperationType == "Decrement" {
			DBInst.counter = DBInst.counter - 1
		}
		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.WriteHeader(http.StatusOK)
		respData := fmt.Sprintf("{\"counter\":\"%d\"}", DBInst.counter)
		responseWriter.Write([]byte(respData))

	}
	//w.Write([]byte("Success"))

}
