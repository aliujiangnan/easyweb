package router

import (
	"io"
    "net/http"
	"encoding/json"
	"easyweb/lib/util"
	"easyweb/lib/htt"

)

func Get(path string, callbacks ...func(params map[string]interface{}, ctx htt.HttpCtx)int){
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		var params = make(map[string]interface{})
		for k,v := range r.Form{
			params[k] = v
		}
		handleReq(params, htt.NewCtx(w, r), callbacks)
	})
}

func Post(path string, callbacks ...func(params map[string]interface{}, ctx htt.HttpCtx)int){
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		e := w.Header().Get("up-encode")
		if e == "1" {
			r.Body = util.DecryptAES(r.Body).(io.ReadCloser)
		}
		d := json.NewDecoder(r.Body)
		var params map[string]interface{}
		d.Decode(&params)

		handleReq(params, htt.NewCtx(w, r), callbacks)
	})
}

func handleReq(params map[string]interface{}, ctx htt.HttpCtx, callbacks []func(params map[string]interface{}, ctx htt.HttpCtx)int){
	for _, cb := range callbacks {
		if cb != nil && cb(params, ctx) != 0 {
			break
		}
	}
}