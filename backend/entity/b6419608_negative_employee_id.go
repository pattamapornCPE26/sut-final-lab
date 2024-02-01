package entity

import (
	"github.com/asaskevich/govalidator"
	 ."github.com/onsi/gomega"
	"testing"
	// "github.com/pattamapornCPE26/sut-final-lab/entity"
)

func TestID(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`EmployeeID not true` , func(t *testing.T) {
		employee := Employee{
			Name:     "test",
			URL:      "https://github.com/pattamapornCPE26/sut-final-lab",
			Employee: "EM001",
		}

		ok, err := govalidator.ValidateStruct(&employee)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("EmployeeID not true"))
	})
}
