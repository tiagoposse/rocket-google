package controller

import (
	"context"
	"fmt"
	"log"
	"os"

	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
)

type GoogleController struct {
	service *admin.Service
}

func NewGoogleController(ctx context.Context) (*GoogleController, error) {
	serviceAccountFile := os.Getenv("GOOGLE_SERVICE_ACCOUNT_FILE")
	// adminUser := os.Getenv("GOOGLE_ADMIN_USER")

	// config, err := google.JWTConfigFromJSON(serviceAccountJSON,
	// 	admin.AdminDirectoryUserScope, admin.AdminDirectoryGroupScope,
	// )
	// config.Subject = adminUser

	srv, err := admin.NewService(ctx, option.WithScopes(admin.AdminDirectoryGroupMemberScope), option.WithCredentialsFile(serviceAccountFile))
	if err != nil {
		log.Fatal(err)
	}
	return &GoogleController{
		service: srv,
	}, nil
}

// Function to remove a Person from a Group
func (c *GoogleController) RemoveUserFromGroup(person, group string) error {
	// Remove user from the group
	err := c.service.Members.Delete(group, person).Do()
	if err != nil {
		return fmt.Errorf("failed to remove user from group: %v", err)
	}

	return nil
}

// Function to add a Person to a Group
func (c *GoogleController) AddPersonToGroup(person, group string) error {
	// Add user to the group
	_, err := c.service.Members.Insert(group, &admin.Member{
		Email: person,
	}).Do()
	if err != nil {
		return fmt.Errorf("failed to add user to group: %v", err)
	}
	return nil
}
