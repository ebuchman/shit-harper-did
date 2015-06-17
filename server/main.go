package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

var (
	listenAddr    = "0.0.0.0:8080"
	webhookSecret = []byte{}
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "site/index.html")
}

// github webhook response (confirm valid post request, git pull)
func hookHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("githook!")
	header := r.Header
	agent := header["User-Agent"][0]
	event := header["X-Github-Event"][0]
	sig := header["X-Hub-Signature"][0]
	// assert GitHub agent
	if !strings.Contains(agent, "GitHub") {
		log.Println("git request from non Github agent")
		return
	}
	// assert event type
	if !(strings.Contains(event, "push") || strings.Contains(event, "ping")) {
		log.Println("git request for non push or ping event")
		return
	}

	if len(webhookSecret) > 0 {
		// check HMAC
		p := make([]byte, r.ContentLength)
		sum := 0
		// read http req - there is almost certainly a oneline for this...
		for sum < int(r.ContentLength) {
			n, err := r.Body.Read(p[sum:])
			if err != nil {
				log.Println("error reading http.req", err)
				return
			}
			sum += n
		}
		sigbytes, err := hex.DecodeString(sig[5:]) // sig begins with "sha1:"
		if err != nil {
			log.Println("no hex to bytes!", err)
		}

		if !CheckMAC(p, sigbytes, webhookSecret) {
			log.Println("git request with invalid signature")
			return
		}
	}

	// all checks passed
	GitPull()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/hook", hookHandler)

	fmt.Println("Serving shit-harper-did on ", listenAddr)
	fmt.Println(http.ListenAndServe(listenAddr, mux))
}

// CheckMAC returns true if messageMAC is a valid HMAC tag for message.
func CheckMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha1.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	log.Println(hex.EncodeToString(expectedMAC), hex.EncodeToString(messageMAC))
	return hmac.Equal(messageMAC, expectedMAC)
}

func GitPull() {
	cmd := exec.Command("git", "fetch", "-a", "origin")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	log.Println(out.String())

	cmd = exec.Command("git", "reset", "--hard", "origin/master")
	cmd.Stdout = &out
	cmd.Run()
	log.Println(out.String())
}
