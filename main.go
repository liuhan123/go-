package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main(){
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if (err != nil){
		fmt.Println("redis conn error| ", err)
		return
	}
	fmt.Println("redis conn success\n")

	defer conn.Close()
	_, err = conn.Do("set", "name", "han")
	if (err != nil){
		fmt.Println("set error |", err)
		return
	}

	_, err = conn.Do("hset", "xian", "time", 4)
	if (err != nil){
		fmt.Println("hset error |", err)
		return
	}

	time, errtime := conn.Do("hget", "xian", "time")
	if (errtime != nil){
		fmt.Println("redis get error |", errtime)
		return
	}

	fmt.Printf("time get %s\n", time)

	//设置key过期时间 10
	_, err = conn.Do("set", "jay", "chou", "ex", 10)
	if (err != nil){
		fmt.Println("set expire error |", err)
		return
	}


	_, err = conn.Do("EXPIRE", "name", 10)
	if (err != nil) {
		fmt.Println("expire fail error |", err)
		return
	}

	//查询键的存在
	bexist, errexist := conn.Do("EXISTS", "xian")
	if (errexist != nil){
		fmt.Println("existerror |", err)
		return
	}

	fmt.Println(bexist)

	//删除键
	_, errdel := conn.Do("del", "xian")
	if (errdel != nil){
		fmt.Println("del error |", err)
	}

}
