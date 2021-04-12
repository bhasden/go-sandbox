package mytypes

type MyInt int

func (i MyInt) Twice() MyInt {
	return i * 2
}

type MyString string

func (s MyString) Len() int {
	return len(s)
}

type MyMultiString []string

func (s MyMultiString) Join(joiner string) string {
	result := ""
	for i := range s {
		result += s[i]
		if i != len(s) - 1 {
			result += joiner
		}
	}
	return result
}
