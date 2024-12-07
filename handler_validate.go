package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	clean := getCleanedBody(params.Body)

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: clean,
	})
}

func getCleanedBody(body string) string {
	cussWords := []string{"kerfuffle", "sharbert", "fornax"}
	words := strings.Split(body, " ")

	for i, s := range words {
		if slices.Contains(cussWords, strings.ToLower(s)) {
			words[i] = "****"
		}
	}

	return strings.Join(words, " ")
}
