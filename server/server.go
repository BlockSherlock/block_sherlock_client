package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func ServerStart(port string, wTime int, rTime int, memory *Memory, apiKey string) {
	fmt.Println(port)
	router := mux.NewRouter()

	server := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("0.0.0.0:%s", port),
		WriteTimeout: time.Duration(wTime) * time.Second,
		ReadTimeout:  time.Duration(rTime) * time.Second,
	}

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(memory.authMiddleware)
	apiRouter.HandleFunc("/signTx", memory.signHandler).Methods("POST")
	apiRouter.HandleFunc("/ping", pingHandler).Methods("GET")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
