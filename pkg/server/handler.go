package server

import (
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/buabaj/fairness-engine/pkg/mpc"
)

type secretRequest struct {
	Secret    string `json:"secret"`
	Threshold int    `json:"threshold"`
	NumShares int    `json:"numShares"`
}

type shareRequest struct {
	Share string `json:"share"`
}

// generates shares for a given secret.
func handleKeyGeneration(w http.ResponseWriter, r *http.Request) {
	var req secretRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	secret := new(big.Int)
	secret, ok := secret.SetString(req.Secret, 10)
	if !ok {
		http.Error(w, "invalid secret", http.StatusBadRequest)
		return
	}

	sss, err := mpc.NewShamirSecretSharing(secret, req.Threshold, req.NumShares)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	prime = sss.Prime
	shares = sss.Shares
	threshold = req.Threshold

	sharesStr := make([]string, len(sss.Shares))
	for i, share := range sss.Shares {
		sharesStr[i] = share.Text(10)
	}

	resp, err := json.Marshal(sharesStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

var (
	prime     *big.Int
	shares    []*big.Int
	threshold int
)

func handleShareSubmission(w http.ResponseWriter, r *http.Request) {
	var req shareRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	share := new(big.Int)
	share, ok := share.SetString(req.Share, 10)
	if !ok {
		http.Error(w, "invalid share", http.StatusBadRequest)
		return
	}

	shares = append(shares, share)

	w.WriteHeader(http.StatusOK)
}

func handleComputation(w http.ResponseWriter, r *http.Request) {
	if len(shares) < threshold {
		http.Error(w, "insufficient number of shares", http.StatusBadRequest)
		return
	}

	sss := &mpc.ShamirSecretSharing{
		Prime:     prime,
		Shares:    shares,
		Threshold: threshold,
	}

	secret, err := sss.Combine(shares)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(secret.Text(10))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
