# Go言語で作る簡単APIアプリケーション
下の記事に沿ってGoでAPIアプリケーションを作成する<br>
[Goで超簡単API](https://qiita.com/k-penguin-sato/items/8088b69304ee7e8f70be)

## 使用しているツール
- gorilla/mux
  - Go言語で有名なWEBツールキットが `gorilla`
  - その中でルーティング機能を提供しているのが `mux`
  - https://qiita.com/gold-kou/items/99507d33b8f8ddd96e3a

## 参照した記事とは異なる点
- 最初に `go mod init ***` を実行する → そうしないと次の `go get` でエラーが発生する
  - https://teratail.com/questions/324895
- `go get` コマンドはGo1.16以降使用しない(代わりに `go install` を使用する)
  - `go install` を使用する場合は `****@version指定` を必ず付ける
    - 最新で良い場合は `@latest` を付ける
  - https://qiita.com/eihigh/items/9fe52804610a8c4b7e41
  - **※ここまで調べたけど、結局 `go install` コマンドではエラーが出てインストールできなかった
    ```
      % go install github.com/gorilla/mux@latest
      package github.com/gorilla/mux is not a main package
    ```

## その他
- postmanの使い方
  - デスクトップアプリから利用しないとlocalhostにアクセスできない
  - https://www.tairaengineer-note.com/postman-how-to-use/
