package custom

import (
	"fmt"
	"log"
	"output/es-cud-output/proto"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("gog")

func CreateToken(uuid string, user string) *proto.JWTToken {

	claims := jwt.MapClaims{
		"uuid":       uuid,
		"authorized": true,
		"client":     user,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	log.Println(claims)

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		panic(err)
	}

	jwtToken := proto.JWTToken{
		User:  user,
		Token: tokenString,
	}

	return &jwtToken
}

func VerifyJWT(tokenString string) string {

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		panic(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println(claims)
		return "valid"
	} else {
		return "invalid"
	}

}
