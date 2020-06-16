package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	// initiate an arrya with out lenght,if there is no element in it, the initial length is 0
	var s0 []int
	// everytime you append an item to the array, the length automatically increased
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	// define a slice, that specify its length:5, and assign first 3 elements for it, without value specified, the default value is 0
	// make is the keyword to define a slice
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 33)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	// The initial length is 0, after 1 element append, the cap of it incremented by 8
	// If the 9th element inserted into the array, the cap of the slice incremented by its current length, that means 16
	s := []int{}
	for i := 0; i < 10; i++ {
		// append, it means append element to slice s, and length increment by 1
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	//it defined a pointer Q2 , which point to the original slice year, just from the 4th element
	// that means either Q2 or year, effects are the some.
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	year[3] = "Kill"
	t.Log(Q2, len(Q2), cap(Q2))
	Q2 = append(Q2, "SS")
	t.Log(year)
}

func TestSliceComparing(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}
	// a==b is invalide, because slice could only be compared with nil
	if a[1] == b[1] {
		t.Log("equal")
		t.Log(a, b)
	} else {
		if a == nil {
			t.Log("Error")
		}

	}
}
