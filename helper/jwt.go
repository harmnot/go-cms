package helper

import (
	"cms/model/respond"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(id, email, role, company string) (token string, err error) {
	//limaBelasMenit, _ := time.ParseDuration("15m")
	//expTime = time.Now().Add(limaBelasMenit)
	claims := jwt.MapClaims{
		"authorized": true,
		"id": id,
		"role": role,
		"company": company,
		"exp": 0,
	}

	str := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	getSecret := ReadENV("", "JWT_KEY")
	if token, err = str.SignedString([]byte(getSecret)); err != nil {
		return
	} else {
		return
	}
}

func ExtractToken(jwtBase string) (token *jwt.Token, err error) {
	getSecret := ReadENV("", "JWT_KEY")
	token, err = jwt.Parse(jwtBase, func(theToken *jwt.Token) (i interface{}, err error) {
		if _, ok := theToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(getSecret), nil
	})
	return
}

func DecodedExtractToken(token *jwt.Token) (user *respond.InitialUser, err error) {
	if claim , ok := token.Claims.(jwt.MapClaims); ok {
		user =  &respond.InitialUser{
			Authorized: true,
			ID:         claim["id"].(string),
			Role:       claim["rule"].(string),
			Company:    claim["company"].(string),
			Exp:        claim["exp"].(int64),
		}
		if err = claim.Valid(); err != nil {
			return nil, fmt.Errorf("token invalid")
		}
		return user, nil
		}
	return nil, err
}

func CheckToken(token string) bool {
	h, _:= ExtractToken(token)
	if !h.Valid {
		return false
	}
	return true
}
