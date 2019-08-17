package secret

var handShakesArray = []string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

// Handshake function returns the array of secret handshakes
func Handshake(code uint) []string {
	var handshakes []string

	for i := 0; code > 0; i++ {
		if code%2 == 1 && i < 4 {
			handshakes = append(handshakes, handShakesArray[i])
		} else if code%2 == 1 {
			length := len(handshakes)
			for i := 0; i < length/2; i++ {
				handshakes[i], handshakes[length-i-1] = handshakes[length-i-1], handshakes[i]
			}
		}
		code /= 2
	}

	return handshakes
}
