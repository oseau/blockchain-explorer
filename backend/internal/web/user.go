package web

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	formaters = map[string]string{
		"en": `Welcome to Blockchain Explorer!

Sign this message to prove you own this wallet and we'll log you in.

This request will NOT trigger any blockchain transactions or cost any gas.

Your authentication status will reset after 24 hours.

To stop hackers from using your identity, here's a one time unique message they can't guess:
%v`,
	}
)

func getNonce(lang string) string {
	length := 16
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	str := base64.URLEncoding.EncodeToString(bytes)
	if _, ok := formaters[lang]; !ok {
		lang = "en"
	}
	return fmt.Sprintf(formaters[lang], str)
}

func verify(address, message, signature string) bool {
	sig := hexutil.MustDecode(signature)
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}
	msg := accounts.TextHash([]byte(message))
	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false
	}
	recoveredAddr := crypto.PubkeyToAddress(*recovered)
	return strings.EqualFold(address, recoveredAddr.Hex())
}
