package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

type rDB struct {
	conn *redis.Client
}

func New(address string) Repo {

	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	ctx := context.Background()

	rdb.Set(ctx, "key:1", 0, 0)
	rdb.Set(ctx, "key:2", 1, 0)
	rdb.Set(ctx, "key:3", 1, 0)
	num, err := rdb.Get(ctx, "max:1").Result()
	if err != nil {
		log.Println(rdb.Set(ctx, "max:1", 3, 0).Result())
	}
	if num == "" {
		rdb.Set(ctx, "max:1", 3, 0)
	}
	log.Println(rdb.Get(ctx, "max:1").Result())
	return &rDB{
		conn: rdb,
	}
}

func (r *rDB) GetDigit(ctx context.Context, key int) (int, error) {
	keyInt := strconv.Itoa(key)
	return r.conn.Get(ctx, "key:"+keyInt).Int()
}

func (r *rDB) SetDigit(ctx context.Context, key int, value int) error {
	keyInt := strconv.Itoa(key)
	return r.conn.Set(ctx, "key:"+keyInt, value, 0).Err()
}

func (r *rDB) SetMax(ctx context.Context, value int) error {
	return r.conn.Set(ctx, "max:1", value, 0).Err()
}

//GetMax getting key for max Fibonacci digit in redis storage
func (r *rDB) GetMax(ctx context.Context) (int, error) {
	return r.conn.Get(ctx, "max:1").Int()
}

func (r *rDB) GetSliceDigits(ctx context.Context, from, to int) ([]int, error) {
	reSlice := make([]int, 0, to-from+1)
	for ; from <= to; from++ {
		digit, err := r.conn.Get(ctx, "key:"+strconv.Itoa(from)).Int()
		if err != nil {
			return reSlice, err
		}
		reSlice = append(reSlice, digit)
	}
	return reSlice, nil
}

func (r *rDB) SetArray(ctx context.Context, from, to int32, digSlice []int32) error {
	counter := 0
	for ; from <= to; from++ {
		err := r.conn.Set(ctx, "key:"+strconv.Itoa(int(from)), digSlice[counter], 0).Err()
		if err != nil {
			return err
		}
		counter++
	}
	return nil
}
