package loop

import "fmt"

// 単純にカウントアップするforループ
func loop_for_1() {
	// for 初期化文; 条件式; 後処理文 { ... }
	for i := 0; i < 5; i++ {
		fmt.Print(i)
	}
}

// 単純にカウントアップするforループ
// 条件式だけでfor文を書くバージョン
func loop_for_2() {
	i := 0
	for i < 5 {
		fmt.Print(i)
		i++
	}
}

// 単純にカウントアップするforループ
// 無限ループのfor文で書くバージョン
func loop_for_3() {
	i := 0
	for {
		fmt.Print(i)
		i++
		if i == 5 {
			break
		}
	}
}

// 単純にカウントアップするforループ
// rangeを使用してfor文で書くバージョン
func loop_for_4() {
	arr := [5]int{0, 1, 2, 3, 4}
	for i := range arr {
		fmt.Print(i)
	}
}

// ここまでに書いた関数を実行する。
func LoopExecute() {
	loop_for_1()
	loop_for_2()
	loop_for_3()
	loop_for_4()
}
