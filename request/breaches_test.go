package request

import (
	"testing"
)

func Test_RetrieveBreaches(t *testing.T) {

	_, err := retrieveBreaches("")

	if err == nil {
		t.Errorf("retrieveBreaches should have returned an error")
	}

	results, err := retrieveBreaches("passw0rd")
	if err != nil {
		t.Errorf("retrieveBreaches expected results without an error")
	} else {
		if len(results) < 1 || len(results[0].Name) < 1 {
			t.Errorf("retrieveBreaches expected results")
		}
	}

}

func Test_RetrievePastes(t *testing.T) {

	_, err := retrievePastes("")

	if err == nil {
		t.Errorf("retrievePastes should have returned an error")
	}

	results, err := retrievePastes("sales@ibm.com")
	if err != nil {
		t.Errorf("retrievePastes expected results without an error")
	} else {
		if len(results) < 1 || len(results[0].Source) < 1 {
			t.Errorf("retrievePastes expected results")
		}
	}

}
