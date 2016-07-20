package morse

import (
	"fmt"
)

func ExampleDecodeITU() {

	decoded, _ := DecodeITU("-- .. ... - . .-. / - . -..- - / ..--- ----- ----- -----")
	fmt.Print(decoded)
	// Output: mister text 2000
}

func ExampleEncodeITU() {

	encoded := EncodeITU("mister text 2000")
	fmt.Print(encoded)
	// Output: -- .. ... - . .-. / - . -..- - / ..--- ----- ----- -----
}
