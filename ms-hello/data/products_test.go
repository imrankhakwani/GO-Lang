package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Chair",
		Price: 1.25,
		SKU:   "abcd-abcd-njio",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}

}
