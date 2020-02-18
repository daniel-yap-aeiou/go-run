package unittest

import (
	"testing"
)

type unitTestData struct {
	counter     int
	action      string
	exp_counter int
	exp_action  string
}

var unitTestValues = []unitTestData{
	{1, "test", 1, "TEST"}, {1, "debug", 1, "DEBUG"}, {1, "run", 1, "RUN"},
}

/*
* TO RUN UNIT TEST:
* PS C:\Users\Daniel Yap\Documents\Go\unittest> go test
*/
func TestMe(t *testing.T) {
	for _, e1 := range unitTestValues {
		tmpAction := MyToUpperCase(e1.action)

		if e1.counter != e1.exp_counter {
			t.Errorf("[Test Name: %s * Input: %d * Expected: %d * Got %d]\n",
				t.Name(), e1.counter, e1.exp_counter, e1.counter)
		}

		if tmpAction != e1.exp_action {
			t.Errorf("[Test Name: %s * Input: %s * Expected: %s * Got %s]\n",
				t.Name(), e1.action, e1.exp_action, e1.action)
		}
	}
}
