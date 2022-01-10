package mymap

import "fmt"

type Map[K comparable, V any] struct {
	count     int
	flags     uint8
	B         uint8
	noverflow uint16 // 溢出的个数
	hash0     uint32 //哈希种子

	buckets []bmap[K, V]
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{}
}

func (m *Map[K, V]) Put(k K, v V) {
	s := fmt.Sprintf("%v", k)
	hash := hashStr(s)
	index := int(hash[0]) / len(m.buckets)
	t := m.buckets[index]
	m.count++
}

type bmap[K comparable, V any] struct {
	tophash []uint8
}

func (b *bmap[K, V]) put(k K, v V) {

}
