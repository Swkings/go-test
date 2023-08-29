package test

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"testing"

	goredis "github.com/go-redis/redis/v8"
	"github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson/v4"
	ZRedis "github.com/zeromicro/go-zero/core/stores/redis"
)

// Name - student name
type Name struct {
	First  string `json:"first,omitempty"`
	Middle string `json:"middle,omitempty"`
	Last   string `json:"last,omitempty"`
}

// Student - student object
type Student struct {
	Name Name `json:"name,omitempty"`
	Rank int  `json:"rank,omitempty"`
}

func Example_JSONSet(rh *rejson.Handler) {

	student := Student{
		Name: Name{
			"Mark",
			"S",
			"Pronto",
		},
		Rank: 1,
	}
	res, err := rh.JSONSet("TestDirB:student", ".", student)
	if err != nil {
		log.Fatalf("Failed to JSONSet: %v", err)
		return
	}

	if res.(string) == "OK" {
		fmt.Printf("Success: %s\n", res)
	} else {
		fmt.Println("Failed to Set: ")
	}

	studentJSON, err := redis.Bytes(rh.JSONGet("student", "."))
	if err != nil {
		log.Fatalf("Failed to JSONGet")
		return
	}

	readStudent := Student{}
	err = json.Unmarshal(studentJSON, &readStudent)
	if err != nil {
		log.Fatalf("Failed to JSON Unmarshal")
		return
	}

	fmt.Printf("Student read from redis : %#v\n", readStudent)
}

func TestReJson(t *testing.T) {
	var addr = flag.String("Server", "localhost:6379", "Redis server address")

	rh := rejson.NewReJSONHandler()
	flag.Parse()

	// // Redigo Client
	// conn, err := redis.Dial("tcp", *addr)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to redis-server @ %s", *addr)
	// }
	// defer func() {
	// 	_, err = conn.Do("FLUSHALL")
	// 	err = conn.Close()
	// 	if err != nil {
	// 		log.Fatalf("Failed to communicate to redis-server @ %v", err)
	// 	}
	// }()
	// rh.SetRedigoClient(conn)
	// fmt.Println("Executing Example_JSONSET for Redigo Client")
	// Example_JSONSet(rh)

	// GoRedis Client
	cli := goredis.NewClient(&goredis.Options{Addr: *addr, Password: "redis"})
	defer func() {
		// if err := cli.FlushAll(context.Background()).Err(); err != nil {
		// 	log.Fatalf("goredis - failed to flush: %v", err)
		// }
		if err := cli.Close(); err != nil {
			log.Fatalf("goredis - failed to communicate to redis-server: %v", err)
		}
	}()
	rh.SetGoRedisClient(cli)
	fmt.Println("\nExecuting Example_JSONSET for GoRedis Client")
	Example_JSONSet(rh)
}

func TestRedis(t *testing.T) {
	RedisClient := ZRedis.New("127.0.0.1:6379", func(r *ZRedis.Redis) {
		r.Pass = "redis"
	})
	err := RedisClient.Set("TestDirA:DirB:key1", "v1")
	fmt.Println(err)
	RedisClient.Set("TestDirA:DirC:key1", "v1")
	RedisClient.Set("key1", "v1")
}
