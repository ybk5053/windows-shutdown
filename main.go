package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) > 1 {
		log.Fatal("Error")
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3030", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	user := r.URL.Query().Get("user")
	pass := r.URL.Query().Get("pass")
	exec.Command("net", "rpc", "shutdown", "-f", "-t", "0", "-C", "shutting down", "-I", ip, "-U", fmt.Sprintf("%s%%%s", user, pass)).Run()
	log.Println("shutdown")
}
