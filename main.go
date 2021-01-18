package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-rest/structures"
	"io/ioutil"
	"log"
	"net/http"
)

// Mocked data
var dogs []structures.Dog

func main() {
	// Router strict by slash
	router := mux.NewRouter().StrictSlash(true)

	// GET
	router.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(dogs)
	})

	// POST
	router.HandleFunc("/dog", func (w http.ResponseWriter, r *http.Request) {
		var newDog structures.Dog
		fmt.Println("here 1")
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "Request's body cannot be null")
		}

		json.Unmarshal(reqBody, &newDog)
		dogs = append(dogs, newDog)
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(newDog)
	})

	// GET
	router.HandleFunc("/dog/{id}", func (w http.ResponseWriter, r *http.Request) {
		registerId := mux.Vars(r)["id"]

		for _, registeredDog := range dogs {
			if registeredDog.ID == registerId {
				json.NewEncoder(w).Encode(registeredDog)
			}
		}
	})

	// PATCH
	router.HandleFunc("/dog/{id}", func (w http.ResponseWriter, r *http.Request) {
		eventID := mux.Vars(r)["id"]
		fmt.Println("here")
		var updatedEvent structures.Dog

		reqBody, err := ioutil.ReadAll(r.Body)

		if err != nil {
			fmt.Fprintf(w, "Request's body cannot be null")
		}
		json.Unmarshal(reqBody, &updatedEvent)

		for i, singleEvent := range dogs {
			if singleEvent.ID == eventID {
				singleEvent.Name = updatedEvent.Name
				singleEvent.Description = updatedEvent.Description
				dogs = append(dogs[:i], singleEvent)
				json.NewEncoder(w).Encode(singleEvent)
			}
		}
	})

	// DELETE
	router.HandleFunc("/dog/{id}", func (w http.ResponseWriter, r *http.Request) {
		eventID := mux.Vars(r)["id"]

		for i, singleEvent := range dogs {
			if singleEvent.ID == eventID {
				dogs = append(dogs[:i], dogs[i+1:]...)
				fmt.Fprintf(w, "Dog Id %v has been successfully deleted", eventID)
			}
		}
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}