package map_ex

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	t.Log(m1)
	t.Logf("Len m1=%d", len(m1))
	m2 := map[int]int{}
	t.Logf("Len m2=%d", len(m2))
	m3 := make(map[int]int, 1)
	t.Logf("len m3=%d", len(m3))
	m3[1] = 1
	m3[2] = 2
	t.Logf("len m3:=%d", len(m3))

}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	if v, ok := m1[5]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 3: 3, 5: 5, 7: 7}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
