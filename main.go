package main
import (
    "fmt"
    "net/http"
)
type httpHandler struct{
	message string
}

func (h httpHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request){
	fmt.Fprint(resp, h.message)
}

func main() {
	handle1 := httpHandler{message: "index"}
	handle2 := httpHandler{message: "Reset"}

	http.Handle("/", handle1)
	http.Handle("/reset", handle2)

	fmt.Println("Server is starting from http://localhost:8181...")
    	http.ListenAndServe(":8181", nil)
}
