package mongodb

import (
	"github.com/globalsign/mgo"
	"log"
	"time"
)

var session *mgo.Session
var database *mgo.Database

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{"localhost:27017"},
		Timeout:   time.Second * 3,
		Database:  "community",
		PoolLimit: 4096,
	}
	var err error
	session, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}

	//Strong
	//session 的读写一直向主服务器发起并使用一个唯一的连接，因此所有的读写操作完全的一致。
	//Monotonic
	//session 的读操作开始是向其他服务器发起（且通过一个唯一的连接），只要出现了一次写操作，session 的连接就会切换至主服务器。由此可见此模式下，能够分散一些读操作到其他服务器，但是读操作不一定能够获得最新的数据。
	//Eventual
	//session 的读操作会向任意的其他服务器发起，多次读操作并不一定使用相同的连接，也就是读操作不一定有序。session 的写操作总是向主服务器发起，但是可能使用不同的连接，也就是写操作也不一定有序。
	session.SetMode(mgo.Monotonic, true)
	database = session.DB("community")
}

func GetMgo() *mgo.Session {
	return session
}

func GetDatabase() *mgo.Database {
	return database
}

func GetErrNotFound() error {
	return mgo.ErrNotFound
}

type sessionPool struct {
	session *mgo.Session
}

func (sp *sessionPool) C(name string) *mgo.Collection {
	return sp.session.DB("community").C(name)
}

func NewSessionPool() *sessionPool {
	return &sessionPool{
		session: session.Copy(),
	}
}

func (sp *sessionPool) Close() {
	sp.session.Close()
}

type Topic struct {
	// bson
	ID      uint64 `bson:"_id"`
	GuildID uint64 `bson:"guild_id"`
}

type User struct {
	ID      int64 `bson:"_id"`
	GuildID int64 `bson:"guild_id"`
}

func Query() {
	var topic Topic
	collection := GetDatabase().C("topics")
	if err := collection.FindId(18).One(&topic); err != nil {
		log.Println(err)
	}

	//if err := collection.Find(bson.M{"guild_id": 1}).One(&topic); err != nil {
	//	log.Println(err)
	//}

	log.Printf("%+v\n", topic)
}

func Query2() {
	var topic Topic
	sp := NewSessionPool()
	defer sp.Close()
	collection := sp.C("topics")
	if err := collection.FindId(18).One(&topic); err != nil {
		log.Println(err)
	}

	//if err := collection.Find(bson.M{"guild_id": 1}).One(&topic); err != nil {
	//	log.Println(err)
	//}

	log.Printf("%+v\n", topic)
}
