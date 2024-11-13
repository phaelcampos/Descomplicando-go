package tokens

import (
	"fmt"

	jwt "github.com/cristalhq/jwt/v5"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Generate() string {
	key := []byte(`secret`)
	signer, err := jwt.NewSignerHS(jwt.HS256, key)
	checkErr(err)
	claims := &jwt.RegisteredClaims{
		Audience: []string{"admin"},
		ID:       "random-unique-string",
	}
	builder := jwt.NewBuilder(signer)
	token, err := builder.Build(claims)
	checkErr(err)
	var valor string = token.String()
	return valor
}
