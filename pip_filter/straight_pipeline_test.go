package pip_filter

import "testing"

func TestStraightPipeline(t *testing.T) {
	splitFilter := NewSplitFilter(",")
	toIntFilter := NewToIntFilter()
	sumFilter := NewSumFilter()

	sp := NewStraightPipeline("p1", splitFilter, toIntFilter, sumFilter)

	ret, err := sp.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}
	if ret != 6 {
		t.Fatalf("The expected is 6, but the actual is %d", ret)
	}
}
