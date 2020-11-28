package service

func mask(str string, start, end int) string {
	maskLen := end - start
	if maskLen < 0 {
		panic("end cannot be greater than start")
	}
	var maskStr string
	for i := 0; i <= maskLen; i++ {
		maskStr += "*"
	}
	runes := []rune(str)
	return string(append(runes[:start], append([]rune(maskStr), runes[end+1:]...)...))
}

func maskLeft(str string, reserve int) string {
	runes := []rune(str)
	if len(runes)-reserve < 0 {
		panic("length of reserved string is out of range")
	}
	for i := 0; i < len(runes)-reserve; i++ {
		runes[i] = '*'
	}
	return string(runes)
}

func maskRight(str string, reserve int) string {
	runes := []rune(str)
	if len(runes)-reserve < 0 {
		panic("length of reserved string is out of range")
	}
	for i := len(runes) - 1; i > len(runes)-reserve; i-- {
		runes[i] = '*'
	}
	return string(runes)
}
