package main

import (
	"bytes"
	"fmt"
	"strconv"
)

//Borad は数独の盤面を示す
//0だと未入力
//1~9だと入力済み
type Borad [9][9]int

func pretty(b Borad) string {
	var buf bytes.Buffer
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			buf.WriteString("+---+---+---+---+---+\n")
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				buf.WriteString("|")
			}
			buf.WriteString(strconv.Itoa(b[i][j]))
			buf.WriteString("|")
		}
		buf.WriteString("\n")

	}
	buf.WriteString("+---+---+---+---+---+\n")
	return buf.String()
}

func duplicate(c [10]int) bool {
	for k, v := range c {
		if k == 0 {
			continue
		}
		if v >= 2 {
			return true
		}
	}
	return false
}

func verify(b Borad) bool {
	//行チェック
	for i := 0; i < 9; i++ {
		//出現回数
		var c [10]int
		for j := 0; j < 9; j++ {
			c[b[i][j]]++
		}
		if duplicate(c) {
			return false
		}

	}
	//列チェック
	for i := 0; i < 9; i++ {
		//出現回数
		var c [10]int
		for j := 0; j < 9; j++ {
			c[b[j][i]]++
		}
		if duplicate(c) {
			return false
		}
	}
	//3*3のチェック
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			var c [10]int
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					c[b[row][col]]++
				}
			}
			if duplicate(c) {
				return false
			}
		}
	}
	return true
}

func solved(b Borad) bool {
	if !verify(b) {
		return false
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

// backtrace
func backtrack(b *Borad) bool {
	if solved(*b) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			//ますがゼロだったら探索開始
			if b[i][j] == 0 {
				for c := 9; c >= 1; c-- {
					b[i][j] = c
					// もし数字がルールに適合するなら
					if verify(*b) {
						//さらに深く探索する
						if backtrack(b) {
							return true
						}
					}
					b[i][j] = 0
				}
				return false
			}
		}
	}
	return false
}

func main() {
	b := Borad{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	fmt.Printf("%+v\n", pretty(b))
	if backtrack(&b){
		fmt.Println(pretty(b))
		fmt.Println("great")
	}
}
