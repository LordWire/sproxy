package main

import (
	"encoding/json"
	"net/http"
)

// Simple healthcheck function. It's not performing any advanced
// application health checking, but it still depends on
// a correct http Server object being up and running,
// which by definition makes it a valid health check for
// the scope of this exercise.

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status OK"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
