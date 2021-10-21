# go-size
print humans readable file size

progressbar

# Install

```bash
go get -u github.com/wux1an/go-size
```

# Usage

```golang
fmt.Println(0 * Byte)
fmt.Println(1 * Byte)
fmt.Println(2*Byte + 1*Bit)
fmt.Println(20*KB + 512*Byte)
fmt.Println(200*MB + 1024*KB)
fmt.Println(1024 * 10000 * GB)
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
