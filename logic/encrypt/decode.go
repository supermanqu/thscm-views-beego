package main

import (
	"crypto/rand"
	"crypto/x509"
	"fmt"
	"os"
)

var pemPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBvzCCASgCAQAwfzELMAkGA1UEBhMCQ04xDTALBgNVBAgTBHNlM3cxDTALBgNV
BAcTBE5pY2UxGTAXBgNVBAoTEE1pY2hlbCBEdXJhbmQgU0ExIDAeBgNVBAsTF0Zv
dXJuaSBwYXIgVEJTIGludGVybmV0MRUwEwYDVQQDEwx3d3cuc2Uzdy5jb20wgZ8w
DQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBAKvaV2O28FWCYIcifu7gYnB8VuO2EXx/
jHHicJjVywIEYyU+VUuh1E4FeR1qf1GOO52MoJp7kVONC1r24tIoHqGb2nMvMn0T
C0tyzVRBnhBqFjFBaCZR2IjJ0vBsHsHIvOtWaKcYvb/4Iw5JpI8I46eTk903MOb8
0DmtP1OGm8u/AgMBAAGgADANBgkqhkiG9w0BAQUFAAOBgQAOmiYcmwJn/vdpZMt+
Ohb2fUI17gMk+d0aPqgNy6nSlymqqtVjNsZbvCr8NoYReMqcy5KjmJWAszmwgGMw
dfCmDxXAl2DuS3jHFQPiBJS1Z0rbaFB+EvQyWFO4LNxOaU7JOFA/pDdy68jhEfut
LxFxae6uIsQM4hfFURj2unPWAg==
-----END PUBLIC KEY-----
`

var pemPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
PNSUjydbXzNSJyZc6/hrSL1khhOjVbZ6bYINmUGxmdwOtYp4UKFXnG7aV68DjZVq
KI00+Zd7VxOKiwG12Sy6MaBErA/JMtRVJj8wgCNug78REuBagTaBR2K+p8ssocZg
woNFlHGynRbRghDzqccXI4uSgsaU68ItKU0m62sXhjwmXkR1ieIUHl4HRvvuVBOm
tSIj69tTVODl7CQU21oEaC+j5ziKD4mnOLH/xZwXUICC36co0XGFsSjaorbn1xVq
Ja7mLrkY7Rx4W9WFC40Vqoh+4iFN0NA+57Rx2mDkpT28sGkHm/0mxPLqFFY8lJ5B
l/sTyYVD9EkjvyVWgY63hH84++eQhpzlbPcCGvFw9osXqsV+O+Sb0HFE6szsziwJ
RLdBFXZGqqcQx0g5lTCjZauyP1sNc48rJa8kl68B7flUnUdH42jM79NinCE1V0sU
Gttb9Qe9+cHKvNvZjrrgNjDRIltLbxLlfbLbRZGzDEv4K6o+a/LSsQkK/2m/s7YT
S7YxFSZZG6/AtGz5TO5+CeUE5ge8JajCJlsJPvPbLP9wkghWmayM06n5mFRu74pc
TTXGT3k/9I9ZH9Jj88FFPXXQDE3JHkg2G6gHgU32C0M8S+ybWx8ThE1AKVkvvDfx
UyHVZUFEvTJnsQGKyxjuQ3bpKmJevIbtXtqHSTVv+SKZ6xE2YDaf4LvHc3oogsaO
iFHFNf0CIlmXDPKmj14CmSPbCbnRaRsTJ+V3qPp4L0HPPSMSQHogSnYy2QbNqXNQ
V3RIINnaNdgTkrj8xBs5RQn2i6qmErDVfjdHqCsoQ1I/HBMFqpxc2Q==
-----END RSA PRIVATE KEY-----
`

func main() {

	blockType := pemPrivateKey
	password := []byte("QAZwsx123")

	// see http://golang.org/pkg/crypto/x509/#pkg-constants
	cipherType := x509.PEMCipherAES256

	EncryptedPEMBlock, err := x509.EncryptPEMBlock(rand.Reader,
		blockType,
		[]byte("测试内容"),
		password,
		cipherType)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// check if encryption is successful or not

	if !x509.IsEncryptedPEMBlock(EncryptedPEMBlock) {
		fmt.Println("PEM Block is not encrypted!")
		os.Exit(1)
	}

	if EncryptedPEMBlock.Type != blockType {
		fmt.Println("Block type is wrong!")
		os.Exit(1)
	}

	fmt.Printf("Encrypted block \n%v\n", EncryptedPEMBlock)

	fmt.Printf("Encrypted Block Headers Info : %v\n", EncryptedPEMBlock.Headers)

	DecryptedPEMBlock, err := x509.DecryptPEMBlock(EncryptedPEMBlock, password)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Decrypted block message is :  \n%s\n", DecryptedPEMBlock)

}
