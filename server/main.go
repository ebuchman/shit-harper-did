package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
)

var (
	listenAddr    = "0.0.0.0:8080"
	webhookSecret = []byte{}
)

// serve files
// TODO: only html/css/js
func rootHandler(w http.ResponseWriter, r *http.Request) {
	f := r.URL.Path[1:]
	p := path.Join("site", f)
	if _, err := os.Stat(p); err != nil {
		p = "site/index.html"
	}
	http.ServeFile(w, r, p)
}

// github webhook response (confirm valid post request, git pull)
func hookHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("githook!")
	header := r.Header

	agentL := header["User-Agent"]
	if len(agentL) == 0 {
		log.Println("Empty User-Agent")
		return
	}
	agent := agentL[0]
	// assert GitHub agent
	if !strings.Contains(agent, "GitHub") {
		log.Println("git request from non Github agent")
		return
	}

	eventL := header["X-Github-Event"]
	if len(eventL) == 0 {
		log.Println("Empty X-Github-Event")
		return
	}
	event := eventL[0]
	// assert event type
	if !(strings.Contains(event, "push") || strings.Contains(event, "ping")) {
		log.Println("git request for non push or ping event")
		return
	}

	// authenticate webhook secret if exists
	if len(webhookSecret) > 0 {

		sigL := header["X-Hub-Signature"]
		if len(sigL) == 0 {
			log.Println("Empty X-Hub-Signature")
			return
		}
		sig := sigL[0]

		p, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("error reading http.req", err)
			return
		}

		// check HMAC
		if len(sig) < 6 || sig[:5] != "sha1:" {
			log.Println("signature does not begin with 'sha1:'")
		}
		sigbytes, err := hex.DecodeString(sig[5:])
		if err != nil {
			log.Println("signature is bad hex", err)
		}

		if !CheckMAC(p, sigbytes, webhookSecret) {
			log.Println("invalid hmac signature")
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
