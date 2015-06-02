package main

import (
	"math/rand"
	"testing"
)

var keys1 [][]byte

func init() {
	rand.Seed(12345)
	keys1 = make([][]byte, 10000)
	for i := 0; i < len(keys1); i++ {
		sz := 5 + rand.Intn(20)
		keys1[i] = make([]byte, sz)
		for j := 0; j < sz; j++ {
			keys1[i][j] = byte(rand.Intn(256))
		}
	}
}

type Map1 map[string]interface{}

func (m Map1) Set(key []byte, value interface{}) {
	m[string(key)] = value
}
func (m Map1) Get(key []byte) interface{} {
	return m[string(key)]
}

func BenchmarkMapSet(b *testing.B) {
	m := make(Map1)
	for i := 0; i < b.N; i++ {
		m.Set(keys1[i%len(keys1)], i)
	}
}

func BenchmarkTrieSet(b *testing.B) {
	t := NewTriew()
	for i := 0; i < b.N; i++ {
		t.Set(keys1[i%len(keys1)], i)
	}
}

func BenchmarkMapGet(b *testing.B) {
	m := make(Map)
	for i := 0; i < b.N; i++ {
		m.Set(keys1[i%len(keys1)], i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get(keys1[i%len(keys1)])
	}
}

func BenchmarkTrieGet(b *testing.B) {
	t := NewTriew()
	for i := 0; i < b.N; i++ {
		t.Set(keys1[i%len(keys1)], i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Get(keys1[i%len(keys1)])
	}
}
