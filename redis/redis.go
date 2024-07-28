package redisaction

import (
	"context"
	"strings"
	"time"

	redis "github.com/redis/go-redis/v9"
)

var c *redis.Client

func init() {
	c = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
}

const (
	ONEDAY = 86400
)

func ArtleVote(user string, article string) error {
	ctx := context.Background()
	cutoff := time.Now().Unix() - 7*ONEDAY
	articleTime, err := c.ZScore(ctx, "time:", article).Result()
	if err != nil {
		return err
	}
	if articleTime < float64(cutoff) {
		return nil
	}
	data := strings.Split(article, ":")
	articleId := data[len(data)-1]
	num, err := c.SAdd(ctx, "voted:"+articleId, user).Result()
	if err != nil {
		return err
	}
	if num > 0 {
		_, err = c.ZIncrBy(ctx, "score:", 432.0, article).Result()
		if err != nil {
			return err
		}
		_, err := c.HIncrBy(ctx, article, "votes", 1).Result()
		if err != nil {
			return err
		}
	}
	return nil
}

func TestPipeLine() {
	ctx := context.Background()
	pipe := c.Pipeline()
	pipe.Set(ctx, "test", 10, 60).Result()
	pipe.Get(ctx, "test").Result()
	// res, err := pipe.Exec(ctx)
	// fmt.Println(res, err)
	// data, err := c.Get(ctx, "test").Result()
	// fmt.Println(data, err)
	time.Sleep(10 * time.Second)
}
