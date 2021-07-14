package initialize

import (
	"fmt"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/wechat/v3"
)

var (
	Appid       string //appid 或者服务商模式的 sp_appid
	Mchid       string //商户ID 或者服务商模式的 sp_mchid
	SerialNo    string //商户证书的证书序列号
	ApiV3Key    string //apiV3Key，商户平台获取
	PkContent   string //pkC私钥 apiclient_key.pem 读取后的内容
	WxSerialNo  string
	WxPkContent string
)

func WxPay() {
	//初始化
	client, err := wechat.NewClientV3(Appid, Mchid, SerialNo, ApiV3Key, PkContent)
	if err != nil {
		xlog.Error(err)
		return
	}

	fmt.Println(client)
/*	//获取微信平台证书公钥
	certs, err := wechat.GetPlatformCerts(Mchid, ApiV3Key, SerialNo, PkContent)
	if err != nil {
		xlog.Error(err)
		return
	}
	if certs.Code == wechat.Success {
		for _, cert := range certs.Certs {
			xlog.Infof("生效时间: %s", cert.EffectiveTime)
			xlog.Infof("到期时间: %s", cert.ExpireTime)
			xlog.Infof("WxSerialNo: %s", cert.SerialNo)
			xlog.Infof("WxContent: \n%s", cert.PublicKey)
			WxSerialNo = cert.SerialNo
			WxPkContent = cert.PublicKey
		}
	}
	//设置 微信支付平台证书 和 证书序列号
	client.SetPlatformCert([]byte(WxPkContent), WxSerialNo).AutoVerifySign()
	client.DebugSwitch = gopay.DebugOn
	//初始化bodyMap
	bm := make(gopay.BodyMap)
	bm.Set("nonce_str", util.GetRandomString(32)).
		Set("out_trade_no", time.Now().Format("20060102150405")).
		Set("total_fee", 1).
		Set("spbill_create_ip", "127.0.0.1").
		Set("notify_url", "https://www.fmm.ink").
		Set("trade_type", "MWEB").
		Set("device_info", "WEB").
		SetBodyMap("scene_info", func(bm gopay.BodyMap) {
			bm.SetBodyMap("h5_info", func(bm gopay.BodyMap) {
				bm.Set("type", "Wap")
				bm.Set("wap_url", "https://www.fmm.ink")
				bm.Set("wap_name", "H5测试支付")
			})
		})
	//H5下单
	wxRsp, err := client.V3TransactionH5(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	//异步验签
	notify, err := wechat.V3ParseNotify()
	if err != nil {
		xlog.Error(err)
		return
	}
	err = notify.VerifySign(WxPkContent)
	if err != nil {
		xlog.Error(err)
		return
	}
	//解密回调
	result, err := notify.DecryptCipherText(ApiV3Key)
	if err != nil {
		xlog.Error(err)
		return
	}*/



}
