package utils

import "testing"

func TestWalshTable(t *testing.T) {
	table := WalshTable(8)
	tt := make([]int, len(table[0]))
	copy(tt, table[0])
	tt[0]=5
	t.Log(table)
	t.Log(tt)
}
