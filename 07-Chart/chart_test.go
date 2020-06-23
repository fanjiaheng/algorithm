package Chart

import (
	"fmt"
	"testing"
)

func TestCreateMgraph(T *testing.T) {
	m := &MGraph{}
	CreateMgraph(m)
	fmt.Println(m.Node)
	fmt.Println(m.Edge)
}
