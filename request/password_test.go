package request

import (
	"log"
	"testing"
)

func TestPassword_FindCount(t *testing.T) {

	results := PasswordCount([]string{"passw0rd"})

	if len(results) < 1 {
		t.Errorf("PasswordCount() expected results")
	} else {
		if results[0].Count < 1 {
			t.Errorf("Password expected non-zero matches")
		} else {
			log.Println("Count = ", results[0].Count)
		}
	}

}
