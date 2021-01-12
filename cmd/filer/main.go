package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"olymp.alabor.me/dev/git/swissmanu/filer/pkg/conf"
	"olymp.alabor.me/dev/git/swissmanu/filer/pkg/rule"
)

type byModifiedDate []os.FileInfo

func (a byModifiedDate) Len() int           { return len(a) }
func (a byModifiedDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byModifiedDate) Less(i, j int) bool { return a[i].ModTime().Before(a[j].ModTime()) }

type applyRequest struct {
	RuleName string
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func main() {
	config, err := conf.NewDefaultConfig()
	if err != nil {
		log.Fatal(err)
	}
	rules, err := rule.ReadRules(*config)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	router.HandleFunc("/inbox", func(w http.ResponseWriter, r *http.Request) {
		files, err := ioutil.ReadDir(config.InboxPath)
		if err != nil {
			http.Error(w, "Could not read from inbox", http.StatusInternalServerError)
			log.Print(err)
			return
		}

		sort.Sort(sort.Reverse(byModifiedDate(files)))

		items := make([]string, 0)
		for _, f := range files {
			name := f.Name()
			if strings.HasSuffix(strings.ToLower(name), ".pdf") {
				items = append(items, name)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	}).Methods("GET")

	router.HandleFunc("/inbox/{item}/apply", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		var request applyRequest

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "Could not parse ApplyRequest", http.StatusBadRequest)
			log.Print(err)
			return
		}

		sourcePath := filepath.Join(config.InboxPath, vars["item"])
		if !fileExists(sourcePath) {
			http.Error(w, "", http.StatusNotFound)
			log.Print(sourcePath + " not found")
			return
		}

		ruleToApply, err := rule.FindRule(rules.Rules, request.RuleName)
		if err != nil {
			http.Error(w, "Rule "+request.RuleName+" not found", http.StatusInternalServerError)
			log.Print("Rule " + request.RuleName + " not found")
			log.Print(err)
			return
		}

		log.Print("Apply " + ruleToApply.Name + " to " + sourcePath)
		err = rule.ApplyRule(ruleToApply, sourcePath, config)
		if err != nil {
			http.Error(w, "Could not apply "+request.RuleName+" to "+sourcePath, http.StatusInternalServerError)
			log.Print(err)
			return
		}
	}).Methods("POST")

	router.HandleFunc("/inbox/{item}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		itemPath := filepath.Join(config.InboxPath, vars["item"])

		if !fileExists(itemPath) {
			http.Error(w, "", http.StatusNotFound)
			log.Print(itemPath + " not found")
			return
		}

		file, err := ioutil.ReadFile(itemPath)
		if err != nil {
			http.Error(w, "Could not read file", http.StatusInternalServerError)
			log.Print(err)
			return
		}

		w.Write(file)
	}).Methods("GET")

	router.HandleFunc("/rules", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rules)
	}).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir(config.UIPath)))

	log.Print("Start HTTP server on " + config.Addr)
	log.Print("Inbox Path: " + config.InboxPath)
	log.Print("Data Path: " + config.DataPath)
	log.Print("Rules Path: " + config.RulesPath)

	srv := &http.Server{
		// Handler: handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router),
		Handler: router,
		Addr:    config.Addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
