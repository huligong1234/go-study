// go-redis.go
//redis download https://github.com/dmajkic/redis/downloads
package main

import (
	"bufio"
	"fmt"
	"github.com/alphazero/Go-Redis"
	"log"
	"os"
)

func main() {

	spec := redis.DefaultSpec().Db(0) //.Password("123456")
	client, e := redis.NewSynchClientWithSpec(spec)
	if e != nil {
		log.Println("failed to create the client", e)
		return
	}

	key := "examples/go/user.name"
	value, e := client.Get(key)
	if e != nil {
		log.Println("error on Get", e)
		return
	}

	if value == nil {
		fmt.Printf("\nHello, don't believe we've met before!\nYour name? ")
		reader := bufio.NewReader(os.Stdin)
		user, _ := reader.ReadString(byte('\n'))
		if len(user) > 1 {
			user = user[0 : len(user)-1]
			value = []byte(user)
			client.Set(key, value)
		} else {
			fmt.Printf("vafanculo!\n")
			return
		}
	}
	fmt.Printf("Hey, %s!\n", fmt.Sprintf("%s", value))
}
