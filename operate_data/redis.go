package operate_data

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var c redis.Conn

func init() {
	dial, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis connection failed", err)
		return
	}
	c = dial
}

func FuncRedis() {
	defer c.Close()
	sgString()
}

func sgString() {
	_, err := c.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}

