package setup

import (
	"fmt"
	"os"
	"strings"

	"open-tutor/internal/services/db"
	"open-tutor/util"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func EnsureDefaultAdmin() {
	// Retrieve default admin credentials from environment variables
	defaultAdminEmail := os.Getenv("DEFAULT_ADMIN_EMAIL")
	defaultAdminPassword := os.Getenv("DEFAULT_ADMIN_PASSWORD")

	// If no default admin email and password are set, we can skip creating the admin
	if defaultAdminEmail == "" || defaultAdminPassword == "" {
		fmt.Println("No default admin email or password provided, skipping creation.")
		return
	}

	// Check if the admin already exists in the database
	var userCount int
	err := db.GetDB().QueryRow(`
		SELECT count(*)
		FROM users
		WHERE email = $1;
	`, defaultAdminEmail).Scan(&userCount)
	if err != nil {
		fmt.Printf("Error checking if default admin exists: %v\n", err)
		return
	}

	if userCount > 0 {
		fmt.Println("Default admin already exists.")
		return
	}

	// Generate user id
	userId := uuid.New().String()

	// Generate password hash via bcrypt for slow hashing (safer)
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(defaultAdminPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error generating password hash: %v\n", err)
		return
	}
	passwordHash := string(hashBytes)

	// Set role to Admin + User
	roleMask := util.RoleMask(util.User)
	roleMask.Add(util.Admin)
	fmt.Println("Default admin roleMask:", roleMask)

	// Insert into DB
	_, err = db.GetDB().Exec(`
		INSERT INTO users (user_id, email, first_name, last_name, password_hash, role_mask)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING user_id, email, first_name, last_name;
	`,
		userId,
		defaultAdminEmail,
		"Admin", // First name
		"User",  // Last name
		passwordHash,
		roleMask,
	)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			fmt.Println("An account already exists with this email, cannot create default admin account.")
		} else {
			fmt.Printf("Error creating default admin: %v\n", err)
		}
		return
	}

	fmt.Printf("Default admin created successfully with email: %s\n", defaultAdminEmail)
}
