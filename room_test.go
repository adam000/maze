package maze

import "testing"

func Test_ConnectTo(t *testing.T) {
	for _, direction := range []Direction{North, South, East, West} {
		roomA := Room{}
		roomB := Room{}
		roomA.ConnectTo(&roomB, direction)

		if roomA.Go(direction) != &roomB {
			t.Errorf("Expected roomB to be %s of roomA", direction)
		}
		if roomB.Go(direction.Opposite()) != &roomA {
			t.Errorf("Expected roomB to be %s of roomA", direction.Opposite())
		}
	}
}
