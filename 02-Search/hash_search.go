package Search

const (
	HASHSIZE = 10
)

// 哈希表的长度
var HashTableLength int

type HashTable struct {
	Elem  []*int
	Count int
}

// 哈希函数
func Hash(data int) int {
	return int(data / L)
}

// 初始化哈希表
func InitHashTable(tbl *HashTable {
	tbl.Count = HashTableLength

	for i := 0; i < HashTableLength; i++ {
		tbl.Elem = append(tbl.Elem, -1)
	}
}

// 插入数据到哈希表
func Insert(tbl *HashTable，data int) {
	adrr := Hash(data)
	for 
}


