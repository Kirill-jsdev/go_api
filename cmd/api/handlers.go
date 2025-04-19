package main

import (
	"backend/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
        Status  string `json:"status"`
        Message string `json:"message"`
        Version string `json:"version"`
    }{
        Status:  "success",
        Message: "Welcome to the API",
        Version: "1.0.0",
    }

	//convert to json
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	releaseDate, _ := time.Parse("2006-01-02", "1986-03-07")
	highlander := models.Movie {
		ID: 1,
		Title: "Highlander",
		ReleaseDate: releaseDate,
		MPAARating: "R",
		Runtime: 116,
		Description: "Nice movie",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	movies = append(movies, highlander)

	releaseDate, _ = time.Parse("2006-01-02", "1986-03-07")
	batman := models.Movie {
		ID: 2,
		Title: "Batman",
		ReleaseDate: releaseDate,
		MPAARating: "R",
		Runtime: 116,
		Description: "Nice movie too",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	movies = append(movies, batman)

	//convert to json
	jsonData, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}