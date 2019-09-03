package addrdec

import (
	"fmt"
	"strings"

	"github.com/blocktree/go-owcdrivers/addressEncoder"
)

var (
	ABBCPublicKeyPrefix       = "PUB_"
	ABBCPublicKeyK1Prefix     = "PUB_K1_"
	ABBCPublicKeyR1Prefix     = "PUB_R1_"
	ABBCPublicKeyPrefixCompat = "ABBC"

	//ABBC stuff
	ABBC_mainnetPublic = addressEncoder.AddressType{"abbc", addressEncoder.BTCAlphabet, "ripemd160", "", 33, []byte(ABBCPublicKeyPrefixCompat), nil}
	// ABBC_mainnetPrivateWIF           = AddressType{"base58", BTCAlphabet, "doubleSHA256", "", 32, []byte{0x80}, nil}
	// ABBC_mainnetPrivateWIFCompressed = AddressType{"base58", BTCAlphabet, "doubleSHA256", "", 32, []byte{0x80}, []byte{0x01}}

	Default = AddressDecoderV2{}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	IsTestNet bool
}

// AddressDecode decode address
func (dec *AddressDecoderV2) AddressDecode(pubKey string) ([]byte, error) {

	var pubKeyMaterial string
	if strings.HasPrefix(pubKey, ABBCPublicKeyR1Prefix) {
		pubKeyMaterial = pubKey[len(ABBCPublicKeyR1Prefix):] // strip "PUB_R1_"
	} else if strings.HasPrefix(pubKey, ABBCPublicKeyK1Prefix) {
		pubKeyMaterial = pubKey[len(ABBCPublicKeyK1Prefix):] // strip "PUB_K1_"
	} else if strings.HasPrefix(pubKey, ABBCPublicKeyPrefixCompat) { // "ABBC"
		pubKeyMaterial = pubKey[len(ABBCPublicKeyPrefixCompat):] // strip "ABBC"
	} else {
		return nil, fmt.Errorf("public key should start with [%q | %q] (or the old %q)", ABBCPublicKeyK1Prefix, ABBCPublicKeyR1Prefix, ABBCPublicKeyPrefixCompat)
	}

	ret, err := addressEncoder.Base58Decode(pubKeyMaterial, addressEncoder.NewBase58Alphabet(ABBC_mainnetPublic.Alphabet))
	if err != nil {
		return nil, addressEncoder.ErrorInvalidAddress
	}
	if addressEncoder.VerifyChecksum(ret, ABBC_mainnetPublic.ChecksumType) == false {
		return nil, addressEncoder.ErrorInvalidAddress
	}

	return ret[:len(ret)-4], nil
}

// AddressEncode encode address
func (dec *AddressDecoderV2) AddressEncode(hash []byte) string {
	data := addressEncoder.CatData(hash, addressEncoder.CalcChecksum(hash, ABBC_mainnetPublic.ChecksumType))
	return string(ABBC_mainnetPublic.Prefix) + addressEncoder.EncodeData(data, "base58", ABBC_mainnetPublic.Alphabet)
}
