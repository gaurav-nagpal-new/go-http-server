package main

import (
	"fmt"
	"log"
	"net/http"
)

type myInterface interface {
	testMethod(int)
}

type myFunc func(int)

// practical example

// here when t() or tn() method is wrapped around myFunc, below myFunc is replaced with that function,
// thats why we are able to call it.
func (m myFunc) testMethod(x int) {
	m(x)
}

// a simple http web server which responds hello world
func main() {

	// need to wrap in HandlerFunc as my function do not implement ServiceHttp method but HandlerFunc does !
	http.Handle("/", http.HandlerFunc(HelloWorld))

	log.Print("server started")

	// when you convert your function into another same function signature, if wrapper function
	// have a method, our function will also have that method and we can call it.
	//
	myFn(5, myFunc(t))
	myFn(5, myFunc(tn))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("error", err)
	}

}

func t(x int) {
	fmt.Println(x)
}

func tn(x int) {
	fmt.Println("new thing ", x)
}

func myFn(x int, in myInterface) {
	in.testMethod(x)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi")
	response := "Hello World"
	w.Header().Set("sample", "sample")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(response))
}
