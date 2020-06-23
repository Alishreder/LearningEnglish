package main

import (
	. "dictionaryProject/algorithms"
	"fmt"
)


func main() {
	var MyCache = CacheType{
		Cache: make(map[uint64]Word),
	}
	MyCache.AddWordToCache("слово", "word")
	MyCache.AddWordToCache("имя", "name")
	MyCache.AddWordToCache("фамилия", "surname")

	for _, v := range MyCache.Cache {
		fmt.Printf("%+v\n", v)
	}

}
