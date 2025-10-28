package main

import (
	"fmt"
	"prepare_bytedance/lib/lru"
)

func main() {
	lru.Put("gg", lru.NewIntValue(10))
	v := lru.Get("gg")
	fmt.Println(v)
	fmt.Println(v.ToInt())

	for i := 0; i < 300; i++ {
		key := fmt.Sprintf("key_%d", i)
		lru.Put(key, lru.NewIntValue(i))
		cap, length := lru.Monitor()
		fmt.Printf("cap: %d, length: %d\n", cap, length)
		if i >= 200 {
			oldKey := fmt.Sprintf("key_%d", i-200)
			value, ok := lru.Get(oldKey).ToInt()
			fmt.Printf("old key %s, exist: %t, value is: %d\n", oldKey, ok, value)
		}
	}
}
