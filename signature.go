package themap

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

type Signature struct {
	Key, Message []byte
	Hash         string
	Valid        bool
}

// NewSignature return new signature
func NewSignature(key, hash string) *Signature {
	return &Signature{Key: []byte(key), Hash: hash}
}

// sign generate signature hash
func (s *Signature) sign() bool {
	mac := hmac.New(sha256.New, s.Key)
	mac.Write(s.Message)
	expectedMAC := mac.Sum(nil)

	if hex.EncodeToString(expectedMAC) == s.Hash {
		s.Valid = true
		return true
	}

	return false
}

func (s *Signature) Verify(p string) bool {

	var keys []string
	var params string

	for _, v := range strings.Split(p, "&") {
		if strings.HasPrefix(v, "Signature") {
			continue // skip signature value
		}
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for k, v := range keys {
		if k > 0 {
			params += "&"
		}
		params += fmt.Sprintf("%s", v)
	}

	s.Message = []byte(params)

	return s.sign()

}
