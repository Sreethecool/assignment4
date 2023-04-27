package db

import (
	"assignment4/internal/contacts"
	"fmt"
)

var contactsdb map[string]contacts.Contact

func initDB() {
	contactsdb = make(map[string]contacts.Contact)
}

func AddContact(c contacts.Contact) error {
	if contactsdb == nil {
		initDB()
	}
	if c.Id == "" {
		return fmt.Errorf("Invalid ID value")
	}

	if _, ok := contactsdb[c.Id]; ok {
		return fmt.Errorf("Contact already present")
	}

	contactsdb[c.Id] = c
	return nil
}
func GetContact(id string) ([]contacts.Contact, error) {
	var contacts []contacts.Contact
	if id == "" {
		for _, v := range contactsdb {
			contacts = append(contacts, v)
		}
	} else if len(contactsdb) != 0 {
		c, ok := contactsdb[id]
		if ok {
			contacts = append(contacts, c)
		} else {
			return nil, fmt.Errorf("Requested contact not present")
		}
	}
	if len(contacts) == 0 {
		return nil, fmt.Errorf("contacts not found")
	}
	return contacts, nil
}
func UpdateContact(newContact contacts.Contact) (contacts.Contact, error) {
	var c contacts.Contact
	var ok bool
	c, ok = contactsdb[newContact.Id]
	if !ok {
		return c, fmt.Errorf("contact ID not present")
	}
	if newContact.Email != "" {
		c.Email = newContact.Email
	}
	if newContact.Name != "" {
		c.Name = newContact.Name
	}
	if newContact.Phno != "" {
		c.Phno = newContact.Phno
	}
	contactsdb[newContact.Id] = c
	return c, nil

}
func DeleteContact(id string) error {
	if _, ok := contactsdb[id]; !ok {
		return fmt.Errorf("contact with id not found")
	}

	delete(contactsdb, id)
	return nil
}

func ClearDB() {
	for k := range contactsdb {
		delete(contactsdb, k)
	}
}
