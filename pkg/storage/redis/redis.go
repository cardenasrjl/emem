package redis

import (
	"fmt"

	"github.com/go-redis/redis/v7"
	"github.com/vmihailenco/msgpack/v4"
)

type RdsClient struct {
	client *redis.Client
}

//newClient ..
func NewClient(rdClient *redis.Client) *RdsClient {
	return &RdsClient{client: rdClient}
}

//Get ...
func (rd *RdsClient) Get(key string, result interface{}) (found bool) {
	rs, err := rd.client.Get(key).Bytes()
	if err != nil {
		return false
	}
	//marshal the package
	err = msgpack.Unmarshal(rs, result)
	if err != nil {
		return
	}
	return true
}

//Get ...
func (rd *RdsClient) Set(key string, value interface{}) (err error) {
	//cache the value
	value, err = msgpack.Marshal(value)
	if err != nil {
		return
	}

	_, err = rd.client.Set(key, value, 0).Result()
	if err != nil {
		return
	}
	return
}

func (rd *RdsClient) Del(key string) (err error) {
	_, err = rd.client.Del(key).Result()
	return err
}

//GetSetKey checks if key exists, if dones't exists then creates it
func (rd *RdsClient) GetSet(key string, value interface{}, fnValue func() (interface{}, error)) (err error) {
	found := rd.Get(key, value)
	if found {
		return
	}

	//get the value
	value, err = fnValue()
	if err != nil {
		return
	}

	err = rd.Set(key, value)
	return err

}

//key get the key for the provided id
func (rd *RdsClient) Key(id int64) string {
	return fmt.Sprintf("mem:%d", id)
}
