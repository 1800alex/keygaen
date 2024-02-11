package git

import (
	"github.com/ProtonMail/gopenpgp/v2/helper"
)

// GenerateKey generates a PGP key using x25519
func GenerateKey(
	fullName string, // Full name of the PGP key's holder
	email string, // Email of the PGP key's holder
	password string, // Password of the PGP key
) ([]byte, error) { // key, error
	key, err := helper.GenerateKey(fullName, email, []byte(password), "x25519", 0)
	if err != nil {
		return []byte{}, err
	}

	return []byte(key), err
}
