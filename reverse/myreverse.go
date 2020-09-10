package reverse

func Myreverse(s1 []int, s2 []int) []int {
	s2 = append(s2, s1[len(s1)-1])
	if len(s1) == 1 {
		return s2
	}
	return Myreverse(s1[:len(s1)-1], s2)
}
