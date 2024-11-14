package tokens

import (
	"testing"

	jwt "github.com/cristalhq/jwt/v4"
	. "github.com/franela/goblin"
)

func TestTokenGeneratorWithGoblin(t *testing.T) {
	g := Goblin(t)

	g.Describe("Token Generator", func() {
		g.Describe("When token gets generated", func() {
			token := Generate()

			g.It("Expect token to exist and be generated successfuly", func() {
				g.Assert(token).IsNotNil()
				g.Assert(len(token) > 100).IsTrue()
				g.Assert(len(token) < 1000)
			})

			g.Describe("and we are verifying it", func() {
				g.It("with the right secret", func() {
					key := []byte(`secret`)
					verifier, err := jwt.NewVerifierHS(jwt.HS256, key)

					g.Assert(err).IsNil()

					newToken, err := jwt.Parse([]byte(token), verifier)

					g.Assert(err).IsNil()
					g.Assert(newToken).IsNotNil()
				})
				g.It("with the wrong secret", func() {
					key := []byte(`different-secret`)
					verifier, err := jwt.NewVerifierHS(jwt.HS256, key)

					g.Assert(err).IsNil()

					newToken, err := jwt.Parse([]byte(token), verifier)

					g.Assert(err).IsNotNil()
					g.Assert(newToken).IsNil()
				})
			})
		})
	})

}
