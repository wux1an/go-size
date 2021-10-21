package size

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(0 * Byte)
	fmt.Println(1 * Byte)
	fmt.Println(2*Byte + 1*Bit)
	fmt.Println(20*KB + 512*Byte)
	fmt.Println(200*MB + 1024*KB)
	fmt.Println(1024 * 10000 * GB)
}
