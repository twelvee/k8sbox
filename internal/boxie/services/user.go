// Package services contains buisness-logic methods of the models
package services

import (
	"github.com/twelvee/boxie/pkg/boxie/structs"
)

// NewUserService creates a new StorageService
func NewUserService() structs.UserService {
	return structs.UserService{}
}
