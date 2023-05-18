package utility

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSignedDetails struct {
	UserID primitive.ObjectID
	jwt.StandardClaims
}

func GenerateTokenForUser(userID primitive.ObjectID) (string, error) {
	var SECRETKEY = os.Getenv("SECRETKEY")
	claims := &UserSignedDetails{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour * 7).Unix(), //Token is valid for 7 days
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRETKEY))

	return token, err
}

func GetUserIDFromToken(tokenString string) (primitive.ObjectID, error) {
	var SECRETKEY = os.Getenv("SECRETKEY")
	token, err := jwt.ParseWithClaims(tokenString, &UserSignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRETKEY), nil
	})

	if claims, ok := token.Claims.(*UserSignedDetails); ok && token.Valid {
		return claims.UserID, nil
	} else {
		return primitive.NilObjectID, err
	}
}
