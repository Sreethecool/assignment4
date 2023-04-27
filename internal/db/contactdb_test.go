package db

import (
	"assignment4/internal/contacts"
	"reflect"
	"sort"
	"testing"
)

var contact1 = contacts.Contact{
	Id:    "test1",
	Name:  "Test1",
	Email: "test@addcontact",
	Phno:  "123456",
}
var contact2 = contacts.Contact{
	Id:    "test2",
	Name:  "Test2",
	Email: "test2@addcontact",
	Phno:  "1234567",
}
var invalidContact = contacts.Contact{
	Id:    "",
	Name:  "Test2",
	Email: "test2@addcontact",
	Phno:  "1234567",
}
var updatedcontact1 = contacts.Contact{
	Id:    "test1",
	Name:  "UpdatedTest1",
	Email: "Updatedtest1@addcontact",
	Phno:  "0123456",
}
var partialupdatedcontact2 = contacts.Contact{
	Id:    "test2",
	Name:  "",
	Email: "",
	Phno:  "",
}

func InsertContacts() {
	AddContact(contact1)
	AddContact(contact2)
}

func TestAddContact(t *testing.T) {
	type args struct {
		c contacts.Contact
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"adding contact 1",
			args{
				c: contact1,
			},
			false,
		},
		{"adding contact 2",
			args{
				c: contact2,
			},
			false,
		},
		{"adding invalid contact",
			args{
				c: invalidContact,
			},
			true,
		},
		{"adding existing contact",
			args{
				c: contact1,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddContact(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("AddContact() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetContact(t *testing.T) {
	type args struct {
		id string
	}
	//test without anycontact
	ClearDB()
	got, err := GetContact("test1")
	if err == nil {
		t.Errorf("Expected contact not found error, got %v", got)
	}

	tests := []struct {
		name    string
		args    args
		want    []contacts.Contact
		wantErr bool
	}{
		{
			"Get one contact",
			args{
				id: "test1",
			},
			[]contacts.Contact{contact1},
			false,
		},
		{
			"Get all contact",
			args{
				id: "",
			},
			[]contacts.Contact{contact1, contact2},
			false,
		},
		{
			"Get invalid contact",
			args{
				id: "invalid contact",
			},
			nil,
			true,
		},
	}
	InsertContacts()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetContact(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) > 1 {
				sort.Slice(got, func(p, q int) bool {
					return got[p].Id < got[q].Id
				})
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetContact() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateContact(t *testing.T) {
	type args struct {
		newContact contacts.Contact
	}
	ClearDB()
	InsertContacts()
	tests := []struct {
		name    string
		args    args
		want    contacts.Contact
		wantErr bool
	}{
		{
			"test update",
			args{
				newContact: updatedcontact1,
			},
			updatedcontact1,
			false,
		},
		{
			"test partial update",
			args{
				newContact: partialupdatedcontact2,
			},
			contact2,
			false,
		},
		{
			"test partial update",
			args{
				newContact: invalidContact,
			},
			contacts.Contact{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateContact(tt.args.newContact)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateContact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateContact() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteContact(t *testing.T) {
	type args struct {
		id string
	}
	ClearDB()
	InsertContacts()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"delete contact",
			args{
				id: "test1",
			},
			false,
		},
		{
			"delete again contact",
			args{
				id: "test1",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteContact(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteContact() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
