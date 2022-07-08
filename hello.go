// Go言語のソースコードは、必ず package 文で始まる。
// ソースコードは「パッケージ」でグルーピングされる。
// hello.go は mainパッケージに属することを宣言している。
package main

// 他パッケージ(fmt)をインポートする。
// つまりは、C言語のinclude文のこと。
import (
	"fmt"
	"unicode/utf8"
)

// func が関数宣言キーワード
// main関数は他の言語と同じく、プログラム実行時に最初に呼ばれる関数。
func main() { // 波カッコはfuncキーワードと同じ行に記載するルール
	// お約束のハローワールド。
	fmt.Println("Hello, World!")

	// --- 論理値型 ---
	var b bool = true
	fmt.Println(b)

	// --- 数値型 ---
	// 整数、小数、複素数のそれぞれが、サイズ毎に数値型がある。
	// 例えば、uint32 は符号なし32bit整数 となる。
	var i32 int32 = 123
	// それぞれの型に互換性はなく、異なる数値型を代入する時は変換が必要。
	var i64 int64 = int64(i32)
	fmt.Println(i64)

	// --- 文字列型 ---
	var str string = "abc"
	// 文字列は +演算子、+=演算子で結合することができる。
	str = str + "d"
	str += "e"
	fmt.Println(str)
	// 文字列の長さを取得するには、len()関数を使用する。
	fmt.Println(len(str))
	// ただし、len()関数で日本語（正確にはマルチバイト文字）の長さを取得する場合は注意が必要
	// len()関数が取得するのはバイト数なので、日本語を含む文字列の長さを取得したい場合は、
	// unicode/utf8パッケージのRuneCountInString()関数を使用する。
	var str2 string = "山田太郎"
	fmt.Println(utf8.RuneCountInString(str2))

	// --- 型宣言 ---
	// 型宣言を行うことで、新たな型を作成することができる。
	// つまりは、C言語のtypedef宣言のこと。
	type myInt32 int32
	var mi32 myInt32 = 123
	fmt.Println(mi32)
	// 構造体を作る場合にも使用される。
	type myStruct struct {
		a int32
		b int32
	}

	// --- 変換 ---
	// Go言語では、型チェックが厳しく、異なる型での代入や比較ができません。
	// 必要がある場合は、明示的に変換を行う必要がある。
	// 変換は「変換先の型(値) 」で行う。
	/* var i32 int32 = 123 */
	var ui32 uint32 = uint32(i32)   // int32  から uint32  への変換
	var f32 float32 = float32(ui32) // uint32 から float32 への変換
	var str3 string = string(i32)   // int32  から string  への変換
	fmt.Println(ui32)
	fmt.Println(f32)
	fmt.Println(str3)

	// --- 演算子 ---
	// TODO
}

/*
セミコロンは原則として不要（※正確には省略可能）
*/
