package hashmap

import (
	"fmt"

	"github.com/hexnaught/go-hashmap-exp/pkg/linkedlist"
)

type HashMap struct {
	buckets []*linkedlist.LinkedList
	size    int
}

func (hm *HashMap) Print() {
	fmt.Printf("\t%+v\n", hm)
	for _, ll := range hm.buckets {
		fmt.Printf("\t\t%+v\n", ll)
		ll.Print()
	}
}

// Has returns if the value already exists in the instance of HashMap
func (hm *HashMap) Has(key string) bool {
	keyHashCode := hm.getHashCode(key)
	hasKey, _ := hm.buckets[keyHashCode].Search(key)
	return hasKey
}

// Get retrieves the stored value for the given key in the HashMap
func (hm *HashMap) Get(key string) interface{} {
	keyHashCode := hm.getHashCode(key)
	return hm.buckets[keyHashCode].Get(key)
}

func (hm *HashMap) Set(key string, value interface{}) {
	keyHashCode := hm.getHashCode(key)
	if hm.Has(key) {
		hm.buckets[keyHashCode].Set(key, value)
		return
	}
	hm.buckets[keyHashCode].Insert(key, value)
}

func (hm *HashMap) Remove(key string) {
	keyHashCode := hm.getHashCode(key)
	hm.buckets[keyHashCode].Remove(key)
}

func (hm *HashMap) getHashCode(key string) int {
	return hash(key) % hm.size
}

func hash(s string) int {
	hashCode := 0
	for _, c := range s {
		hashCode += int(c)
	}
	return hashCode
}

func New(size int) *HashMap {
	// TODO: Force init to a 'closest size to the power of 2
	// TODO: ability to resize hashmap capactiy. Lazy rehash? Full rehash?
	hashMap := &HashMap{
		size:    size,
		buckets: make([]*linkedlist.LinkedList, size),
	}

	for i := 0; i < size; i++ {
		hashMap.buckets[i] = &linkedlist.LinkedList{}
	}

	return hashMap
}
