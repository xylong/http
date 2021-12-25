package main

import (
	"net/http"
	"os"
)

type web1handler struct {

}

func (web1handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("web1"))
}

type web2handler struct {

}

func (web2handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("web2"))
}


func main() {
	c:=make(chan os.Signal)

	go func() {
		http.ListenAndServe(":9091", web1handler{})
	}()

	go func() {
		http.ListenAndServe(":9092", web2handler{})
	}()

	<-c
}
