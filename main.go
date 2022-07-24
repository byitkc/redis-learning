package main

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

var redisAddr = "localhost:6379"
var redisOptions = redis.DialPassword("password")

func main() {
	readWriteCache()
}

func mustConnect() (redis.Conn) {
	c, err := redis.Dial("tcp", redisAddr, redisOptions)
	if err != nil {
		panic(err)
	}
	return c
}

func readWriteCache() {
	c := mustConnect()
	defer c.Close()

	c.Do("SET", "mykey", "myvalue2")

	s, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		panic(err)
	}
	log.Println(s)
}

func streams() {
	c := mustConnect()
	
}
