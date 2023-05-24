package _390_Removing_Stars_From_a_String_

func removeStars(s string) string {
	bs := []byte(s)
	var res []byte
	for _, v := range bs {
		if v == '*' {
			res = res[:len(res)-1]
			continue
		}
		res = append(res, v)
	}
	return string(res)
}