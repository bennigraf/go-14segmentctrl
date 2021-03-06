package main

// import "fmt"
import "strconv"

// chars taken from here: https://github.com/adafruit/Adafruit_Python_LED_Backpack/blob/master/Adafruit_LED_Backpack/AlphaNum4.py
var chars map[byte]uint16 = map[byte]uint16{
	' ':  parseBinary("0000000000000000"),
	'!':  parseBinary("0100000000000010"),
	'"':  parseBinary("0000001000100000"),
	'#':  parseBinary("0001001011001110"),
	'$':  parseBinary("0001001011101101"),
	'%':  parseBinary("0000110000100100"),
	'&':  parseBinary("0010001101011101"),
	'\'': parseBinary("0000010000000000"),
	'(':  parseBinary("0010010000000000"),
	')':  parseBinary("0000100100000000"),
	'*':  parseBinary("0011111111000000"),
	'+':  parseBinary("0001001011000000"),
	',':  parseBinary("0000100000000000"),
	'-':  parseBinary("0000000011000000"),
	'.':  parseBinary("0100000000000000"),
	'/':  parseBinary("0000110000000000"),
	'0':  parseBinary("0000110000111111"),
	'1':  parseBinary("0000000000000110"),
	'2':  parseBinary("0000000011011011"),
	'3':  parseBinary("0000000010001111"),
	'4':  parseBinary("0000000011100110"),
	'5':  parseBinary("0010000001101001"),
	'6':  parseBinary("0000000011111101"),
	'7':  parseBinary("0000000000000111"),
	'8':  parseBinary("0000000011111111"),
	'9':  parseBinary("0000000011101111"),
	':':  parseBinary("0001001000000000"),
	';':  parseBinary("0000101000000000"),
	'<':  parseBinary("0010010000000000"),
	'=':  parseBinary("0000000011001000"),
	'>':  parseBinary("0000100100000000"),
	'?':  parseBinary("0001000010000011"),
	'@':  parseBinary("0000001010111011"),
	'A':  parseBinary("0000000011110111"),
	'B':  parseBinary("0001001010001111"),
	'C':  parseBinary("0000000000111001"),
	'D':  parseBinary("0001001000001111"),
	'E':  parseBinary("0000000011111001"),
	'F':  parseBinary("0000000001110001"),
	'G':  parseBinary("0000000010111101"),
	'H':  parseBinary("0000000011110110"),
	'I':  parseBinary("0001001000000000"),
	'J':  parseBinary("0000000000011110"),
	'K':  parseBinary("0010010001110000"),
	'L':  parseBinary("0000000000111000"),
	'M':  parseBinary("0000010100110110"),
	'N':  parseBinary("0010000100110110"),
	'O':  parseBinary("0000000000111111"),
	'P':  parseBinary("0000000011110011"),
	'Q':  parseBinary("0010000000111111"),
	'R':  parseBinary("0010000011110011"),
	'S':  parseBinary("0000000011101101"),
	'T':  parseBinary("0001001000000001"),
	'U':  parseBinary("0000000000111110"),
	'V':  parseBinary("0000110000110000"),
	'W':  parseBinary("0010100000110110"),
	'X':  parseBinary("0010110100000000"),
	'Y':  parseBinary("0001010100000000"),
	'Z':  parseBinary("0000110000001001"),
	'[':  parseBinary("0000000000111001"),
	'\\': parseBinary("0010000100000000"),
	']':  parseBinary("0000000000001111"),
	'^':  parseBinary("0000110000000011"),
	'_':  parseBinary("0000000000001000"),
	'`':  parseBinary("0000000100000000"),
	'a':  parseBinary("0001000001011000"),
	'b':  parseBinary("0010000001111000"),
	'c':  parseBinary("0000000011011000"),
	'd':  parseBinary("0000100010001110"),
	'e':  parseBinary("0000100001011000"),
	'f':  parseBinary("0000000001110001"),
	'g':  parseBinary("0000010010001110"),
	'h':  parseBinary("0001000001110000"),
	'i':  parseBinary("0001000000000000"),
	'j':  parseBinary("0000000000001110"),
	'k':  parseBinary("0011011000000000"),
	'l':  parseBinary("0000000000110000"),
	'm':  parseBinary("0001000011010100"),
	'n':  parseBinary("0001000001010000"),
	'o':  parseBinary("0000000011011100"),
	'p':  parseBinary("0000000101110000"),
	'q':  parseBinary("0000010010000110"),
	'r':  parseBinary("0000000001010000"),
	's':  parseBinary("0010000010001000"),
	't':  parseBinary("0000000001111000"),
	'u':  parseBinary("0000000000011100"),
	'v':  parseBinary("0010000000000100"),
	'w':  parseBinary("0010100000010100"),
	'x':  parseBinary("0010100011000000"),
	'y':  parseBinary("0010000000001100"),
	'z':  parseBinary("0000100001001000"),
	'{':  parseBinary("0000100101001001"),
	'|':  parseBinary("0001001000000000"),
	'}':  parseBinary("0010010010001001"),
	'~':  parseBinary("0000010100100000"),
}

func parseBinary(binaryString string) uint16 {
	// var number uint16
	number, _ := strconv.ParseUint(binaryString, 2, 16)
	return uint16(number)
}
