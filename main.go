package main

import (
	"net/http"
	"strconv"

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

type Entity struct{}

type EndPoint struct {
	EntityId int64
}

func (this *EndPoint) NewRequest() miggle.HttpContextHandler {
	return new(EndPoint)
}

func (this *EndPoint) Init(request *http.Request) error {
	vars := mux.Vars(request)

	if id, ok := vars["id"]; ok {
		eid, _ := strconv.ParseInt(id, 10, 64)
		this.EntityId = eid
	}
	return nil
}

func (this *EndPoint) Get(request *http.Request) (int, interface{}, http.Header) {
	return 200, "get data", nil
}

func (this *EndPoint) Post(request *http.Request) (int, interface{}, http.Header) {
	return 201, "post data", nil
}

func (this *EndPoint) GetEntities() ([]*Entity, error) {
	var entities []*Entity
	return entities, nil
}

func (this *EndPoint) GetEntity(id int64) (*Entity, error) {
	entity := new(Entity)
	return entity, nil
}
