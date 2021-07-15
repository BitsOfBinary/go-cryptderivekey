# go-cryptderivekey
Golang implementation of Window's `CryptDeriveKey` API:
https://docs.microsoft.com/en-us/windows/win32/api/wincrypt/nf-wincrypt-cryptderivekey

Currently this package supports MD5 and SHA1 for derived keys.

## Reference
This is adapted from FireEye's Python 2 implementation in FlareOn 2016:
https://www.fireeye.com/content/dam/fireeye-www/global/en/blog/threat-research/flareon2016/challenge2-solution.pdf
