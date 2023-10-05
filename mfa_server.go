package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"sync"
)

var (
	// In-memory store of valid MFA tokens
	tokens = make(map[string]bool)
	// Mutex to ensure concurrent safety
	mfaMu sync.Mutex
)

// Generate a random MFA token
func generateToken() string {
	buf := make([]byte, 6) // 48 bits for a 6-byte token
	rand.Read(buf)
	return hex.EncodeToString(buf)
}

// Issue a new MFA token
func issueTokenHandler(w http.ResponseWriter, r *http.Request) {
	mfaMu.Lock()
	defer mfaMu.Unlock()

	newToken := generateToken()
	tokens[newToken] = true
	fmt.Fprintf(w, "Issued MFA token: %s\n", newToken)
}

// Validate an MFA token
func validateTokenHandler(w http.ResponseWriter, r *http.Request) {
	mfaMu.Lock()
	defer mfaMu.Unlock()

	token := r.URL.Query().Get("token")
	if _, exists := tokens[token]; exists {
		delete(tokens, token) // Remove the token since it's one-time-use
		fmt.Fprint(w, "Token is valid\n")
		return
	}
	fmt.Fprint(w, "Token is invalid\n")
}

func main() {
	http.HandleFunc("/issue-token", issueTokenHandler)
	http.HandleFunc("/validate-token", validateTokenHandler)

	fmt.Println("MFA Server running on :8081")
	http.ListenAndServe(":8081", nil)
}
