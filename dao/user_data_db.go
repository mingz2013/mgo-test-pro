package dao

import (
	"fmt"
	"mgo-test-pro/datastore"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"k8s.io/apimachinery/pkg/util/rand"
	"log"
	"os"
	"strconv"
	"time"
)

/*

   # {
   #     _id: <userId>,
   #     stoken: asdfsadfas,   # stoken, http接口使用
   #     serverId: 1,          # 玩家在哪个服
   #     mcExpireAt: Date(),  # 月卡到期时间
   #     mcDailyRewardAtDay: "2019-11-12",  # 发送月卡每日奖励的日期
   #     wcExpireAt: Date(),  # 周卡到期时间
   #     wcExpireNotice: 0,   # 周卡到期提醒是否发送过
   #     newbie: {doneList: [k1,k2], currProgress: k}   # 新手引导
   #     enterGameAt: Date(),  # 进入游戏时间，从first_enter_hall开始算
   #     onlineAt: Data(),     # 上线时间
   #     offlineAt: Date(),    # 离线时间
   #     exp: 123,             # 经验
   #     level: 12,            # 等级
   #     vipExp: 999           # vip经验
   #     vipLevel: 5,          # vip等级
   #     fightPower:           # 玩家的战斗力
   #     activeTimestamp       # 玩家处于活跃状态的超时时间戳(例如，新玩家第一次登陆，这个时间会设置为第二天0时，其他玩家当每日活跃度达成>=50后，会设置这个值为第三天0时)
   #     achievement:          # 玩家的成就点
   #     merit:                # 玩家的功勋点
   #     militaryRank          # 军衔
   #     package:{}            # 背包购买次数   package.hero
   #     originSharerInfo: {}  # 源头是否别人邀请而来，不是别人邀请来的时该字段不存在。 {sharerId: xx, funcType: xx}
   # }

*/

type objectType map[string]interface{}

type UserData struct {
	UserId             int        `json:"_id" bson:"_id"`
	SToken             string     `json:"stoken" bson:"stoken"`
	ServerId           int        `json:"serverId" bson:"serverId"`
	McExpireAt         string     `json:"mcExpireAt" bson:"mcExpireAt"`
	McDailyRewardAtDay string     `json:"mcDailyRewardAtDay" bson:"mcDailyRewardAtDay"`
	WcExpireAt         string     `json:"wcExpireAt" bson:"wcExpireAt"`
	WcExpireNotice     bool       `json:"wcExpireNotice" bson:"wcExpireNotice"`
	Newbie             objectType `json:"newbie" bson:"newbie"`
	EnterGameAt        string     `json:"enterGameAt" bson:"enterGameAt"`
	OnlineAt           string     `json:"onlineAt" bson:"onlineAt"`
	OfflineAt          string     `json:"offlineAt" bson:"offlineAt"`
	Exp                int        `json:"exp" bson:"exp"`
	Level              int        `json:"level" bson:"level"`
	VipExp             int        `json:"vipExp" bson:"vipExp"`
	VipLevel           int        `json:"vipLevel" bson:"vipLevel"`
	FightPower         int        `json:"fightPower" bson:"fightPower"`
	ActiveTimestamp    int        `json:"activeTimestamp" bson:"activeTimestamp"`
	Achievement        int        `json:"achievement" bson:"achievement"`
	Merit              string     `json:"merit" bson:"merit"`
	MilitaryRank       string     `json:"militaryRank" bson:"militaryRank"`
	Package            objectType `json:"package" bson:"package"`
	OriginSharerInfo   objectType `json:"originSharerInfo" bson:"originSharerInfo"`
	TestData           objectType `json:"testData" bson:"testData"`
}

type FindData struct {
	SToken             string     `json:"stoken" bson:"stoken"`
	ServerId           int        `json:"serverId" bson:"serverId"`
	McExpireAt         string     `json:"mcExpireAt" bson:"mcExpireAt"`
	McDailyRewardAtDay string     `json:"mcDailyRewardAtDay" bson:"mcDailyRewardAtDay"`
	WcExpireAt         string     `json:"wcExpireAt" bson:"wcExpireAt"`
	WcExpireNotice     bool       `json:"wcExpireNotice" bson:"wcExpireNotice"`
	Newbie             objectType `json:"newbie" bson:"newbie"`
	EnterGameAt        string     `json:"enterGameAt" bson:"enterGameAt"`
	OnlineAt           string     `json:"onlineAt" bson:"onlineAt"`
	OfflineAt          string     `json:"offlineAt" bson:"offlineAt"`
	Exp                int        `json:"exp" bson:"exp"`
	Level              int        `json:"level" bson:"level"`
	VipExp             int        `json:"vipExp" bson:"vipExp"`
	VipLevel           int        `json:"vipLevel" bson:"vipLevel"`
	FightPower         int        `json:"fightPower" bson:"fightPower"`
	ActiveTimestamp    int        `json:"activeTimestamp" bson:"activeTimestamp"`
	Achievement        int        `json:"achievement" bson:"achievement"`
	Merit              string     `json:"merit" bson:"merit"`
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < l; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func RandString() string {
	return GetRandomString(rand.IntnRange(1, 50))
}

func RandSlice() (l []string) {
	num := rand.IntnRange(1, 20)

	for i := 0; i < num; i++ {
		l = append(l, RandString())
	}

	return
}

func NewUserData(index int) (data *UserData) {

	data = &UserData{}

	startUserId := 10000
	var err error
	s := os.Getenv("START_USER_ID")
	if s != "" {
		startUserId, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("error START_USER_ID", err)
		}
	}

	//data.UserId = string(datastore.GetKeyId())
	data.UserId = index + startUserId
	data.Achievement = rand.Intn(999999)
	data.ActiveTimestamp = 1559364073
	data.EnterGameAt = ""
	data.Exp = rand.Intn(999999)
	data.FightPower = rand.Intn(999999)
	data.Level = rand.Intn(999999)
	data.McDailyRewardAtDay = ""
	data.McExpireAt = ""
	data.WcExpireNotice = false
	data.FightPower = rand.Intn(999999)
	data.Merit = ""
	data.MilitaryRank = ""
	data.Newbie = map[string]interface{}{
		"currProgress": "GuideEnd",
		"doneList":     []string{"GuidePubNormal", "GuideEnd", "GuideEquipmentUP", "GuideEnd", "GuideStory", "GuideEnd", "GuideFuseHero", "GuideEnd", "GuideLord", "GuideEnd", "GuideFigtingFailure"},
	}
	data.OfflineAt = ""
	data.OnlineAt = ""
	data.Package = map[string]interface{}{
		"aaa": "aaa",
	}

	data.ServerId = rand.Intn(999999)
	data.SToken = "aaaaaaaaaa"
	data.VipLevel = rand.Intn(999999)
	data.VipExp = rand.Intn(999999)

	//data.TestData = map[string]interface{}{
	//	"testList": RandSlice(),
	//	"hahaha":   RandSlice(),
	//}
	//
	//num := rand.Intn(100)
	//for i := 0; i < num; i++ {
	//	data.TestData[RandString()] = RandSlice()
	//}
	//
	//num = rand.Intn(100)
	//for i := 0; i < num; i++ {
	//	data.Package[RandString()] = RandSlice()
	//}
	//
	//num = rand.Intn(100)
	//for i := 0; i < num; i++ {
	//	data.Newbie[RandString()] = RandSlice()
	//}

	return
}

var DB_NAME, USER_DATA_COLLECTION_NAME string

func init() {
	DB_NAME = "idlethree"
	USER_DATA_COLLECTION_NAME = "user_data"

	s := os.Getenv("DB_NAME")
	if s == "" {
	} else {
		DB_NAME = s
	}

	s = os.Getenv("USER_DATA_COLLECTION_NAME")
	if s == "" {
	} else {
		USER_DATA_COLLECTION_NAME = s
	}

}

type UserDataC struct {
	//session *mgo.Session
	collectionIndex int
}

func NewUserDataC(collectionIndex int) *UserDataC {
	c := &UserDataC{
		collectionIndex: collectionIndex,
	}

	//var err error
	//c.session, err = datastore.CreateSession(host)
	//if err != nil {
	//	fmt.Println("NewUserDataC...", err)
	//}

	return c
}

//func (c *UserDataC) Close() {
//	c.session.Close()
//}

//func (c *UserDataC) collection() *mgo.Collection {
//	return c.session.DB(DB_NAME).C(USER_DATA_DB_NAME)
//}

func (c *UserDataC) DoQuery(query func(c *mgo.Collection) (err error)) (err error) {

	name := USER_DATA_COLLECTION_NAME + "_" + strconv.Itoa(c.collectionIndex)
	log.Println("name", name, "dbname", DB_NAME)
	err = datastore.HandlerCollection(DB_NAME, name, query)
	return
}

func (c *UserDataC) Insert(data *UserData) (err error) {
	query := func(c *mgo.Collection) error {
		return c.Insert(data)
	}

	err = c.DoQuery(query)
	return
}

func (c *UserDataC) FindByUserId(userId int) (data *FindData, err error) {
	//err = c.collection().FindId(userId).One(data)

	data = &FindData{}

	query := func(c *mgo.Collection) error {
		return c.FindId(userId).One(data)
	}
	err = c.DoQuery(query)
	return
}

func (c *UserDataC) TestUpdate(userId int, data *UserData) (err error) {

	//data := NewUserData(userId)

	query := func(c *mgo.Collection) error {
		_, err := c.Upsert(bson.M{"userId": userId}, data)
		return err
	}
	err = c.DoQuery(query)

	return
}
