package main

import (
	"fmt"

	"github.com/hiroto60/mymodule/mathutils"
)

func main() {
	fmt.Println("main関数が開始されました")
	result1 := mathutils.Add(5, 3)
	result2 := mathutils.Subtract(10, 7)
	fmt.Println("5 + 3 =", result1)
	fmt.Println("10 - 7 =", result2)
	fmt.Println("main関数が終了しました")
}
