package Search

import (
	"fmt"
)

const (
	HASHSIZE = 10
)

type hash_table struct {
	elem  []int
	count int
}

var m int = 0

func hash_search(hash_tbl *hash_table, k int) int {
	addr := hash(k)//求哈希地址

	for {
		if hash_tbl->elem[addr] != k {//开放定址法解决冲突
			addr = (addr+1) % m;
			if hash_tbl->elem[addr] == 0 || addr == Hash(k) {
				return -1;
			}
		}
	}

	return addr;
}

func hash_init(hash_tbl *hash_table) {
	m = HASHSIZE
	hash_tbl.count = m
}

func hash(k int) int {
	return k % m //除留余数法
}

func insert(hash_tbl *hash_table, k int) {
	addr := hash(k);
	for {
		if hash_tbl.elem[addr] != NULLKEY {
			addr = (addr+1) % m;//开放定址法
		} else {
			break
		}
	}

	hash_tbl.elem[addr] = k;
}

func show(hash_tbl *hash_table) {
	for i:= 0; i < hash_tbl.count; i++ {
		fmt.Println(hash_tbl.elem[i])
	}
}
