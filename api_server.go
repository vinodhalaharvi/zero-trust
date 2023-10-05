package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	dynamicIdentityServerURL = "http://localhost:8080/validate?identity="
	mfaServerURL             = "http://localhost:8081/validate-token?token="
)

// Validate Dynamic Identity
func validateDynamicIdentity(identity string) bool {
	resp, err := http.Get(dynamicIdentityServerURL + identity)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return strings.Contains(string(body), "Identity is valid")
}

// Validate MFA Token
func validateMFAToken(token string) bool {
	resp, err := http.Get(mfaServerURL + token)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	defer resp.Body.Close()

	return true
}

// Handler for fetching current time
func timeHandler(w http.ResponseWriter, r *http.Request) {
	identity := r.URL.Query().Get("identity")
	mfaToken := r.URL.Query().Get("token")

	if !validateDynamicIdentity(identity) || !validateMFAToken(mfaToken) {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized\n")
		return
	}

	currentTime := time.Now().Format(time.RFC1123)
	fmt.Fprintf(w, "Current local time: %s\n", currentTime)
}

func main() {
	http.HandleFunc("/current-time", timeHandler)

	fmt.Println("API Server running on :8082")
	http.ListenAndServe(":8082", nil)
}
