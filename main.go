package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Message string `json:"message"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Moo")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, World!"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}
	response := Response{Message: fmt.Sprintf("Hello, %s!", name)}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func mooHandler(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")
	if strings.HasPrefix(userAgent, "curl") {
		cow := `
					 ______________
					< Hello, curl! >
					 --------------
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣾⡟⠷⢦⡄⠀⠀⠀⠀/⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⡶⢶⡦⣄⡀⠀⠀⠀⢀⣤⠞⠻⠛⠃⠀⠈⠛⢦⡀⠀/⠀⠀⠀⠀⢀⣀⣀⡀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⣀⡀⠀⠀⠀⠀⠀⠀⠀⢸⠀⠀⠉⠒⠬⢍⡉⠉⣹⠃⠀⠀⠀⠀⠀⠀⠀⠀⢹⠒⢶⣶⡾⠛⠛⠁⠀⠙
⠀⠀⠀⠀⠀⠀⠀⢀⣠⠾⠋⠉⠉⠉⠛⠲⠤⢤⣀⣀⣀⡘⣧⡀⠀⠀⠀⠀⠈⣹⠁⠀⠀⠀⠀⠀⠀⠀⢀⠀⠸⡦⠊⠁⠀⠀⠀⠀⢄⡼
⠀⠀⠀⢀⣀⡤⠶⠿⠅⠀⢀⠀⠔⠊⠀⠀⠀⠀⠀⠀⠈⠉⠙⠛⠮⣄⣀⣠⣼⣧⡄⠀⠀⠀⠆⠀⠀⣤⣼⠲⣄⢀⡀⢀⣀⣤⣤⠴⠋⠀
⠀⠀⢀⡟⠁⠀⠀⠀⠀⠀⠀⠀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢘⣧⠀⠀⠀⠀⠀⠀⠉⠀⢰⠃⠈⠉⠉⠉⠁⠀⠀⠀⠀
⠀⠀⣸⠁⠀⠀⡰⠁⠀⠀⠀⠊⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡴⠀⠀⠀⠀⠀⣸⠃⠀⠀⠀⠀⠀⠁⠀⠀⣼⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⢰⣿⠀⠀⣰⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠞⠀⠀⠀⠀⠀⣰⡃⠀⠀⠀⠀⠀⠀⠀⢀⣴⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⢠⠇⣿⠀⢠⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠎⠀⠀⠀⠀⠀⠀⢿⢻⡍⢠⡄⠀⠀⢠⣠⠎⣼⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⡸⢰⡏⠀⢸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡏⠀⠀⠀⠀⠀⠀⠀⠈⠻⣗⣚⣿⣴⣺⠟⢁⡼⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⣇⣸⡇⠀⠸⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⠀⠀⠀⢀⡔⠀⠀⠀⠀⠀⠀⠉⠉⠉⠀⢀⡞⠁⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⣇⢸⡇⠀⠀⢿⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⠇⠀⠀⠀⠏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡠⠊⠀⠀⣷⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⢹⣸⡇⠀⠀⠘⣷⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⠞⠁⠀⠀⠀⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠈⡏⢷⠀⠀⠀⠘⣷⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠃⠀⠀⠀⠀⣼⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠘⣼⡀⠀⠀⠀⣼⢷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠈⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢻⡇⠀⠀⢠⢿⠀⠉⠛⢷⠶⠤⢤⣄⣀⣀⣀⣹⠰⠀⠀⠀⠀⣠⠄⠀⠀⠀⠀⠀⣀⡀⠀⢠⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢸⠁⠀⠀⡏⢸⠀⠀⠀⡿⠀⠀⠀⠀⠀⠀⠈⠙⡇⠀⠀⠀⢰⣇⠀⠀⠀⢀⡠⢾⠉⠀⠀⣸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢺⠀⠀⢸⠇⢸⠀⠀⣼⠃⠀⠀⠀⠀⠀⠀⠀⠀⢻⡀⠀⠀⠀⡏⠓⠦⠴⠋⠁⢸⠀⠀⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢸⡄⠀⡾⠀⡞⠀⢰⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣇⠀⠀⠀⡇⠀⠀⠀⠀⠀⣼⠀⠀⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠘⠃⢀⡇⢰⠁⠀⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠀⠀⢰⠃⠀⠀⠀⠀⠀⢻⠀⠀⢰⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢀⡀⠈⣇⡸⠀⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠀⠀⢸⠀⠀⠀⠀⠀⠀⢸⡀⠀⡜⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⣼⠀⠀⢸⡇⠀⠠⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⠀⠀⢸⠀⠀⠀⠀⠀⠀⢸⠀⠀⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⢰⡟⠀⠀⠚⡧⣄⣀⣸⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⡆⠀⣼⠀⠀⠀⠀⠀⠀⣸⠀⠀⣇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠻⠤⠤⠖⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⠀⢹⡀⠀⠀⠀⠀⠀⢹⡀⠀⠸⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⠃⠀⠈⣇⠀⠀⠀⠀⠀⠀⣧⡀⠐⠻⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡞⠒⠀⠐⠻⡄⠀⠀⠀⠀⠀⠙⠒⠒⠊⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠳⣤⣠⠤⠴⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
        `
		w.Write([]byte(cow))
	} else {
		response := Response{Message: "This endpoint is for curl users only!"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/greet", greetHandler)
	              http.HandleFunc("/moo", mooHandler)
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
