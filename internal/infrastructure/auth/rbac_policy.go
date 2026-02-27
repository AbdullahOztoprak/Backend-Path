package auth

import (
	"errors"
	"strings"
)

// Role represents a user role in the system.
type Role string

const (
	Admin Role = "admin"
	User  Role = "user"
)

// Policy defines the permissions for each role.
type Policy struct {
	Role       Role
	Permissions []string
}

// RBAC holds the role-based access control policies.
var RBAC = []Policy{
	{
		Role:       Admin,
		Permissions: []string{"create_user", "delete_user", "view_user", "transfer_funds", "view_balance"},
	},
	{
		Role:       User,
		Permissions: []string{"transfer_funds", "view_balance"},
	},
}

// Authorize checks if a user has permission to perform an action.
func Authorize(role Role, action string) error {
	for _, policy := range RBAC {
		if policy.Role == role {
			for _, permission := range policy.Permissions {
				if strings.EqualFold(permission, action) {
					return nil
				}
			}
		}
	}
	return errors.New("access denied")
}