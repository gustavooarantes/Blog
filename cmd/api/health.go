package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("ok!"))
	if err != nil {
		app.logger.Printf("Error on writing the answer: %v", err)
	}
}
