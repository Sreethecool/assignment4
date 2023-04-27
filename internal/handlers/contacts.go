package handlers

import (
	"assignment4/internal/contacts"
	"assignment4/internal/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

func GetContacts(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	contactlist, err := db.GetContact(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sort.Slice(contactlist, func(i, j int) bool {
		return contactlist[i].Id < contactlist[j].Id
	})
	json.NewEncoder(w).Encode(contactlist)

}

func SetContact(w http.ResponseWriter, r *http.Request) {
	var c contacts.Contact
	body, bodyErr := io.ReadAll(r.Body)
	if bodyErr != nil {
		fmt.Println("error reading body. got error", bodyErr)
	}

	err := json.Unmarshal(body, &c)
	if err != nil {
		fmt.Printf("error parsing body, getting error %v, body:%v", err, string(body))
		http.Error(w, "Unable to parse request body", http.StatusBadRequest)
		return
	}
	err = db.AddContact(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(c)
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id not passed", http.StatusBadRequest)
		return
	}
	var contact contacts.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, "Unable to parse request body", http.StatusBadRequest)
		return
	}
	contact.Id = id
	newContact, err := db.UpdateContact(contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(newContact)
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id not passed", http.StatusBadRequest)
		return
	}
	err := db.DeleteContact(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "contact deleted successfully")
}
