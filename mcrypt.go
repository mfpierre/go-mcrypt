package mcrypt

/*
#cgo LDFLAGS: -lmcrypt
#include <stdlib.h>
#include <string.h>
#include <mcrypt.h>
#define FAILED_MCRYPT_MODULE 1
#define INVALID_KEY_LENGTH 2
#define INVALID_IV_LENGTH 3
#define FAILED_TO_ENCRYPT_DATA 4
#define FAILED_TO_DECRYPT_DATA 5
#define INVALID_DATA_LENGTH 6

// getError convert a given error code to its string representation
const char* getError(int err) {
	switch (err) {
	case FAILED_MCRYPT_MODULE:
		return "Failed to open mcrypt module";
	case INVALID_KEY_LENGTH:
		return "Invalid key length";
	case INVALID_IV_LENGTH:
		return "Invalid iv length";
	case FAILED_TO_ENCRYPT_DATA:
		return "Failed to encrypt data";
	case FAILED_TO_DECRYPT_DATA:
		return "Failed to decrypt data";
	case INVALID_DATA_LENGTH:
		return "Invalid data length";
	}
	return mcrypt_strerror(err);
}

char* encrypt(void* key, int keyLength, void* iv, int ivLength, char* data, int* length, char* algo, char* mode, int* err) {
	if (*length <= 0) {
		*err = INVALID_DATA_LENGTH;
		return NULL;
	}
	int i;
	MCRYPT td = mcrypt_module_open(algo, NULL, mode, NULL);
	if (td == MCRYPT_FAILED) {
		*err = FAILED_MCRYPT_MODULE;
		return NULL;
	}
	int requiredKeySize = mcrypt_enc_get_key_size(td);
	int requiredIvSize = mcrypt_enc_get_iv_size(td);
	// make sure the key and iv are the correct sizes
	if (keyLength != requiredKeySize) {
		*err = INVALID_KEY_LENGTH;
		mcrypt_module_close(td);
		return NULL;
	}
	if (ivLength != requiredIvSize) {
		*err = INVALID_IV_LENGTH;
		mcrypt_module_close(td);
		return NULL;
	}
	*err = mcrypt_generic_init(td, key, keyLength, iv);
	if (*err) {
		mcrypt_generic_deinit(td);
		mcrypt_module_close(td);
		return NULL;
	}
	// get the block size
	int blockSize = mcrypt_enc_get_block_size(td);
	// determine the new length if needed
	int newLength = 0;
	// if blockSize is greater than length, expand length to the blockSize
	if (blockSize > *length) {
		newLength = blockSize;
	} else {
		int lengthBlockMod = *length % blockSize;
		if (lengthBlockMod) {
			// if length is not multiple of blockSize, make it the next highest blockSize
			newLength = *length - lengthBlockMod + blockSize;
		} else {
			// we do not need to change the length
			newLength = *length;
		}
	}
	// allocate and copy the data to the output value
	char* output = malloc(sizeof *output * newLength);
	// append byte zeroes to the output array if needed
	for (i = *length; i < newLength; ++i) {
		output[i] = 0;
	}
	memcpy(output, data, *length);
	// update the length to the reallocated length
	*length = newLength;
	// loop through the output data by blockSize at a time
	for (i = 0; i < *length; i += blockSize) {
		// encrypt the block of output[i] plus blockSize
		if (mcrypt_generic(td, output+i, blockSize)) {
			*err = FAILED_TO_ENCRYPT_DATA;
			mcrypt_generic_deinit(td);
			mcrypt_module_close(td);
			free(output);
			return NULL;
		}
	}
	// finish up mcrypt
	mcrypt_generic_deinit(td);
	mcrypt_module_close(td);
	// return the encrypted data
	return output;
}

char* decrypt(void* key, int keyLength, void* iv, int ivLength, char* data, int* length, char* algo, char* mode, int* err) {
	int i;
	MCRYPT td = mcrypt_module_open(algo, NULL, mode, NULL);
	if (td == MCRYPT_FAILED) {
		*err = FAILED_MCRYPT_MODULE;
		mcrypt_module_close(td);
		return NULL;
	}
	int requiredKeySize = mcrypt_enc_get_key_size(td);
	int requiredIvSize = mcrypt_enc_get_iv_size(td);
	// make sure the key and iv are the correct sizes
	if (keyLength != requiredKeySize) {
		*err = INVALID_KEY_LENGTH;
		mcrypt_module_close(td);
		return NULL;
	}
	if (ivLength != requiredIvSize) {
		*err = INVALID_IV_LENGTH;
		mcrypt_module_close(td);
		return NULL;
	}
	*err = mcrypt_generic_init(td, key, keyLength, iv);
	if (*err) {
		mcrypt_generic_deinit(td);
		mcrypt_module_close(td);
		return NULL;
	}
	// get the block size
	int blockSize = mcrypt_enc_get_block_size(td);
	if (*length < blockSize || *length % blockSize) {
		*err = INVALID_DATA_LENGTH;
		mcrypt_generic_deinit(td);
		mcrypt_module_close(td);
		return NULL;
	}
	// allocate and copy the data to the output value
	char* output = malloc(sizeof *output * *length);
	memcpy(output, data, *length);
	// loop through the output data by blockSize at a time
	for (i = 0; i < *length; i += blockSize) {
		// decrypt the block of output[i] plus blockSize
		if (mdecrypt_generic(td, output+i, blockSize)) {
			*err = FAILED_TO_DECRYPT_DATA;
			mcrypt_generic_deinit(td);
			mcrypt_module_close(td);
			free(output);
			return NULL;
		}
	}
	// finish up mcrypt
	mcrypt_generic_deinit(td);
	mcrypt_module_close(td);
	// return the decrypted data
	return output;
}
*/
import "C"

import (
	"bytes"
	"errors"
	"unsafe"
)

func Encrypt(key []byte, iv []byte, data []byte, algo string, mode string) ([]byte, error) {
	// keep track of the size of the input data
	length := C.int(len(data))
	if length == 0 {
		return nil, errors.New("Invalid data size of 0")
	}
	// keep track of any errors that occur on encryption
	err := C.int(0)
	// convert algo and mode strings
	var calgo *C.char = C.CString(algo)
	defer C.free(unsafe.Pointer(calgo))
	var cmode *C.char = C.CString(mode)
	defer C.free(unsafe.Pointer(cmode))
	// encrypt the data
	encryptedData := C.encrypt(unsafe.Pointer(&key[0]), C.int(len(key)), unsafe.Pointer(&iv[0]), C.int(len(iv)), (*C.char)(unsafe.Pointer(&data[0])), (*C.int)(unsafe.Pointer(&length)), (*C.char)(unsafe.Pointer(calgo)), (*C.char)(unsafe.Pointer(cmode)), (*C.int)(unsafe.Pointer(&err)))
	// ensure that memory is freed on the encrypted data after it is converted to Go bytes
	defer C.free(unsafe.Pointer(encryptedData))

	// if err is not 0, there is an error
	if int(err) != 0 {
		return nil, errors.New(C.GoString(C.getError(err)))
	}

	// return the Go bytes of the encrypted data
	return C.GoBytes(unsafe.Pointer(encryptedData), length), nil
}

func Decrypt(key []byte, iv []byte, data []byte, algo string, mode string) ([]byte, error) {
	// keep track of the size of the input data
	length := C.int(len(data))
	if length == 0 {
		return nil, errors.New("Invalid data size of 0")
	}
	// keep track of any errors that occur on decryption
	err := C.int(0)
	// convert algo and mode
	var calgo *C.char = C.CString(algo)
	defer C.free(unsafe.Pointer(calgo))
	var cmode *C.char = C.CString(mode)
	defer C.free(unsafe.Pointer(cmode))
	// decrypt the data
	decryptedData := C.decrypt(unsafe.Pointer(&key[0]), C.int(len(key)), unsafe.Pointer(&iv[0]), C.int(len(iv)), (*C.char)(unsafe.Pointer(&data[0])), (*C.int)(unsafe.Pointer(&length)), (*C.char)(unsafe.Pointer(calgo)), (*C.char)(unsafe.Pointer(cmode)), (*C.int)(unsafe.Pointer(&err)))
	// ensure that memory is freed on the decrypted data after it is converted to Go bytes
	defer C.free(unsafe.Pointer(decryptedData))

	// if err is not 0, there is an error
	if int(err) != 0 {
		return nil, errors.New(C.GoString(C.getError(err)))
	}

	decryptedBytes := C.GoBytes(unsafe.Pointer(decryptedData), length)
	// trim ending null bytes
	decryptedBytes = bytes.Trim(decryptedBytes, "\x00")

	return decryptedBytes, nil
}
