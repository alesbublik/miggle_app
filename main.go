package main

import (
	"fmt"
	"net/http"

	"github.com/alesbublik/miggle"
	"github.com/gorilla/mux"
)

func main() {
	res := miggle.NewResource()
	r := mux.NewRouter()
	r.HandleFunc("/aaa", res.RequestHandler(new(EndPoint)))
	http.Handle("/", r)

	http.ListenAndServe(":8000", nil)
}

type EndPoint struct{}

func (this *EndPoint) NewRequest() miggle.HttpContextHandler {
	return new(EndPoint)
}

func (this *EndPoint) Init(request *http.Request) error {
	fmt.Println("Init")
	return nil
}

func (this *EndPoint) Post(request *http.Request) (int, interface{}, http.Header) {
	return 200, "post data", nil
}

func (this *EndPoint) Get(request *http.Request) (int, interface{}, http.Header) {
	return 201, "get data", nil
}
