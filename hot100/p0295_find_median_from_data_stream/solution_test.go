package p0295_find_median_from_data_stream

import "testing"

func TestMedianFinder(t *testing.T) {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	if got := mf.FindMedian(); got != 1.5 {
		t.Errorf("got %f, want 1.5", got)
	}
	mf.AddNum(3)
	if got := mf.FindMedian(); got != 2.0 {
		t.Errorf("got %f, want 2.0", got)
	}
}
