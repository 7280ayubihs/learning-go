// Go言語のソースコードは、必ず package 文で始まる。
// ソースコードは「パッケージ」でグルーピングされる。
// hello.go は mainパッケージに属することを宣言している。
package main

// 他パッケージ(fmt)をインポートする。
// つまりは、C言語のinclude文のこと。
import (
	"fmt"
	"learning-go/src/basic"
	"learning-go/src/loop"
)

// func が関数宣言キーワード
// main関数は他の言語と同じく、プログラム実行時に最初に呼ばれる関数。
func main() { // 波カッコはfuncキーワードと同じ行に記載するルール
	fmt.Println("Hello, World!")
	//basic.BasicExecute()
	basic.BasicExecuteDummy()
	loop.LoopExecute()
}
