package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nikkiDEEE/go-blockchain/pkg/controllers"
	"github.com/nikkiDEEE/go-blockchain/pkg/models"
)

var bc *models.Blockchain

func main() {
	bc = controllers.NewBlockchain()

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.getBlockchain).Methods("GET")
	r.HandleFunc("/", controllers.writeBlock).Methods("POST")
	r.HandleFunc("/new", controllers.newProduct).Methods("POST")

	log.Println("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
