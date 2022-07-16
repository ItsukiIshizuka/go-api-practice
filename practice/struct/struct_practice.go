/**
struct内にある`json:"**"`についての解説
    構造体にメタ情報を付与することができるタグ
        ライブラリによってはタグを参照しているものもあるので各ライブラリの使い方を良く見る
    reflectパッケージを利用すると取り出すことができる
        key:"value"の形であればkeyを指定してvalueを参照することもできる
        その場合の注意点として、json:"**" の : と "**" の間にスペースを挟まない
        複数のkeyを持たせることもできる(json1:"***" json2:"***" json3"***")
    https://qiita.com/itkr/items/9b4e8d8c6d574137443c
**/

package main

import (
	"fmt"
	"reflect"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `タグなのでなんでも良い`
	Author string `key1:"test1" key2:"test2" key3:"test3"`
}

func main() {
	book := Book{
		ID:     "1",
		Title:  "testTitle",
		Author: "testAuthor",
	}

	fmt.Println("【構造体の表示】")
	fmt.Printf("%v", book) // 構造体の表示

	fmt.Println("\n")

	fmt.Println("【タグの表示】")
	fmt.Println(reflect.TypeOf(book).Field(0).Tag) //タグの取得(key:value型)
	fmt.Println(reflect.TypeOf(book).Field(1).Tag) //タグの取得(文字列)
	fmt.Println(reflect.TypeOf(book).Field(2).Tag) //タグの取得(複数のkey:value型)

	fmt.Println("\n")

	fmt.Println("【タグのvalue表示】")
	fmt.Println(reflect.TypeOf(book).Field(0).Tag.Get("json")) //タグのvalueを取得
	fmt.Println(reflect.TypeOf(book).Field(2).Tag.Get("key1")) //タグのvalueを取得
	fmt.Println(reflect.TypeOf(book).Field(2).Tag.Get("key2")) //タグのvalueを取得
	fmt.Println(reflect.TypeOf(book).Field(2).Tag.Get("key3")) //タグのvalueを取得
}
