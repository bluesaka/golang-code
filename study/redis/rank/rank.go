/**
排行榜设计

对于百万级别以下的排行榜，redis的sorted set不需要消耗太大内存就可实现高效快速的排序
对于千万级别、亿级别的排行榜，有多种方法。
如可以在sorted set的基础上进行分桶，进行桶排序
或者使用脚本每隔一定的时间去同步top n的排名到redis缓存
*/
package rank

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"go-code/study/redis/goredis"
	"log"
	"math/rand"
)

var (
	pageSize int64 = 10

	sortKey       = "zadd-key"
	bucketSortKey = "bucket-sort-key"
)

func SortRaw() {
	redisClient := goredis.GetRedis()
	defer redisClient.Close()

	//initDataWithRow(redisClient)
	//initDataWithPipeline(redisClient)
	//initDataWithMulti(redisClient)

	getTopByPage(redisClient, 1)
}

// 100w数据，预估分几个桶，以及各个桶的score区间
// 10个桶，各个桶的区间为0-99, 100-199, ..., 990-999
func SortBucket() {
	redisClient := goredis.GetRedis()
	defer redisClient.Close()

	//initBucketDataWithMulti(redisClient)
	getBucketTopByPage(redisClient, 2)
	getBucketSelfRank(redisClient, "m617688")
}

func getBucketTopByPage(redisClient *redis.Client, page int64) {
	list, err := redisClient.ZRevRangeByScoreWithScores(context.Background(), bucketSortKey+"9", &redis.ZRangeBy{
		Min:    "0",
		Max:    "+inf",
		Offset: (page - 1) * pageSize,
		Count:  pageSize,
	}).Result()

	if err != nil {
		log.Println("getBucketTopByPage error:", err)
	}

	for _, z := range list {
		log.Printf("member:%s, score=%g\n", z.Member, z.Score)
	}
}

func getBucketSelfRank(redisClient *redis.Client, member string) {
	// get member score
	score := 801
	bucketNum, frontNumSlice := getBucketNum(score)
	var frontRank int64 = 0

	// 判断在那个桶，排名就是在该桶的排名+在这之前的桶的个数
	// 801在第八个桶，前面有一个桶
	for i := range frontNumSlice {
		fr, err := redisClient.ZCard(context.Background(), bucketSortKey+cast.ToString(i)).Result()
		if err != nil {
			log.Println("getBucketSelfRank zCard error:", err)
		}
		frontRank = frontRank + fr
	}

	count, err := redisClient.ZRevRank(context.Background(), bucketSortKey+cast.ToString(bucketNum), member).Result()
	if err != nil {
		log.Println("getBucketSelfRank ZRevRank error:", err)
	}

	log.Println("rank is", frontRank+count)
}

func getTopByPage(redisClient *redis.Client, page int64) {
	list, err := redisClient.ZRevRangeByScoreWithScores(context.Background(), sortKey, &redis.ZRangeBy{
		Min:    "0",
		Max:    "+inf",
		Offset: (page - 1) * pageSize,
		Count:  pageSize,
	}).Result()

	if err != nil {
		log.Println("getTopByPage error:", err)
	}

	//log.Println(list)
	for _, z := range list {
		log.Printf("member:%s, score=%g\n", z.Member, z.Score)
	}
}

func initDataWithRow(redisClient *redis.Client) {
	var z *redis.Z
	var ctx = context.Background()
	for i := 0; i < 10; i++ {
		rand.Seed(int64(i))
		z = &redis.Z{
			Member: "m" + cast.ToString(i),
			Score:  float64(rand.Int31n(1000)),
		}
		log.Println(z.Score)
		if _, err := redisClient.ZAdd(ctx, sortKey, z).Result(); err != nil {
			log.Println("goredis zadd err:", err)
		}
	}
}

// initDataWithMulti 批量初始化sorted set的值
// 每个值占用6byte(视key的长度而定)，10w个约586kb，100w约5.8M，1000w约57M，1亿约572M
func initDataWithMulti(redisClient *redis.Client) {
	var zz []*redis.Z
	for i := 0; i < 100000; i++ {
		rand.Seed(int64(i))
		zz = append(zz, &redis.Z{
			Member: "m" + cast.ToString(i),
			Score:  float64(rand.Int31n(1000)),
		})
	}
	if _, err := redisClient.ZAdd(context.Background(), sortKey, zz...).Result(); err != nil {
		log.Println("goredis zadd err:", err)
	}
}

func initDataWithPipeline(redisClient *redis.Client) {
	var z *redis.Z
	var ctx = context.Background()
	var pipe = redisClient.Pipeline()
	for i := 0; i < 100000; i++ {
		rand.Seed(int64(i))
		z = &redis.Z{
			Member: "m" + cast.ToString(i),
			Score:  float64(rand.Int31n(1000)),
		}
		pipe.ZAdd(ctx, sortKey, z)
	}
	cmders, err := pipe.Exec(ctx)
	if err != nil {
		log.Println("goredis initDataWithPipeline error:", err)
	}
	for _, cmder := range cmders {
		if _, err := cmder.(*redis.IntCmd).Result(); err != nil {
			log.Println("goredis initDataWithPipeline exec error:", err)
		}
	}
}

func initBucketDataWithMulti(redisClient *redis.Client) {
	// 100个bucket桶
	m := make(map[int]string, 100)
	n := make(map[int][]*redis.Z, 100)

	for i := 0; i < 100; i++ {
		m[i] = bucketSortKey + cast.ToString(i)
		n[i] = []*redis.Z{}
	}

	var score float64
	for i := 0; i < 1000000; i++ {
		rand.Seed(int64(i))
		score = float64(rand.Int31n(1000))
		// 10个桶，各个桶的区间为0-99, 100-199, ..., 990-999
		n[int(score/100)] = append(n[int(score/100)], &redis.Z{
			Member: "m" + cast.ToString(i),
			Score:  score,
		})
	}

	for i := 0; i < 10; i++ {
		if len(n[i]) > 0 {
			if _, err := redisClient.ZAdd(context.Background(), m[i], n[i]...).Result(); err != nil {
				log.Println("goredis zadd err:", err)
			}
		}
	}
}

func getBucketNum(score int) (int, []int) {
	var frontBucket []int
	index := score / 100
	max := 9
	for {
		if index < max {
			frontBucket = append(frontBucket, max)
			max--
		}
		break
	}
	return index, frontBucket
}
