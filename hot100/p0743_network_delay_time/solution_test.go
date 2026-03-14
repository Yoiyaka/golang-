package p0743_network_delay_time

import "testing"

func TestNetworkDelayTime(t *testing.T) {
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	if got := networkDelayTime(times, 4, 2); got != 2 {
		t.Errorf("got %d, want 2", got)
	}
	if got := networkDelayTime([][]int{{1, 2, 1}}, 2, 1); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
	if got := networkDelayTime([][]int{{1, 2, 1}}, 2, 2); got != -1 {
		t.Errorf("got %d, want -1", got)
	}
}
