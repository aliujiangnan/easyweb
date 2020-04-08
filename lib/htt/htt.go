package htt

import (
    "net/http"
	"encoding/json"
)

type HttpCtx struct{
	Writer http.ResponseWriter
	Req *http.Request
	UserId int
	GameId int
}

func NewCtx(w http.ResponseWriter, r *http.Request) HttpCtx{
	return HttpCtx{w, r, 0, 0}
}

func (this *HttpCtx)Send(d interface{}) error{
	this.Writer.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(this.Writer)
	return encoder.Encode(d)
}