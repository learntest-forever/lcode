package longestsubstring

func LongestSubString(s string){
	m := make(map[string]int)
	var start, end, sublen int
	var s1 string
	for k,v := range s{
		if k == 0 {
			s1 = s1+string(v)
			sublen = 1
		}else {
			for i,j := range s1{
				if v == j{
					continue
				}else {
					s1 = s1+string(v)

				}
			}
		}
	}
}