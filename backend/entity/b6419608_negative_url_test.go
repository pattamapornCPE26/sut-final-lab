package entity

import (
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"testing"
	// "github.com/pattamapornCPE26/sut-final-lab/entity"
)

// "testing"
// "github.com/asaskevich/govalidator"
// "github.com/onsi/gomega"

func TestURL(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`URL is invalid` , func(t *testing.T) {
		employee := Employee{
			Name:     "test",
			URL:      "aaa",
			Employee: "EM1234567890",
		}

		ok, err := govalidator.ValidateStruct(&employee)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("URL is invalid"))
	})
}
