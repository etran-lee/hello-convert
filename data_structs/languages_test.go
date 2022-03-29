package languages

import "testing"

func TestKeyCheck(t *testing.T) {
	testKeys := []string{"english", "eng", "ENgLiSh", "gERm"}

	for _, testKey := range testKeys {
		if output := KeyCheck(testKey); output.Language == "Error" {
			t.Errorf("The key '%s' did not properly match any languages within the index", testKey)
		}
	}
}
