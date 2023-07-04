// Package structs contain every boxie public structs
package structs

// User is boxie user in a struct
type User struct {
	ID         int32
	Name       string
	Email      string
	CreatedAt  string
	UpdatedAt  string
	Password   string
	InviteCode string
	Token      string
}

// CreateUserRequest is Create User Request DTO
type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest is User Login Request DTO
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AcceptInviteRequest is Accept Invite Request DTO
type AcceptInviteRequest struct {
	Code string `json:"code"`
}

// SetPasswordRequest is Set User Password Request DTO
type SetPasswordRequest struct {
	Code                 string `json:"code"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

// DeleteUserRequest is Delete User Request DTO
type DeleteUserRequest struct {
	ID int32 `json:"id"`
}

// UserService is a public UserService
type UserService struct {
	Login       func() error
	Invite      func() error
	SetPassword func() error
}
