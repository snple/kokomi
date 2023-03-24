package token

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// DeviceClaims struct
type DeviceClaims struct {
	DeviceID string `json:"device_id"`
	jwt.StandardClaims
}

var DeviceTokenSigningKey = []byte(os.Getenv("TOKEN_SALT"))

// ValidateToken func
func ValidateDeviceToken(myToken string) (bool, string) {
	token, err := jwt.ParseWithClaims(myToken, &DeviceClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(DeviceTokenSigningKey), nil
	})

	if err != nil {
		return false, ""
	}

	claims := token.Claims.(*DeviceClaims)
	return token.Valid, claims.DeviceID
}

// ClaimToken func
func ClaimDeviceToken(deviceID string) (string, error) {
	claims := DeviceClaims{
		deviceID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret
	return token.SignedString(DeviceTokenSigningKey)
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
