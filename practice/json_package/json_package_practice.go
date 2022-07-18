/**
json パッケージの encoding について学習
	エンコードは、GO -> JSON。デコードは、JSON -> GO を行う。
		エンコードには Marshal() または NewEncoder().Encode() を使用する
		デコードには Unmarshal() または NewDecoder().Decode() を使用する
	https://zenn.dev/hsaki/articles/go-convert-json-struct
	https://osamu-tech-blog.com/go-encord-decord/
**/

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// GoStruct is 構造体
type GoStruct struct {
	A string
	B string
}

// TagStruct is jsonタグのテスト用構造体
type TagStruct struct {
	C string `json:"first"`
}

func main() {

	// MarshalとEncoderで使用
	structData := GoStruct{A: "aaa", B: "bbb"}

	// ##### Marshal (Struct->JSON を byte型 で変数に格納) ######
	enMaJSON, errEnMa := json.Marshal(structData)
	if errEnMa != nil {
		fmt.Println(errEnMa)
		return
	}
	fmt.Printf("MarshalJSON: %s\n", enMaJSON)
	// #####################################

	// ###### Encoder (Struct->JSON を ファイルに出力) ########
	file, _ := os.Create("encoder_json.txt")
	data := []byte("EncoderJSON: ")
	file.Write(data)
	enErr := json.NewEncoder(file).Encode(structData)
	if enErr != nil {
		fmt.Println(enErr)
		return
	}
	// #####################################

	// UnmarshalとDecoderで使用
	jsonString := `{"A":"testA", "B":"testB"}`

	// ###### Unmarshal (JSON->Struct) ######
	var unmData GoStruct
	unmErr := json.Unmarshal([]byte(jsonString), &unmData) // []byte("string型") で文字列型→byte型に変換している
	if unmErr != nil {
		fmt.Println(unmErr)
		return
	}
	fmt.Printf("UnmarshalStruct: %+v\n", unmData)
	// #####################################

	// ###### Decoder (JSON->Struct) ######
	var decoData GoStruct
	decoErr := json.NewDecoder(strings.NewReader(jsonString)).Decode(&decoData)
	if decoErr != nil {
		fmt.Println(decoErr)
		return
	}
	fmt.Printf("DecoderStruct: %+v\n", decoData)
	// #####################################

	// ##### タグ付与時のエンコード (Struct->JSON にしたとき、キー名がタグを参照するようになる) ######
	tagStruct := TagStruct{C: "testC"}
	tagJSON, tagErr := json.Marshal(tagStruct)
	if tagErr != nil {
		fmt.Println(tagErr)
		return
	}
	fmt.Printf("tagJSON: %s\n", tagJSON)
	// #####################################

	// ###### タグ付与時のデコード (JSON->Struct) #####
	typeAJSON := `{"C":"ccc"}`
	typeBJSON := `{"first":"ccc"}`
	typeCJSON := `{"First":"ccc"}`

	var typeAStruct TagStruct
	var typeBStruct TagStruct
	var typeCStruct TagStruct

	// タグと一致していないキー名は反映されない
	decodeToTagStruct(typeAJSON, &typeAStruct)
	fmt.Printf("tagDecode %+v\n", typeAStruct)
	// タグと完全一致しているキー名は反映される
	decodeToTagStruct(typeBJSON, &typeBStruct)
	fmt.Printf("tagDecode %+v\n", typeBStruct)
	// 頭が大文字でも反映される
	decodeToTagStruct(typeCJSON, &typeCStruct)
	fmt.Printf("tagDecode %+v\n", typeCStruct)
	// #####################################
}

func decodeToTagStruct(j string, s *TagStruct) {
	if err := json.Unmarshal([]byte(j), s); err != nil {
		fmt.Println(err)
		return
	}
}
