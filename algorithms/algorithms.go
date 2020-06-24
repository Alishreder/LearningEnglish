package algorithms

import "sync"

/*

It was in main

var MyCache = CacheType{
	Cache: make(map[uint64]Word),
}
MyCache.AddWordToCache("слово", "word")
MyCache.AddWordToCache("имя", "name")
MyCache.AddWordToCache("фамилия", "surname")

for _, v := range MyCache.Cache {
fmt.Printf("%+v\n", v)
}

*/

type Word struct {
	Rus string `json:"rus" bson:"rus"`
	Eng string `json:"eng" bson:"eng"`
	Id  uint64 `json:"id" bson:"_id"`
}

type CacheType struct {
	Cache map[uint64]Word
	sync.RWMutex
	sync.WaitGroup
}

func (c *CacheType) SetId() uint64 {
	return uint64(len(c.Cache) + 1)
}

func (c *CacheType) AddWordToCache(rus, eng string) {
	newId := c.SetId()
	c.Cache[c.SetId()] = Word{
		Rus: rus,
		Eng: eng,
		Id:  newId,
	}
}

func (c *CacheType) FromEngToRus(rus, eng string) {
	// if ok, v := c.Cache[]; ok {
	//
	// }
}
