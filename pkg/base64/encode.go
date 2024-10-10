package base64

import "encoding/base64"

func Encode(src []byte) (dest []byte) {
	dest = make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dest, src)
	return dest
}
