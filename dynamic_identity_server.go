package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"sync"
)

var (
	// In-memory store of valid dynamic identities
	// Using a map for O(1) lookups and deletions
	identities = make(map[string]bool)
	// Mutex to ensure concurrent safety
	mu sync.Mutex
)

// Generate a random identity
func generateIdentity() string {
	buf := make([]byte, 16) // 128 bits
	rand.Read(buf)
	return hex.EncodeToString(buf)
}

// Issue a new dynamic identity
func issueIdentityHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	newIdentity := generateIdentity()
	identities[newIdentity] = true
	fmt.Fprintf(w, "Issued identity: %s\n", newIdentity)
}

func validateIdentity(identity string) bool {
	_, exists := identities[identity]
	return exists
}

// Revoke a given dynamic identity
func revokeIdentityHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	identityToRevoke := r.URL.Query().Get("identity")
	if _, exists := identities[identityToRevoke]; !exists {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Identity not found\n")
		return
	}

	delete(identities, identityToRevoke)
	fmt.Fprint(w, "Identity revoked\n")
}

func main() {
	http.HandleFunc("/issue", issueIdentityHandler)
	http.HandleFunc("/revoke", revokeIdentityHandler)
	http.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
		identity := r.URL.Query().Get("identity")
		if validateIdentity(identity) {
			fmt.Fprint(w, "Identity is valid\n")
		} else {
			fmt.Fprint(w, "Identity is invalid\n")
		}
	})

	fmt.Println("Dynamic Identity Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
