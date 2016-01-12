package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func SimplePanicRecover(v int, w http.ResponseWriter) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("SimplePanic info is: ", err)
			w.WriteHeader(597)
		}
	}()

	ss := []string{"a", "b"}
	fmt.Println(ss[v])
}
func serverHandle(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Println("value", req.Form["value"][0])
	i, _ := strconv.Atoi(req.Form["value"][0])
	SimplePanicRecover(i, w)
}
func main() {
	http.HandleFunc("/", serverHandle)
	http.ListenAndServe(":8001", nil)
}
