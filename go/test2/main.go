package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Contact struct {
	Name        string
	PhoneNumber string
	Email       string
}

type ContactsManager struct {
	contacts []Contact
}

func NewContactsManager() *ContactsManager {
	return &ContactsManager{
		contacts: []Contact{},
	}
}

func (cm *ContactsManager) AddContact(name, phoneNumber, email string) {
	contact := Contact{
		Name:        name,
		PhoneNumber: phoneNumber,
		Email:       email,
	}
	cm.contacts = append(cm.contacts, contact)
	fmt.Println("Contact added successfully!")
}

func (cm *ContactsManager) ViewAllContacts() {
	if len(cm.contacts) == 0 {
		fmt.Println("No contacts found.")
		return
	}

	fmt.Println("\n===== All Contacts =====")
	for i, contact := range cm.contacts {
		fmt.Printf("[%d] Name: %s, Phone: %s, Email: %s\n", i+1, contact.Name, contact.PhoneNumber, contact.Email)
	}
	fmt.Println("=======================")
}

func (cm *ContactsManager) UpdateContact(index int, name, phoneNumber, email string) bool {
	if index < 0 || index >= len(cm.contacts) {
		return false
	}

	cm.contacts[index].Name = name
	cm.contacts[index].PhoneNumber = phoneNumber
	cm.contacts[index].Email = email
	return true
}

func (cm *ContactsManager) DeleteContact(index int) bool {
	if index < 0 || index >= len(cm.contacts) {
		return false
	}

	cm.contacts = append(cm.contacts[:index], cm.contacts[index+1:]...)
	return true
}

func getInput(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	cm := NewContactsManager()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n===== Contacts Management System =====")
		fmt.Println("1. Add Contact")
		fmt.Println("2. View All Contacts")
		fmt.Println("3. Update Contact")
		fmt.Println("4. Delete Contact")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice (1-5): ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			name := getInput(scanner, "Enter name: ")
			phoneNumber := getInput(scanner, "Enter phone number: ")
			email := getInput(scanner, "Enter email: ")
			cm.AddContact(name, phoneNumber, email)

		case "2":
			cm.ViewAllContacts()

		case "3":
			cm.ViewAllContacts()
			if len(cm.contacts) == 0 {
				continue
			}

			indexStr := getInput(scanner, "Enter the index of the contact to update: ")
			index, err := strconv.Atoi(indexStr)
			if err != nil || index < 1 || index > len(cm.contacts) {
				fmt.Println("Invalid index!")
				continue
			}

			name := getInput(scanner, "Enter new name: ")
			phoneNumber := getInput(scanner, "Enter new phone number: ")
			email := getInput(scanner, "Enter new email: ")

			if cm.UpdateContact(index-1, name, phoneNumber, email) {
				fmt.Println("Contact updated successfully!")
			} else {
				fmt.Println("Failed to update contact.")
			}

		case "4":
			cm.ViewAllContacts()
			if len(cm.contacts) == 0 {
				continue
			}

			indexStr := getInput(scanner, "Enter the index of the contact to delete: ")
			index, err := strconv.Atoi(indexStr)
			if err != nil || index < 1 || index > len(cm.contacts) {
				fmt.Println("Invalid index!")
				continue
			}

			confirmStr := getInput(scanner, fmt.Sprintf("Are you sure you want to delete %s? (y/n): ", cm.contacts[index-1].Name))
			confirm := strings.ToLower(confirmStr) == "y" || strings.ToLower(confirmStr) == "yes"

			if confirm {
				if cm.DeleteContact(index - 1) {
					fmt.Println("Contact deleted successfully!")
				} else {
					fmt.Println("Failed to delete contact.")
				}
			}

		case "5":
			fmt.Println("Exiting the Contacts Management System. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}
}
