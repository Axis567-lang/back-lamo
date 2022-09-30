package persistence

import (
	"obra-blanca/distributor"
	"testing"
)

func TestPersistence(t *testing.T) {
	const expectedDistributorName = "test-name"
	const expectedSecondDistributorName = "test-name2"
	const testFileName = "distributor.test"
	testPersistence := persistence{testFileName}
	testDistributor := distributor.DTO{
		DistributorName:         expectedDistributorName,
		DistributorNegotiations: nil,
	}
	secondDistributor := distributor.DTO{
		DistributorName:         expectedSecondDistributorName,
		DistributorNegotiations: nil,
	}
	err := testPersistence.AddDistributor(testDistributor)
	if err != nil {
		t.Fatal(err.Error())
	}
	err = testPersistence.AddDistributor(secondDistributor)
	if err != nil {
		t.Fatal(err.Error())
	}

	otherPersistence := persistence{testFileName}
	retrievedDistributors, err := otherPersistence.GetDistributors()
	if err != nil {
		t.Fatal(err.Error())
	}

	if _, exists := retrievedDistributors[expectedDistributorName]; !exists {
		t.Fatal("distributor not retrieved")
	}
	if _, exists := retrievedDistributors[expectedSecondDistributorName]; !exists {
		t.Fatal("distributor not retrieved")
	}

	err = otherPersistence.DeleteDistributor(expectedSecondDistributorName)
	if err != nil {
		t.Fatal(err.Error())
	}
	retrievedDistributors, err = otherPersistence.GetDistributors()
	if err != nil {
		t.Fatal(err.Error())
	}

	if _, exists := retrievedDistributors[expectedDistributorName]; !exists {
		t.Fatal("distributor not retrieved")
	}
	if _, exists := retrievedDistributors[expectedSecondDistributorName]; exists {
		t.Fatal("distributor not deleted")
	}
}
