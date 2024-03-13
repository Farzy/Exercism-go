package isbn

func IsValidISBN(isbn string) bool {
	sum := 0
	position := 10

	for _, chr := range isbn {
		switch chr {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			sum += int(chr-'0') * position
			position--
		case '-':
			// Skip dashes
		case 'X':
			// 'X' can only be in position 1, counting backwards
			if position != 1 {
				return false
			}
			sum += 10
			position--
		default:
			// All other characters are invalid
			return false
		}
	}
	return position == 0 && sum%11 == 0
}
