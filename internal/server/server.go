package server

import (
	"fmt"
	"net/http"
)


func Start(port string)  {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "Hi! Server run!!!")
	})

	fmt.Printf("Server start on http://localhost%s\n", port)
	fmt.Println("Stop --> push Ctrl+C")

	err:= http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Printf("Error run server: %v\n", err)
	}
}