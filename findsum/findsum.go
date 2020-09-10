package findsum

func Findsum(s1 []int, target int) []int {
	m := make(map[int]int)
	for k, v := range s1 {
		t2 := target - v
		if _, ok := m[t2]; ok {
			return []int{m[t2], k}
		}
		m[v] = k
	}
	return nil
}
