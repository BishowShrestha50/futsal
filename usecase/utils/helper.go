package utils

import (
	"encoding/json"
	"fmt"
	"futsal/model"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func ParseStringToTime(date, timeData string) (*time.Time, error) {
	dateString := date

	// Time string
	timeString := timeData

	// Concatenate date and time strings
	combinedString := dateString + " " + timeString

	// Define the layout for the combined string
	layout := "2006-01-02 15:04"

	// Parse the combined string and convert to time.Time
	t, err := time.Parse(layout, combinedString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil, err
	}

	fmt.Println("Parsed time:", t)
	return &t, nil
}
func GenerateToken(userID string) (string, error) {
	// Create a new token object.
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims for the token, including the user ID and expiration time.
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign the token with a secret key.
	secretKey := "qwerty" // Replace with your own secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecodeToken(tokenString string) (string, error) {
	// Parse the token and extract the claims.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Provide the secret key used for signing the token.
		secretKey := "qwerty" // Replace with your own secret key
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", err
	}

	// Verify the token's signature.
	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	// Extract the user ID from the claims.
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	userID, ok := claims["userID"].(string)
	if !ok {
		return "", fmt.Errorf("invalid user ID claim")
	}

	return userID, nil

}

func RemoveDuplicates(arr1 []string, arr2 []model.BookFutsal) []string {
	result := make([]string, 0)

	// Create a map to track elements in arr2
	existingElements := make(map[string]bool)
	for _, element := range arr2 {
		existingElements[element.Time] = true
	}

	// Iterate through arr1 and add non-duplicate elements to result
	for _, element := range arr1 {
		if !existingElements[element] {
			result = append(result, element)
		}
	}

	return result
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
