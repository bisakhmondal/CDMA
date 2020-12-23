package utils

func WalshTable(nodes int) [][]int{
	table := make([][]int, nodes)
	for i := range table{
		table[i] = make([]int, nodes)
	}
	buildTable(table, nodes, 0,0,nodes-1, nodes-1, false)
	return table
}

func buildTable(table [][]int, length, i1, i2, j1, j2 int, bar bool ){
	if length==1{
		table[i1][i2]=1
		return
	}
	if length==2{
		if bar{
			table[i1][i2]=-1
			table[i1][j2]=-1
			table[j1][i2]=-1
			table[j1][j2]=1
		}else{
			table[i1][i2]=1
			table[i1][j2]=1
			table[j1][i2]=1
			table[j1][j2]=-1
		}
		return
	}
	midh := i1 +(j1-i1)/2
	midw := i2 +(j2-i2)/2

	buildTable(table, length/2, i1, i2, midh, midw, bar)
	buildTable(table, length/2, i1, midw+1, midh, j2, bar)
	buildTable(table, length/2, midh+1, i2, j1, midw, bar)
	buildTable(table, length/2, midh+1, midw+1, j1, j2, !bar)
	return
}