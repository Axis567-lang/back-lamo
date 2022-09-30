package persistence

import (
	"obra-blanca/promoter"
	"testing"
)

func TestPersistence(t *testing.T) {
	var firstPersistence promoter.Persistence = persistence{testFileName}
	const firstPromoterName = "test-promoter-guy"
	const secondPromoterName = "test-other-promoter-guy"
	var firstPromoter = promoter.New(firstPromoterName)
	var secondPromoter = promoter.New(secondPromoterName)
	err := firstPersistence.AddPromoter(firstPromoter)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = firstPersistence.AddPromoter(secondPromoter)
	if err != nil {
		t.Fatal(err.Error())
	}

	var secondPersistence promoter.Persistence = persistence{testFileName}
	retrievedPromoters, err := secondPersistence.GetPromoters()
	if err != nil {
		t.Fatal(err.Error())
	}
	if _, exists := retrievedPromoters[firstPromoterName]; !exists {
		t.Fatal("promoter not found")
	}
	if _, exists := retrievedPromoters[secondPromoterName]; !exists {
		t.Fatal("promoter not found")
	}

	secondPersistence.DeletePromoter(secondPromoterName)
	retrievedPromoters, err = secondPersistence.GetPromoters()
	if err != nil {
		t.Fatal(err.Error())
	}
	if _, exists := retrievedPromoters[firstPromoterName]; !exists {
		t.Fatal("promoter not found")
	}
	if _, exists := retrievedPromoters[secondPromoterName]; exists {
		t.Fatal("promoter not deleted")
	}
}

const testFileName = "promoters.test"
