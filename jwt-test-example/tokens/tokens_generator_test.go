package tokens

import (
	"fmt"
	"testing"

	"github.com/cristalhq/jwt/v5"
)

func TestGenerate(t *testing.T) {
	token := Generate()
	fmt.Println(token)
	if token == "" {
		t.Error("Token Vazio")
	}

	if len(token) < 100 {
		t.Error("Token Muito curto")
	}

	if len(token) > 1000 {
		t.Error("Token Muito longo")
	}

}

func TestVerifyToken(t *testing.T) {
	token := Generate()
	key := []byte(`secret`)
	verifier, err := jwt.NewVerifierHS(jwt.HS256, key)
	if err != nil {
		t.Error("Falhou na verificação")
	}
	newToken, err := jwt.Parse([]byte(token), verifier)
	if err != nil {
		t.Error("Falhou ao fazer o parse")
	}
	if newToken == nil {
		t.Error("Token não pode ser nulo")
	}
	newToken.Claims()
}
