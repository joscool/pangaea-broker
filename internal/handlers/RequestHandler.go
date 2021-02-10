package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/custom-broker/internal/data"
	"github.com/custom-broker/internal/interfaces"
	"github.com/custom-broker/internal/services"
	"github.com/gorilla/mux"
)

var pubSubStore interfaces.PubSub

func init() {
	pubSubStore = services.NewCustomPubsub()
}

func HandleSubsription(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Subscribe")
	vars := mux.Vars(r)
	topic := vars["topic"]

	subscriptionRequest := data.SubscriptionRequest{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&subscriptionRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	pubSubStore.Subscribe(topic, subscriptionRequest.Url)

	respondWithJSON(w, http.StatusOK, subscriptionRequest)
}

func HandlePublishing(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Publish")
	vars := mux.Vars(r)
	topic := vars["topic"]

	publishingRequest := data.PublishingRequest{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&publishingRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	pubSubStore.Publish(topic, publishingRequest.Message)

	respondWithJSON(w, http.StatusOK, publishingRequest)
}

func HandleEevent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Event")

	var postData map[string]string
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&postData); err != nil {
		log.Printf("Invalid request payload : %v", err)
		return
	}

	log.Printf("received : %v", postData)

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
