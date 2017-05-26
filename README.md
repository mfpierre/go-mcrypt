# go-mcrypt

[![Build Status](https://travis-ci.org/mfpierre/go-mcrypt.svg?branch=master)](https://travis-ci.org/mfpierre/go-mcrypt)

Go bindings for mcrypt library.

Should be compatible with all algo/modes supported by libmcrypt

## Usage
```go
key := []byte("here is a random key of 32 bytes") // 32 bytes
plaintext := []byte("here is what I want to encrypt")
iv := make([]byte, 16)

// using cast-256 in ECB mode
encrypted, _ := encrypt(key, iv, plaintext, "cast-256", "ecb")
```

## Requirements
 * libmcrypt (http://mcrypt.sourceforge.net/)

## Credits
Thanks to https://github.com/tblyler/go-mcrypt for initial implementation with rijndael
