package main

import (
	"flag"
	"fmt"

	"github.com/33cn/chain33/common/db"
	"github.com/33cn/chain33/types"
)

var dir = flag.String("dir", "", "data dir of dump db")
var name = flag.String("name", "", "db  name of dump db")

func main() {
	flag.Parse()
	leveldb, err := db.NewGoLevelDB(*name, *dir, 128)
	if err != nil {
		panic(err)
	}
	defer leveldb.Close()
	it := leveldb.Iterator(nil, types.EmptyValue, false)
	for it.Rewind(); it.Valid(); it.Next() {
		fmt.Printf("%s -> %s\n", string(it.Key()), string(it.Value()))
	}
	it.Close()
}
