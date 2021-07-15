# go-cryptderivekey
Golang implementation of Window's `CryptDeriveKey` API:
https://docs.microsoft.com/en-us/windows/win32/api/wincrypt/nf-wincrypt-cryptderivekey

Currently this package supports MD5 and SHA1 for derived keys.

## Example usage
```go
package main

import cryptderivekey "github.com/BitsOfBinary/go-cryptderivekey"
import "fmt"

func main() {
	fmt.Print(cryptderivekey.CryptDeriveKey([]byte("abcd"), "md5"))
}
```
Output:
```
[22 118 170 136 96 216 12 224 194 0 173 22 180 217 69 251 181 84 183 196 135 6 58 240 105 175 24 200 194 172 16 21] <nil>
```

## Reference
This is adapted from FireEye's Python 2 implementation in FlareOn 2016:
https://www.fireeye.com/content/dam/fireeye-www/global/en/blog/threat-research/flareon2016/challenge2-solution.pdf
