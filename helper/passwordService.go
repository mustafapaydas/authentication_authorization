package helper

import (
	"authenticaiton-authorization/utils"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

type Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

func newDefaultParams() *Params {
	return &Params{
		Memory:      128 * 1024,
		Iterations:  2,
		Parallelism: 24,
		SaltLength:  1024,
		KeyLength:   1024,
	}
}

func GeneratePassword(password *string) (*string, error) {
	if password == nil {
		return nil, errors.New("password is not empty")
	}
	params := newDefaultParams()
	salt := make([]byte, params.SaltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	hash := argon2.IDKey([]byte(*password), salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	hashedPassword := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, params.Memory, params.Iterations, params.Parallelism, b64Salt, b64Hash)
	return &hashedPassword, nil

}

func ComparePasswordAndHash(password, encodedHash string) (bool, error) {

	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func decodeHash(encodedHash string) (p *Params, salt, hash []byte, err error) {
	values := strings.Split(encodedHash, "$")
	if len(values) != 6 {
		return nil, nil, nil, &utils.BusinessException{Message: "the encoded hash is not in the correct format"}
	}

	var version int
	_, err = fmt.Sscanf(values[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, &utils.BusinessException{
			Message: "incompatible version of argon2",
		}
	}

	p = &Params{}
	_, err = fmt.Sscanf(values[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism)
	if err != nil {
		return nil, nil, nil, &utils.BusinessException{
			Message: err.Error(),
		}
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(values[4])
	if err != nil {
		return nil, nil, nil, &utils.BusinessException{
			Message: err.Error(),
		}
	}
	p.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(values[5])
	if err != nil {
		return nil, nil, nil, &utils.BusinessException{
			Message: err.Error(),
		}
	}
	p.KeyLength = uint32(len(hash))

	return p, salt, hash, nil
}
