package blind_sign

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/pkg/errors"
)

var (
	filePath = os.Getenv("PUBLIC_TOKEN_PEM_PATH")
)

func Execute(msg []byte) (string, error) {
	pem, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("Invalid file's path: %w", err)
	}

	priKey, err := generatePrivateKey(pem)
	if err != nil {
		return "", fmt.Errorf("Invalid file's value: %w", err)
	}

	signature, err := blindSign(msg, priKey)
	if err != nil {
		return "", fmt.Errorf("Failed blind sign operation: %w", err)
	}

	return base64.RawStdEncoding.EncodeToString(signature), nil
}

func blindSign(msg []byte, key *rsa.PrivateKey) ([]byte, error) {
	//step1
	bitLen := key.PublicKey.N.BitLen()
	if len(msg)*8 > bitLen {
		return nil, rsa.ErrMessageTooLong
	}

	//step2
	c := new(big.Int).SetBytes(msg)

	//step3, step4（decrypt=step3、check=step4）
	m, err := decryptAndCheck(rand.Reader, key, c)
	if err != nil {
		return nil, err
	}

	//step5
	return m.Bytes(), nil
}

func generatePrivateKey(pf []byte) (key *rsa.PrivateKey, err error) {
	block, _ := pem.Decode(pf)
	if block == nil {
		return nil, errors.New("Invalid private key data")
	}

	if block.Type == "RSA PRIVATE KEY" {
		key, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
	} else if block.Type == "PRIVATE KEY" {
		keyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		var ok bool
		key, ok = keyInterface.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("Not RSA private key")
		}
	} else {
		return nil, errors.Errorf("invalid private key type: %w ", block.Type)
	}

	key.Precompute()

	if err := key.Validate(); err != nil {
		return nil, err
	}

	return
}
