package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func validateArgs(args []string) {
	if len(args) != 2 {
		fmt.Println("Wrong number of args! Use like <port> <message>")
		os.Exit(1)
	}
}

func getPortAndMessage(args []string) (port int, message string) {
	validateArgs(args)
	port, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Failed to convert port arg to int, exiting...")
		panic(err)
	}
	message = args[1]
	return
}

func makeHandlerWithMessage(message string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(message))
	}

}

func startServer(port int, message string) {
	fmt.Println("Listening on port:", port)
	fmt.Println("Responding with message:", message)
	http.HandleFunc("/", makeHandlerWithMessage(message))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
func main() {
	args := os.Args[1:]
	port, message := getPortAndMessage(args)
	startServer(port, message)
}
