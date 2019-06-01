package dao

import (
	"github.com/mingz2013/mgo-test-pro/datastore"
	"gopkg.in/mgo.v2"
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
}

func NewUserData(index int) (data *dao.UserData) {

	data = &dao.UserData{}

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

	return
}

const DB_NAME = "idlethree"
const USER_DATA_DB_NAME = "user_data"

type UserDataC struct {
	//session *mgo.Session
}

func NewUserDataC() *UserDataC {
	c := &UserDataC{}

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

func (c *UserDataC) Insert(data *UserData) error {
	query := func(c *mgo.Collection) error {
		return c.Insert(data)
	}

	return datastore.HandlerCollection(DB_NAME, USER_DATA_DB_NAME, query)
}

func (c *UserDataC) FindByUserId(userId int) (data *UserData, err error) {
	//err = c.collection().FindId(userId).One(data)

	query := func(c *mgo.Collection) error {
		return c.FindId(userId).One(data)
	}
	err = datastore.HandlerCollection(DB_NAME, USER_DATA_DB_NAME, query)
	return
}
