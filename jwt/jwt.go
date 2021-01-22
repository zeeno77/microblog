package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/zeeno77/microblog/models"
)

func GeneroJWT(usu models.Usuario) (string, error) {
	miClave := []byte("SuperSecretJuampiKey") //Private key

	payload := jwt.MapClaims{
		"email":            usu.Email,
		"nombre":           usu.Nombre,
		"apellidos":        usu.Apellidos,
		"fecha_nacimiento": usu.FechaNacimiento,
		"biografia":        usu.Biografia,
		"ubicacion":        usu.Ubicacion,
		"sitioWeb":         usu.SitioWeb,
		"_id":              usu.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(), //expira en 24hs
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
