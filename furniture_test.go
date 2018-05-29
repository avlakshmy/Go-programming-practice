package furniture_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "Practice/Furniture"
)

func TestFurniture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Furniture Suite")
}

var _ = Describe("Furniture", func() {
	var (
		c Furniture
		t Furniture
		s Furniture
		b Furniture
	)		
	
	BeforeEach(func() {
        	c = Furniture{
            		Name:  "Chair",
            		Material: "Teak wood",
            		Price: 50,
        	}

        	t = Furniture{
            		Name:  "Table",
            		Material: "Oak wood",
            		Price: 100,
        	}

		s = Furniture{
            		Name:  "Sofa",
            		Material: "Rosewood",
            		Price: 300,
        	}

		b = Furniture{
            		Name:  "Bed",
            		Material: "Mahogany",
            		Price: 350,
        	}
    	})

	Context("Chair", func() {
		It("is a chair", func() {
			Expect(c.RetName()).Should(Equal("Chair"))
		})
		It("is made of teak wood", func() {
			Expect(c.RetMaterial()).Should(Equal("Teak wood"))
		}) 
		It("has price $50", func() {
			Expect(c.RetPrice()).Should(Equal(50))
		})
	})

	Context("Table", func() {
		It("is a table", func() {
			Expect(t.RetName()).Should(Equal("Table"))
		})
		It("is made of oak wood", func() {
			Expect(t.RetMaterial()).Should(Equal("Oak wood"))
		}) 
		It("has price $100", func() {
			Expect(t.RetPrice()).Should(Equal(100))
		})
	})

	Context("Sofa", func() {
		It("is a sofa", func() {
			Expect(s.RetName()).Should(Equal("Sofa"))
		})
		It("is made of rosewood", func() {
			Expect(s.RetMaterial()).Should(Equal("Rosewood"))
		}) 
		It("has price $300", func() {
			Expect(s.RetPrice()).Should(Equal(300))
		})
	})

	Context("Bed", func() {
		It("is a bed", func() {
			Expect(b.RetName()).Should(Equal("Bed"))
		})
		It("is made of mahogany", func() {
			Expect(b.RetMaterial()).Should(Equal("Mahogany"))
		}) 
		It("has price $350", func() {
			Expect(b.RetPrice()).Should(Equal(350))
		})
	})
})
