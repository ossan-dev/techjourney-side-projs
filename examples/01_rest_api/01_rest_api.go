package main

import (
	"fmt"
	"net/http"
)

// curl http://localhost:8089/example			=> GET
// curl http://localhost:8089/health			=> GET
// curl -X POST http://localhost:8089/health	=> POST

func HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HandleHealthCheck triggered...")
	if r.Method != http.MethodGet {
		w.Write([]byte(fmt.Sprintln("health check must be a GET request")))
		return
	}
	w.Write([]byte(fmt.Sprintln("the system is health")))
}

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/example route was hit...")
		_ = r
		w.Write([]byte(fmt.Sprintln("this is a demo endpoint")))
	})
	r.HandleFunc("/health", HandleHealthCheck)
	if err := http.ListenAndServe(":8089", r); err != nil {
		panic(err)
	}
}
