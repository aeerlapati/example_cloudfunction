package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func addFun(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var n1 = vars["num1"]
	var n2 = vars["num2"]
	var operation = "add"
	var j = util(w, n1, n2, operation)
	fmt.Fprintf(w, j)

}

func subFun(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var n1 = vars["num1"]
	var n2 = vars["num2"]
	var operation = "sub"
	var j = util(w, n1, n2, operation)
	fmt.Fprintf(w, j)
}

func mulFun(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var n1 = vars["num1"]
	var n2 = vars["num2"]
	var operation = "mul"
	var j = util(w, n1, n2, operation)
	fmt.Fprintf(w, j)
}

func divFun(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var n1 = vars["num1"]
	var n2 = vars["num2"]
	var operation = "div"
	var j = util(w, n1, n2, operation)
	fmt.Fprintf(w, j)
}

func util(w http.ResponseWriter, n1 string, n2 string, operation string) string {
	i, err := strconv.Atoi(n1)
	j, err2 := strconv.Atoi(n2)
	if err == nil && err2 == nil {
		switch {
		case operation == "add":
			j = i + j
		case operation == "div":
			j = i / j
		case operation == "sub":
			j = i - j
		case operation == "mul":
			j = i * j
		}
		return strconv.Itoa(j)
	}
	fmt.Println(err)
	return err.Error()
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/add/{num1}/{num2}", addFun)
	router.HandleFunc("/mul/{num1}/{num2}", mulFun)
	router.HandleFunc("/div/{num1}/{num2}", divFun)
	router.HandleFunc("/sub/{num1}/{num2}", subFun)
	log.Fatal(http.ListenAndServe(":9099", router))
}
