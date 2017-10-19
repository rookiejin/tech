package setting

import (
	"gopkg.in/ini.v1"
	"gopkg.in/macaron.v1"
	"log"
	"strings"
	"gopkg.in/mgo.v2"
	"os"
)

var (
	AppVer  string
	AppName string

	// server settings
	HttpPort      int
	ArchivePath   string
	MaxUploadSize int64

	// security settings
	SecretKey         = "!#$)@U)($)(jdNSLAD#$@#"
	LogInRememberDays = 7

	// admin
	AccessToken string

	// globals
	Cfg      *ini.File
	ProdMode bool
	PageSize = 30

	// Qiniu Settings
	BucketName   string
	BucketUrl    string
	QNACCESS_KEY string
	QNSECRET_KEY string

	DbHost []string
	DbName string
	DbPort int
	DbUser string
	DbPass string
)

var Service struct {
	RegisterEmailConfirm bool
	ActiveCodeLives      int
	ResetPwdCodeLives    int
}

func init() {
	sources := []interface{}{"conf/app.ini"}

	var err error

	Cfg, err = macaron.SetConfig(sources[0])
	if err != nil {
		log.Fatal(4, "fail to load the ini file: %v", err)
	}
	AppName = Cfg.Section("").Key("APP_NAME").String()
	if Cfg.Section("").Key("RUN_MODE").MustString("dev") == "prod" {
		ProdMode = true
		macaron.Env = macaron.PROD
		macaron.ColorLog = false
		mgo.SetDebug(true)
		var aLogger *log.Logger
		aLogger = log.New(os.Stdout , "" , log.LstdFlags)
		mgo.SetLogger(aLogger)
	}
	HttpPort = Cfg.Section("server").Key("HTTP_PORT").MustInt(8084)
	MaxUploadSize = Cfg.Section("server").Key("MAX_UPLOAD_SIZE").MustInt64(5)

	DbHost = strings.Split(Cfg.Section("mongodb").Key("HOST").String(), ",")
	DbName = Cfg.Section("mongodb").Key("DB").String()
	DbUser = Cfg.Section("mongodb").Key("USER").String()
	DbPass = Cfg.Section("mongodb").Key("PASS").String()
	DbPort = Cfg.Section("mongodb").Key("PORT").MustInt(27017)

	BucketName = Cfg.Section("qiniu").Key("BUCKET_NAME").String()
	BucketUrl = Cfg.Section("qiniu").Key("BUCKET_URL").String()
	QNACCESS_KEY = Cfg.Section("qiniu").Key("ACCESS_KEY").String()
	QNSECRET_KEY = Cfg.Section("qiniu").Key("SECRET_KEY").String()

}
