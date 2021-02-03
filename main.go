package main

import "log"

func main() {
	rs := pow(2, 3)
	log.Println(rs)
}

func pow(x, y int) int {
	rs := 1
	for y > 0 {
		// 基数
		if y%2 == 1 {
			rs = rs * x
		}
		x = x * x
		y = y >> 1
	}
	return rs
}
