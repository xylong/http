package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type ProxyHandler struct{}

func (ProxyHandler) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	defer func() {
		if err:=recover();err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
	}()

	if r.URL.Path =="/a" {
		req, err:=http.NewRequest(r.Method, "http://127.0.0.1:9091",r.Body)
		if err!=nil {
			log.Println(err)
		}

		rep,err:= http.DefaultClient.Do(req)
		if err!=nil {
			log.Println(err)
		}
		defer rep.Body.Close()


		bytes, err:=ioutil.ReadAll(rep.Body)
		if err!=nil {
			log.Println(err)
		}
		w.Write(bytes)

		return
	}

	w.Write([]byte("default"))
}


func main() {
	http.ListenAndServe(":8080", ProxyHandler{})
}
