package iteration

import "testing"

func TestRepeat(t *testing.T) {
	var repeated = Repeat("oua", 5)
	var expected = "ouaouaouaouaoua"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("oua", 5)
	}
}
