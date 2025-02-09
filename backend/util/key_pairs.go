package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"errors"
	"fmt"
	"time"

	"open-tutor/internal/services/db"
)

type KeyPair struct {
	ID         string
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	CreatedAt  time.Time
}

// Serialize keys to PEM format for storage //
func serializeKeys(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (string, string, error) {
	// Encode private key
	privateBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privatePEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateBytes,
	})

	// Encode public key
	publicBytes := x509.MarshalPKCS1PublicKey(publicKey)
	publicPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicBytes,
	})

	return string(privatePEM), string(publicPEM), nil
}

func generateKeyPair(keyId string) (*KeyPair, error) {
	// Generate new key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("error generating key pair: %w", err)
	}
	publicKey := &privateKey.PublicKey

	// Serialize to PEM for storage
	privatePEM, publicPEM, err := serializeKeys(privateKey, publicKey)
	if err != nil {
		return nil, fmt.Errorf("error serializing keys: %w", err)
	}

	var keyPair KeyPair
	err = db.GetDB().QueryRow(`
		INSERT INTO key_pairs (id, private_key, public_key)
		VALUES ($1, $2, $3)
		RETURNING created_at
	`,
		keyId,
		privatePEM,
		publicPEM,
	).Scan(&keyPair.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("error storing key pair: %w", err)
	}

	keyPair.ID = keyId
	keyPair.PrivateKey = privateKey
	keyPair.PublicKey = publicKey

	return &keyPair, nil
}

func GetKeyPair(keyId string) (*KeyPair, error) {
	var (
		keyPair    KeyPair
		privatePEM string
		publicPEM  string
	)

	// Query keys from Postgres //
	err := db.GetDB().QueryRow(`
    SELECT public_key, private_key, created_at
    FROM key_pairs
    WHERE id = $1
  `, keyId).Scan(&publicPEM, &privatePEM, &keyPair.CreatedAt)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		keyPair, err := generateKeyPair(keyId)
		if err != nil {
			return nil, err
		}

		return keyPair, nil
	}

	// Decode private key //
	privateBlock, _ := pem.Decode([]byte(privatePEM))
	if privateBlock == nil {
		// Generate new key //
		return nil, fmt.Errorf("failed to decode private key PEM")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	// Decode public key //
	publicBlock, _ := pem.Decode([]byte(publicPEM))
	if publicBlock == nil {
		return nil, fmt.Errorf("failed to decode public key PEM")
	}

	publicKey, err := x509.ParsePKCS1PublicKey(publicBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	keyPair.PrivateKey = privateKey
	keyPair.PublicKey = publicKey

	return &keyPair, nil
}
