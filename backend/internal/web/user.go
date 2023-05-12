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
		"zh-CN": `欢迎使用区块链浏览器！

签名此消息以证明您拥有该钱包，我们将为您进行登录操作。

此请求不会触发任何区块链交易或产生任何费用。

您的身份验证状态将在 24 小时后重置。

为防止黑客盗用您的身份，以下是一条一次性的登陆信息：
%v`,
		"zh-HK": `歡迎使用區塊鏈瀏覽器！

簽名此消息以證明您擁有此錢包，我們將為您登錄。

此請求不會觸發任何區塊鏈交易或產生任何費用。

您的身份驗證狀態將在 24 小時後重置。

為了阻止黑客使用您的身份，這裡有一條他們無法猜到的一次性消息：
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
