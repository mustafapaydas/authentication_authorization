package helper

import (
	"authenticaiton-authorization/entity"
	"authenticaiton-authorization/utils"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
)

func getRoleNames(user *entity.User) []string {
	roleNames := []string{}
	for _, role := range user.Roles {
		roleNames = append(roleNames, role.Name)
	}
	return roleNames
}

func CreateJWTToken(user *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.MapClaims{
		"userId":      user.UserId,
		"iss":         "issuer",
		"username":    user.UserName,
		"sub":         "subject",
		"verify":      user.Verify,
		"roles":       getRoleNames(user),
		"phoneNumber": user.PhoneNumber,
		"email":       user.Email,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})

	_key := loadPrivateKey(os.Getenv("PRIVATE_KEY_FILE_PATH"))
	tokenString, err := token.SignedString(_key)
	if err != nil {
		return "", &utils.BusinessException{
			Message: err.Error(),
		}
	}

	return tokenString, nil
}

func loadPrivateKey(path string) []byte {
	key, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to load private key: %v", err)
	}
	return key
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	publicKeyPEM := loadPrivateKey("PUBLIC_KEY_FILE_PATH")

	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPub, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return rsaPub, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
