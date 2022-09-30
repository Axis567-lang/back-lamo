package distributor

import "testing"

func TestDistributor(t *testing.T) {
	var _ Distributor = DTO{}
}

func TestNew(t *testing.T) {
	var testName Name
	var testDistributor Distributor = New(testName)
	testDistributor.Negotiations()
}

func TestDistributor_Name(t *testing.T) {
	var testName Name = "test name"
	var testDistributor Distributor = New(testName)
	var testDistributorName Name = testDistributor.Name()
	if testDistributorName != testName {
		t.Fatalf("did not get expected name (%s) instead got: %s", testName, testDistributorName)
	}
}
