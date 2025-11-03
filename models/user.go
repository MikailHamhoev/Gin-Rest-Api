// gin-rest-api/models/user.go
package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email" binding:"required,email"`
	Username  string    `json:"username" binding:"required,min=3"`
	Password  string    `json:"password" binding:"required,min=6"`
	FullName  string    `json:"full_name" binding:"required"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateProfileRequest struct {
	Username string `json:"username" binding:"min=3"`
	FullName string `json:"full_name"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

// In-memory database
var users = []*User{}
var nextID = 1

func CreateUser(user *User) error {
	// Check if email already exists
	for _, u := range users {
		if u.Email == user.Email {
			return &AppError{Message: "Email already exists", Code: 409}
		}
		if u.Username == user.Username {
			return &AppError{Message: "Username already exists", Code: 409}
		}
	}

	user.ID = nextID
	nextID++
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	if user.Role == "" {
		user.Role = "user"
	}

	users = append(users, user)
	return nil
}

func GetUserByEmail(email string) (*User, error) {
	for _, user := range users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, &AppError{Message: "User not found", Code: 404}
}

func GetUserByID(id int) (*User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, &AppError{Message: "User not found", Code: 404}
}

func GetAllUsers() []*User {
	// Return copy without passwords
	var result []*User
	for _, user := range users {
		userCopy := *user
		userCopy.Password = ""
		result = append(result, &userCopy)
	}
	return result
}

func UpdateUser(id int, updates *UpdateProfileRequest) (*User, error) {
	for _, user := range users {
		if user.ID == id {
			if updates.Username != "" {
				// Check if username is taken by another user
				for _, u := range users {
					if u.ID != id && u.Username == updates.Username {
						return nil, &AppError{Message: "Username already taken", Code: 409}
					}
				}
				user.Username = updates.Username
			}
			if updates.FullName != "" {
				user.FullName = updates.FullName
			}
			user.UpdatedAt = time.Now()
			return user, nil
		}
	}
	return nil, &AppError{Message: "User not found", Code: 404}
}

func GetStats() map[string]interface{} {
	stats := make(map[string]interface{})
	stats["total_users"] = len(users)

	roleCount := make(map[string]int)
	for _, user := range users {
		roleCount[user.Role]++
	}
	stats["users_by_role"] = roleCount

	return stats
}

// Custom error type
type AppError struct {
	Message string
	Code    int
}

func (e *AppError) Error() string {
	return e.Message
}
