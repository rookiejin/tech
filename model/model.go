package model

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"tech/modules/setting"
	"time"
)

const (
	MasterSession    = "master"
	MonotonicSession = "monotonic"
)

var (
	singleton mongoManager
)

type (
	mongoConfiguration struct {
		Hosts    string
		Database string
		Username string
		Password string
		Port     int
	}

	mongoSession struct {
		mongoDialInfo *mgo.DialInfo
		mongoSession  *mgo.Session
	}

	mongoManager struct {
		sessions map[string]mongoSession
	}

	DBCall func(*mgo.Collection) error
)

func init() {
	if singleton.sessions != nil {
		return
	}
	singleton = mongoManager{
		sessions: make(map[string]mongoSession),
	}
	if err := CreateSession("monotonic", MonotonicSession, setting.DbHost, setting.DbName, setting.DbUser, setting.DbPass); err != nil {
		log.Fatalf("mongodb connection failed : %v", err)
	}
	fmt.Print("create session success")
}

func CreateSession(mode string, sessionName string, host []string, databaseName string, username string, password string) error {
	mongoSession := mongoSession{
		mongoDialInfo: &mgo.DialInfo{
			Addrs:    host,
			Timeout:  10 * time.Second,
			Database: databaseName,
			Username: username,
			Password: password,
		},
	}
	var err error
	mongoSession.mongoSession, err = mgo.DialWithInfo(mongoSession.mongoDialInfo)
	if err != nil {
		return err
	}
	switch mode {
	case "strong":
		mongoSession.mongoSession.SetMode(mgo.Strong, true)
	case "monotonic":
		mongoSession.mongoSession.SetMode(mgo.Monotonic, true)
	}

	mongoSession.mongoSession.SetSafe(&mgo.Safe{})
	singleton.sessions[sessionName] = mongoSession
	return nil
}

func CopySession(useSession string) (*mgo.Session, error) {
	session := singleton.sessions[useSession]
	if session.mongoSession == nil {
		return nil, errors.New("unable to Locate Session " + useSession)
	}
	mongoSession := session.mongoSession.Copy()
	return mongoSession, nil
}

func GetDB() *mgo.Database{
	m, err := CopySession(MonotonicSession)
	if err != nil {
		log.Fatal("unable to get db")
	}
	return m.DB(setting.DbName)
}
