package promoter

import (
	"testing"
)

func TestPromoter_HasAName(t *testing.T) {
	const testName Name = "test name"
	var testPromoter Promoter = New(testName)
	var testPromoterName Name = testPromoter.Name()
	if testPromoterName != testName {
		t.Fatalf("invalid name returned: %s", testPromoterName)
	}
}
