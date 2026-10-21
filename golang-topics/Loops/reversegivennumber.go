package loops

func ReversetheGivenNumber(x uint) uint {
	var Rev uint

	for x > 0 {

		lastdigit := x % 10
		Rev = (Rev * 10) + uint(lastdigit)
		x = x / 10
	}
	return Rev

}
