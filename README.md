# go-size
print humans readable file size

progressbar

# Install

```bash
go get -u github.com/wux1an/go-size
```

# Usage

```golang
package main

import (
	"fmt"
	"github.com/wux1an/go-size"
)

func main() {
	fmt.Println(0 * size.Byte)
	fmt.Println(1 * size.Byte)
	fmt.Println(2*size.Byte + 1*size.Bit)
	fmt.Println(20*size.KB + 512*size.Byte)
	fmt.Println(200*size.MB + 1024*size.KB)
	fmt.Println(1024 * 10000 * size.GB)
}
```

output

```
0 Bit
1 Byte
2 Byte
20.5 KB
201 MB
9.8 PB
```
