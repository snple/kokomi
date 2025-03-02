package token

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// NodeClaims struct
type NodeClaims struct {
	NodeID string `json:"node_id"`
	jwt.StandardClaims
}

var NodeTokenSigningKey = []byte(os.Getenv("TOKEN_SALT"))

// ValidateToken func
func ValidateNodeToken(myToken string) (bool, string) {
	token, err := jwt.ParseWithClaims(myToken, &NodeClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(NodeTokenSigningKey), nil
	})

	if err != nil {
		return false, ""
	}

	claims := token.Claims.(*NodeClaims)
	return token.Valid, claims.NodeID
}

// ClaimToken func
func ClaimNodeToken(nodeID string) (string, error) {
	claims := NodeClaims{
		nodeID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret
	return token.SignedString(NodeTokenSigningKey)
}

type SlotClaims struct {
	SlotID string `json:"Slot_id"`
	jwt.StandardClaims
}

var SlotTokenSigningKey = []byte(os.Getenv("TOKEN_SALT"))

// ValidateToken func
func ValidateSlotToken(myToken string) (bool, string) {
	token, err := jwt.ParseWithClaims(myToken, &SlotClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SlotTokenSigningKey), nil
	})

	if err != nil {
		return false, ""
	}

	claims := token.Claims.(*SlotClaims)
	return token.Valid, claims.SlotID
}

// ClaimToken func
func ClaimSlotToken(slotID string) (string, error) {
	claims := SlotClaims{
		slotID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret
	return token.SignedString(SlotTokenSigningKey)
}
