package encipher

var (
	secret    = []byte("234sdfn234ksjdf")
	length    = len(secret)
	jwtSecret = byte(7)
)

func Encrypt(bytes []byte) {
	j := 0
	for i := 0; i < len(bytes); i++ {
		bytes[i] = bytes[i] ^ secret[j]
		j = (j + 1) % length
	}
}

func Jwt(bytes []byte) {
	for i := 0; i < len(bytes); i++ {
		bytes[i] = bytes[i] ^ jwtSecret
	}
}
