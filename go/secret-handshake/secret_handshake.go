package secret

import (
	"fmt"
	"strconv"
)

func Handshake(code uint) []string {
	var binary string
	handShakes := []string{
		"wink",
		"double blink",
		"close your eyes",
		"jump",
		"",
	}

	var handshake []string

	if code < 15 {
		if code % 2 == 1 { handshake = append(handshake, handShakes[0]) }
	}
	for(code > 0){
		binary = strconv.Itoa(int(code % 2)) + binary
		code /= 2
	}

	fmt.Println(binary)

	if(len(binary) > 4){
		binary = binary[1:]
		strlen := len(binary)
		for index, char := range binary {
			if(char == '1'){
				handshake = append(handshake, handShakes[strlen - index - 1])
			}
		}
	} else {
		for index, char := range binary {
			if(char == '1'){
				handshake = append(handshake, handShakes[index])
			}
		}
	}

	return handshake
}