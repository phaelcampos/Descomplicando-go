package tokens

import (
	"testing"

	jwt "github.com/cristalhq/jwt/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTokenGenerator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Token Generator Suite")
}

var _ = Describe("Token Generator", func() {
	Context("When token gets generated", func() {
		token := Generate()
		It("Expect token to exist and be generated successfuly", func() {
			Expect(token).ShouldNot(BeEmpty())
			Expect(len(token)).ShouldNot(BeNumerically("<", 100))
			Expect(len(token)).ShouldNot(BeNumerically(">", 1000))
		})

		Context("and we are verifying it", func() {
			It("with the right secret", func() {
				key := []byte(`secret`)
				verifier, err := jwt.NewVerifierHS(jwt.HS256, key)

				Expect(err).ToNot(HaveOccurred())

				newToken, err := jwt.Parse([]byte(token), verifier)

				Expect(err).ToNot(HaveOccurred())
				Expect(newToken).ToNot(BeNil())
			})
			It("with the wrong secret", func() {
				key := []byte(`different-secret`)
				verifier, err := jwt.NewVerifierHS(jwt.HS256, key)
				Expect(err).ToNot(HaveOccurred())

				newToken, err := jwt.Parse([]byte(token), verifier)
				Expect(err).To(HaveOccurred())
				Expect(newToken).To(BeNil())
			})
		})
	})
})
