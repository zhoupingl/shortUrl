package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Env struct {
	S         Storage
	ShortHost string // 短链接访问地址
}

const NULL = "null"

func getEnv() *Env {
	addr := os.Getenv("APP_REDIS_ADDR")
	if addr == "" {
		addr = "r-bp1ecyvm2s303ugmh9.redis.rds.aliyuncs.com:6379"
	} else if addr == NULL {
		addr = ""
	}

	passwd := os.Getenv("APP_REDIS_PASSWD")
	if passwd == "" {
		passwd = "CKVm4j6YL3t5"
	} else if passwd == NULL {
		passwd = ""
	}

	dbs := os.Getenv("APP_REDIS_DB")
	if dbs == "" {
		dbs = "5"
	} else if dbs == NULL {
		dbs = ""
	}
	db, err := strconv.Atoi(dbs)
	if err != nil {
		log.Fatal(err)
	}

	shorthost := os.Getenv("APP_SHORT_HOST")
	if shorthost == "" {
		shorthost = "http://url.jzwp.cn"
	} else if shorthost == NULL {
		shorthost = ""
	}


	log.Printf("connect to redis (addr: %s password: %s db: %d)", addr, passwd, db)
	r := NewRedisCli(addr, passwd, db)
	return &Env{
		S:         r,
		ShortHost: strings.TrimRight(shorthost, "/"),
	}
}
