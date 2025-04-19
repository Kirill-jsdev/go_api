package main

import "net/http"

func (app *application) enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Controll-Allow-Origin", "http://*")
		
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Controll-Allow-Credentials", "true")
			w.Header().Set("Access-Controll-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
			w.Header().Set("Access-Controll-Allow-Headers", "Accept, Content-Type, X-CFRS-Token, Authorization")
			return
		} else {
			h.ServeHTTP(w, r)
		}
	})
}