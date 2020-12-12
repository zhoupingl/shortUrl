package main

import (
	"log"
	"os"
	"strconv"
)

type Env struct {
	S         Storage
	ShortHost string // 短链接访问地址
}

func getEnv() *Env {
	addr := os.Getenv("APP_REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	passwd := os.Getenv("APP_REDIS_PASSWD")
	if passwd == "" {
		//passwd =""
	}
	dbs := os.Getenv("APP_REDIS_DB")
	if dbs == "" {
		dbs = "0"
	}

	shorthost := os.Getenv("APP_SHORT_HOST")

	db, err := strconv.Atoi(dbs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("connect to redis (addr: %s password: %s db: %d)", addr, passwd, db)
	r := NewRedisCli(addr, passwd, db)
	return &Env{
		S:         r,
		ShortHost: shorthost,
	}
}
