package shape

import (
	"testing"
)

func TestArea(t *testing.T) {
	react := React{4, 5}
	if react.Area() != 20 {
		t.Errorf("error area failed!")
	}
}