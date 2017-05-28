# go-mcrypt

[![Build Status](https://travis-ci.org/mfpierre/go-mcrypt.svg?branch=master)](https://travis-ci.org/mfpierre/go-mcrypt)

Go bindings for mcrypt library.

Should be compatible with most algo/modes supported by libmcrypt.

## Requirements
 * libmcrypt (http://mcrypt.sourceforge.net/)

## Usage
```go
key := []byte("here is a random key of 32 bytes")
plaintext := []byte("here is what you want to encrypt")
iv := make([]byte, 16)

// using CAST-256 in ECB mode
encrypted, _ := encrypt(key, iv, plaintext, "cast-256", "ecb")
decrypted, _ := decrypt(key, iv, encrypted, "cast-256", "ecb")
```

Below a reminder of IV & Key size that you can use depending on algo/mode settings

|Cipher Name|Block Mode|Block Size|IV Size|Default Key Size|All Key Size(s)|
|--- |--- |--- |--- |--- |--- |
|CAST-128|CBC|8|8|16|16|
|CAST-128|ECB|8|8|16|16|
|CAST-128|OFB|8|8|16|16|
|CAST-128|NOFB|8|8|16|16|
|CAST-128|CFB|8|8|16|16|
|CAST-128|NCFB|8|8|16|16|
|CAST-128|CTR|8|8|16|16|
|GOST|CBC|8|8|32|32|
|GOST|ECB|8|8|32|32|
|GOST|OFB|8|8|32|32|
|GOST|NOFB|8|8|32|32|
|GOST|CFB|8|8|32|32|
|GOST|NCFB|8|8|32|32|
|GOST|CTR|8|8|32|32|
|Rijndael-128|CBC|16|16|32|16 24 32|
|Rijndael-128|ECB|16|16|32|16 24 32|
|Rijndael-128|OFB|16|16|32|16 24 32|
|Rijndael-128|NOFB|16|16|32|16 24 32|
|Rijndael-128|CFB|16|16|32|16 24 32|
|Rijndael-128|NCFB|16|16|32|16 24 32|
|Rijndael-128|CTR|16|16|32|16 24 32|
|Twofish|CBC|16|16|32|16 24 32|
|Twofish|ECB|16|16|32|16 24 32|
|Twofish|OFB|16|16|32|16 24 32|
|Twofish|NOFB|16|16|32|16 24 32|
|Twofish|CFB|16|16|32|16 24 32|
|Twofish|NCFB|16|16|32|16 24 32|
|Twofish|CTR|16|16|32|16 24 32|
|RC4|STREAM|1|0|256||
|CAST-256|CBC|16|16|32|16 24 32|
|CAST-256|ECB|16|16|32|16 24 32|
|CAST-256|OFB|16|16|32|16 24 32|
|CAST-256|NOFB|16|16|32|16 24 32|
|CAST-256|CFB|16|16|32|16 24 32|
|CAST-256|NCFB|16|16|32|16 24 32|
|CAST-256|CTR|16|16|32|16 24 32|
|LOKI97|CBC|16|16|32|16 24 32|
|LOKI97|ECB|16|16|32|16 24 32|
|LOKI97|OFB|16|16|32|16 24 32|
|LOKI97|NOFB|16|16|32|16 24 32|
|LOKI97|CFB|16|16|32|16 24 32|
|LOKI97|NCFB|16|16|32|16 24 32|
|LOKI97|CTR|16|16|32|16 24 32|
|Rijndael-192|CBC|24|24|32|16 24 32|
|Rijndael-192|ECB|24|24|32|16 24 32|
|Rijndael-192|OFB|24|24|32|16 24 32|
|Rijndael-192|NOFB|24|24|32|16 24 32|
|Rijndael-192|CFB|24|24|32|16 24 32|
|Rijndael-192|NCFB|24|24|32|16 24 32|
|Rijndael-192|CTR|24|24|32|16 24 32|
|Safer+|CBC|16|16|32|16 24 32|
|Safer+|ECB|16|16|32|16 24 32|
|Safer+|OFB|16|16|32|16 24 32|
|Safer+|NOFB|16|16|32|16 24 32|
|Safer+|CFB|16|16|32|16 24 32|
|Safer+|NCFB|16|16|32|16 24 32|
|Safer+|CTR|16|16|32|16 24 32|
|WAKE|STREAM|1|0|32|32|
|Blowfish|CBC|8|8|56||
|Blowfish|ECB|8|8|56||
|Blowfish|OFB|8|8|56||
|Blowfish|NOFB|8|8|56||
|Blowfish|CFB|8|8|56||
|Blowfish|NCFB|8|8|56||
|Blowfish|CTR|8|8|56||
|DES|CBC|8|8|8|8|
|DES|ECB|8|8|8|8|
|DES|OFB|8|8|8|8|
|DES|NOFB|8|8|8|8|
|DES|CFB|8|8|8|8|
|DES|NCFB|8|8|8|8|
|DES|CTR|8|8|8|8|
|Rijndael-256|CBC|32|32|32|16 24 32|
|Rijndael-256|ECB|32|32|32|16 24 32|
|Rijndael-256|OFB|32|32|32|16 24 32|
|Rijndael-256|NOFB|32|32|32|16 24 32|
|Rijndael-256|CFB|32|32|32|16 24 32|
|Rijndael-256|NCFB|32|32|32|16 24 32|
|Rijndael-256|CTR|32|32|32|16 24 32|
|Serpent|CBC|16|16|32|16 24 32|
|Serpent|ECB|16|16|32|16 24 32|
|Serpent|OFB|16|16|32|16 24 32|
|Serpent|NOFB|16|16|32|16 24 32|
|Serpent|CFB|16|16|32|16 24 32|
|Serpent|NCFB|16|16|32|16 24 32|
|Serpent|CTR|16|16|32|16 24 32|
|xTEA|CBC|8|8|16|16|
|xTEA|ECB|8|8|16|16|
|xTEA|OFB|8|8|16|16|
|xTEA|NOFB|8|8|16|16|
|xTEA|CFB|8|8|16|16|
|xTEA|NCFB|8|8|16|16|
|xTEA|CTR|8|8|16|16|
|Blowfish|CBC|8|8|56||
|Blowfish|ECB|8|8|56||
|Blowfish|OFB|8|8|56||
|Blowfish|NOFB|8|8|56||
|Blowfish|CFB|8|8|56||
|Blowfish|NCFB|8|8|56||
|Blowfish|CTR|8|8|56||
|enigma|STREAM|1|0|13||
|RC2|CBC|8|8|128||
|RC2|ECB|8|8|128||
|RC2|OFB|8|8|128||
|RC2|NOFB|8|8|128||
|RC2|CFB|8|8|128||
|RC2|NCFB|8|8|128||
|RC2|CTR|8|8|128||
|3DES|CBC|8|8|24|24|
|3DES|ECB|8|8|24|24|
|3DES|OFB|8|8|24|24|
|3DES|NOFB|8|8|24|24|
|3DES|CFB|8|8|24|24|
|3DES|NCFB|8|8|24|24|
|3DES|CTR|8|8|24|24|


## Credits
Thanks to https://github.com/tblyler/go-mcrypt for initial implementation with rijndael
