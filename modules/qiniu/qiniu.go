package qiniu

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"tech/modules/setting"
	"time"
)

func Sign(key string) string {
	mac := qbox.NewMac(setting.QNACCESS_KEY, setting.QNSECRET_KEY)
	domain := setting.BucketUrl
	deadline := time.Now().Add(time.Second * 3600).Unix()
	privateAccess := storage.MakePrivateURL(mac, domain, key, deadline)
	return privateAccess
}
