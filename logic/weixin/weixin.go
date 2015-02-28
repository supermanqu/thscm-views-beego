package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"
)

var token string = "ofMA_tz8JXR_Grf6Rn3A5x6kClCk"

type SignatureBody struct {
	signature string
	timestamp string
	nonce     string
	echostr   string
}

func CheckSignature(body *SignatureBody) bool {
	strs := []string{token, body.timestamp, body.nonce}
	sort.Strings(strs)
	var str string

	for _, value := range strs {
		str += value
	}
	fmt.Println(str)
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)

	str = hex.EncodeToString(bs)
	fmt.Println(str)
	return str == body.signature
}
func main() {
	sb := SignatureBody{"a6c0b1fc0c57d12fa258d763a551e64408a4f3e7", "d", "c", "b"}

	fmt.Println(CheckSignature(&sb))
}
