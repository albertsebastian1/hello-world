package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMethod(c *gin.Context) {
	fmt.Println("\napi.go 'GetMethod' called")
	message := "GetMethod called"
	c.JSON(http.StatusOK, message)
}

func helloworld() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})
	fmt.Printf("Server running (port=80), route: http://localhost\n")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}

func main() {
	helloworld()
}
