package datastore

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"strings"
)

var MgoSession *mgo.Session

//
//func CreateSession(host string) (session *mgo.Session, err error) {
//	session, err = mgo.Dial(host)
//	if err != nil {
//		return nil, err
//	}
//	session.SetMode(mgo.Monotonic, true)
//	//sessionInstance = session
//	return session, nil
//
//}

func InitSession(host string, username, password string) {

	addrs := strings.Split(host, ",")

	dialInfo := &mgo.DialInfo{
		Addrs:     addrs,
		Direct:    false,
		PoolLimit: 4096,
		Database:  "idlethree",
		Username:  username,
		Password:  password,
	}

	var err error

	MgoSession, err = mgo.DialWithInfo(dialInfo)

	if err != nil {
		fmt.Println(err)
		log.Fatal("err session init..", err)
	} else {
		fmt.Println("mongo init....")
	}
}

func GetSession() *mgo.Session {
	return MgoSession
}

type dbCollection func(collection *mgo.Collection) (err error)

func HandlerCollection(dbName, colName string, query dbCollection) error {
	s := MgoSession.Copy()
	defer s.Close()

	c := s.DB(dbName).C(colName)

	return query(c)
}

const defaultDBHost = "127.0.0.1:27017"

func init() {
	log.Println("init....datastore...")
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultDBHost
	}
	username := os.Getenv("USERNAME")
	if host == "" {
		username = ""
	}
	password := os.Getenv("PASSWORD")
	if host == "" {
		host = ""
	}

	log.Println("host", host, "username", username, "password", password)
	InitSession(host, username, password)
}

//var redisClient *RedisClient
//
//func CreateRedisClient() {
//
//	defaultRedisHost := "redis"
//	defaultRedisPort := "6379"
//
//	host := os.Getenv("REDIS_HOST")
//
//	if host == "" {
//		host = defaultRedisHost
//	}
//
//	port := os.Getenv("REDIS_PORT")
//	if port == "" {
//		port = defaultRedisPort
//	}
//
//	db := "0"
//
//	jsonConf := fmt.Sprintf(`{"host": "%s","port": "%s","db": %s}`, host, port, db)
//
//	log.Println("jsonCOnf", jsonConf)
//
//	redisClient = NewRedisClient(jsonConf)
//}
//
//func GetRedisClient() *RedisClient {
//	return redisClient
//}
//
//const KEY_ID = "key_id"
//
//func GetKeyId() int {
//	id, err := redisClient.Int(redisClient.Do("incr", KEY_ID))
//	if err != nil {
//		id = 1000
//		redisClient.Do("set", KEY_ID, id)
//	}
//	if id < 1000 {
//		id = 1000
//		redisClient.Do("set", KEY_ID, id)
//	}
//
//	return id
//}
