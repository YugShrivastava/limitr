package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/time/rate"
)

const (
	eventsAllowed = 10
	tokensAllowed = 4
)

func rateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.Handler{
	limiter := rate.NewLimiter(eventsAllowed, tokensAllowed)
	
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			message := Message{
				Status: "Request Failed",
				Body: "The API is at capacity, try again later.",
			}
			
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&message)
			return
		} else {
			next(w, r)
		}
	})
}