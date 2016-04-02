package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/garyburd/redigo/redis"
)

var include = [...][]string{{"module", "%s"}, {"module_version", "%s", "version"}, {"type", "%s"}, {"host", "%s"}, {"level", "%s"}}

var conn redis.Conn

func LogsInit() {
	c, err := redis.DialURL("redis://h:p4rmqocfi3snf82b6r7pg1pbnp1@ec2-54-235-164-4.compute-1.amazonaws.com:17979")
	if err != nil {
		panic(err)
	}
	conn = c
}

func Logs(module string) {
	psc := redis.PubSubConn{conn}
	pattern := "babl:log:module:" + module
	if strings.ContainsAny(module, "?*[") {
		psc.PSubscribe(pattern)
	} else {
		psc.Subscribe(pattern)
	}
	for {
		switch v := psc.Receive().(type) {
		case redis.Message:
			printMessage(v.Data)
		case redis.PMessage:
			printMessage(v.Data)
		case error:
			panic(v)
		}
	}
}

func printMessage(data []byte) {
	var d map[string]interface{}
	if err := json.Unmarshal(data, &d); err != nil {
		panic(err)
	}

	remove := []string{"image_id", "@version", "source_host", "container_id", "@timestamp", "SERVICE_TAGS", "created", "image_name"}
	for _, key := range remove {
		delete(d, key)
	}

	msg := []string{}

	for _, inc := range include {
		key := inc[0]
		format := inc[1]
		name := key
		if len(inc) > 2 {
			name = inc[2]
		}

		if d[key] != nil {
			msg = append(msg, fmt.Sprintf(fmt.Sprintf("%s=%s", name, format), d[key]))
			delete(d, key)
		}
	}

	rest := fmt.Sprintf("%+v", d)
	rest = strings.Replace(rest, "map[", "", 1)
	rest = strings.TrimSuffix(rest, "]")
	msg = append(msg, rest)

	fmt.Println(strings.Join(msg, " "))
	// msg := fmt.Sprintf("module=%s version=%s ")

	// if d["type"] && d["type"].(string) == "module" &&

	// fmt.Println(d["foo"].(string))

	// if d["type"].

	// %s\n", d["module"], d["module_version"], d["message"])
}
