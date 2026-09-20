# C#開発者のためのGo言語チュートリアル

C#での開発経験がある方向けに、約5時間でGoの基本を学ぶチュートリアルのサンプルコード集です。このリポジトリには、各章のサンプルコードと、総合演習「Todo HTTP API」のコード・テストを収録しています。サンプルはGoの標準ライブラリだけで動作します。

- [最初に実行する](#最初に実行する)
- [学習の進め方](#学習の進め方)
- [章とファイルの対応](#章とファイルの対応)
- [必須演習と解答の扱い](#必須演習と解答の扱い)
- [総合演習のTodo APIを動かす](#総合演習のtodo-apiを動かす)
- [全体の確認と修了の目安](#全体の確認と修了の目安)
- [困ったとき](#困ったとき)

## 最初に実行する

必要なものは **Go 1.22以降**、エディター、ターミナルです。第6章ではGo 1.22で導入されたHTTPのメソッド付きルーティングを使います。

Gitを使う場合は、次のコマンドで入手してサンプルのディレクトリに移動します。Gitを使わない場合は、GitHubの **Code → Download ZIP** から入手して展開できます。

```sh
git clone https://github.com/usairuka/go-tutorial-for-c-sharp-developers.git
cd go-tutorial-for-c-sharp-developers/go-lab
```

すでに入手済みの場合は、このREADMEがあるリポジトリ直下から `cd go-lab` で移動してください。

**以降のGoコマンドは、すべて `go-lab/` 内で実行します。**

```sh
go version
go env GOMOD
go run ./cmd/hello
go test ./...
```

- `go version`：Go 1.22以降であることを確認します。
- `go env GOMOD`：このリポジトリの `go-lab/go.mod` へのパスが表示されます。
- `go run ./cmd/hello`：`Hello, Go!` と表示されれば、最初の実行は成功です。
- `go test ./...`：`internal/task` と `internal/httpapi` が `ok` になれば、収録済みのテストは成功です。`cmd/` 以下の `[no test files]` は失敗ではありません。

このサンプルには [go.mod](go-lab/go.mod) が用意されています。サンプルを使う場合、`mkdir go-lab` や `go mod init` を再実行する必要はありません。module名の `example.com/go-lab` は教材内の識別名で、ドメインの所有や公開は不要です。

```text
go-tutorial-for-c-sharp-developers/
├── README.md
└── go-lab/                           # Goコマンドの実行場所
    ├── go.mod
    ├── cmd/                          # 各章の実行プログラムとAPIの起動処理
    └── internal/
        ├── task/                     # Todoの型・保存処理とテスト
        └── httpapi/                  # HTTP処理とテスト
```

以降で使う `cmd/...` や `internal/...` というパスは、リポジトリ内では `go-lab/cmd/...`、`go-lab/internal/...` に対応します。

## 学習の進め方

1. [最初に実行する](#最初に実行する)の手順で環境を確認し、第0章から順に進めます。
2. [各章に対応するサンプル](#章とファイルの対応)を開いて実行します。出力を予想してから動かし、値や条件を1か所変えて違いを確かめます。
3. 各章の必須演習に取り組みます。収録コードには解答例が含まれるため、自分で実装する場合は[解答の扱い](#必須演習と解答の扱い)を先に確認してください。5分以上詰まったら収録済みの解答例を読み、動かして次へ進みます。
4. 第6章でAPIを起動してテストし、第7章で理解度を確認します。

学習時間の目安は次のとおりです。学習280分に、第3章・第5章の後の休憩各10分を加えて計300分です。

| 章 | 目安 | 確認したいこと |
| --- | --- | --- |
| 0. 環境設定と最初の実行 | 15分 | moduleとpackage、実行・整形・ビルド |
| 1. 型・変数・制御構文・関数 | 35分 | 短い関数と複数の戻り値 |
| 2. slice・map・文字列 | 35分 | 値のコピーとデータの共有 |
| 3. struct・ポインタ・interface・error | 40分 | Todoの型と振る舞い、明示的なエラー処理 |
| 4. package設計・テスト・ジェネリクス | 35分 | packageの分割、境界条件のテスト |
| 5. goroutine・channel・context・排他制御 | 45分 | 結果の収集、キャンセル、共有データの保護 |
| 6. 総合演習：Todo HTTP API | 65分 | HTTP処理と保存処理を分け、JSON APIをテストする |
| 7. 理解度チェック | 10分 | 学んだ内容を自分の言葉で説明し、実行結果と合わせて振り返る |

## 章とファイルの対応

リンクはリポジトリ直下からのパス、コマンドは **`go-lab/` 内** を基準にしています。各プログラムはディレクトリ単位で実行してください。`go run ./...` で全プログラムをまとめて起動することはできません。

| 章・節 | ファイル | 実行・確認 |
| --- | --- | --- |
| 0.2 Hello World | [cmd/hello/main.go](go-lab/cmd/hello/main.go) | `go run ./cmd/hello` |
| 1.1・1.3 型・関数、偶数の合計 | [cmd/basics/main.go](go-lab/cmd/basics/main.go) | `go run ./cmd/basics` |
| 2.1・2.5 slice・map・文字列、filterAtLeast | [cmd/collections/main.go](go-lab/cmd/collections/main.go) | `go run ./cmd/collections` |
| 3.1・3.4 Task、Rename | [internal/task/task.go](go-lab/internal/task/task.go) | `go test ./internal/task -v` |
| 3.2 ポインタとinterface | [cmd/model/main.go](go-lab/cmd/model/main.go) | `go run ./cmd/model` |
| 3.3 defer | [cmd/cleanup/main.go](go-lab/cmd/cleanup/main.go) | `go run ./cmd/cleanup` |
| 4.2・4.4 テーブル駆動テスト、Renameのテスト | [internal/task/task_test.go](go-lab/internal/task/task_test.go) | `go test ./internal/task -v` |
| 4.3 ジェネリクス | [cmd/generics/main.go](go-lab/cmd/generics/main.go) | `go run ./cmd/generics` |
| 5.2・5.5 並行処理とキャンセル、期限切れ | [cmd/concurrency/main.go](go-lab/cmd/concurrency/main.go) | `go run ./cmd/concurrency` |
| 6.2 メモリ上の保存先 | [internal/task/memory.go](go-lab/internal/task/memory.go) | `go test ./internal/task -v` |
| 6.3・6.7 HTTP処理、health endpoint | [internal/httpapi/handler.go](go-lab/internal/httpapi/handler.go) | `go test ./internal/httpapi -v` |
| 6.4 サーバーの起動 | [cmd/api/main.go](go-lab/cmd/api/main.go) | `go run ./cmd/api` |
| 6.6・6.7 HTTPとhealthのテスト | [internal/httpapi/handler_test.go](go-lab/internal/httpapi/handler_test.go) | `go test ./internal/httpapi -v` |
| 6.6 並行アクセスのテスト | [internal/task/memory_test.go](go-lab/internal/task/memory_test.go) | `go test -race ./internal/task` |

`-race` の実行条件は[全体の確認](#全体の確認と修了の目安)を参照してください。第7章は理解度チェックのため、対応する実行プログラムはありません。

## 必須演習と解答の扱い

**このリポジトリは、必須演習の解答例を反映した完成コードです。** 演習で作成する関数・テスト・出力がすでに含まれています。同じ解答を追記すると、関数などが重複する場合があります。

自分で書きながら学ぶ場合は、このリポジトリとは別の作業用ディレクトリで `go mod init example.com/go-lab` を実行し、各章のコードと演習を順に実装すると、収録コードを比較用に残せます。収録コードを直接編集する場合は、対象の実装を書き換えて動作を確かめてください。

<details>
<summary>解答例の収録箇所と確認方法を見る</summary>

コード内の `必須演習の解答例` コメントが目印です。

| 演習 | 収録済みの内容 | 確認方法・期待結果 |
| --- | --- | --- |
| 1.3 偶数の合計 | `cmd/basics/main.go` の `sumEven` と呼び出し | `go run ./cmd/basics` の末尾に `12`、`0` |
| 2.5 条件に合う要素の抽出 | `cmd/collections/main.go` の `filterAtLeast` | 関数のみ収録。下記の呼び出しを追加すると `[30 20]` |
| 3.4・4.4 名前の変更とテスト | `internal/task/task.go` の `Task.Rename`、`task_test.go` の `TestRename` | `go test ./internal/task -run TestRename -v` |
| 5.5 期限切れcontext | `cmd/concurrency/main.go` の `main` 末尾 | `go run ./cmd/concurrency` の末尾に `deadline: true` |
| 6.7 health endpoint | `internal/httpapi/handler.go` の `GET /health`、`handler_test.go` の `TestHealth` | `go test ./internal/httpapi -run TestHealth -v` |

第2章の演習を確認するには、`cmd/collections/main.go` の既存の `main` 関数内に次の1行を追加し、`go run ./cmd/collections` を実行します。

```go
fmt.Println(filterAtLeast([]int{10, 30, 20}, 20))
```

</details>

## 総合演習のTodo APIを動かす

第6章では、`internal/task` がTodoの検証と保存、`internal/httpapi` がHTTP処理、`cmd/api` がそれらを組み合わせたサーバー起動を担当します。

まず、`go-lab/` 内でサーバーを起動します。

```sh
go run ./cmd/api
```

`listening on http://127.0.0.1:8080` と表示されたら、そのターミナルを開いたまま、**別のターミナル**からリクエストを送ります。以下のBashかPowerShellのどちらかを使ってください。

### Bash

```sh
curl -i http://127.0.0.1:8080/tasks
curl -i -H 'Content-Type: application/json' --data '{"title":"Learn Go"}' http://127.0.0.1:8080/tasks
curl -i http://127.0.0.1:8080/tasks
curl -i http://127.0.0.1:8080/health
```

### PowerShell

```powershell
Invoke-RestMethod -Uri 'http://127.0.0.1:8080/tasks'
Invoke-RestMethod -Uri 'http://127.0.0.1:8080/tasks' -Method Post -ContentType 'application/json' -Body '{"title":"Learn Go"}'
Invoke-RestMethod -Uri 'http://127.0.0.1:8080/tasks'
Invoke-RestMethod -Uri 'http://127.0.0.1:8080/health'
```

### 期待する結果

サーバーを起動した直後に上の4つを順に実行した場合、HTTPステータスとレスポンス本文は次のようになります。Bashの `curl -i` はヘッダーも表示し、PowerShellの `Invoke-RestMethod` はJSONをオブジェクトに変換して表示します。

| 操作 | HTTPステータス | レスポンス本文（JSON） |
| --- | --- | --- |
| 一覧取得 `GET /tasks` | 200 | `[]` |
| 追加 `POST /tasks` | 201 | `{"id":1,"title":"Learn Go","done":false}` |
| 追加後の一覧取得 `GET /tasks` | 200 | `[{"id":1,"title":"Learn Go","done":false}]` |
| 稼働確認 `GET /health` | 200 | `{"status":"ok"}` |

空白だけのタイトルが拒否されることも、Bashでは次のコマンドで確認できます。HTTP 400と `title is required` が返り、Todoは追加されません。

```sh
curl -i -H 'Content-Type: application/json' --data '{"title":"   "}' http://127.0.0.1:8080/tasks
```

サーバーの終了は、起動したターミナルで **Ctrl+C** です。Todoはメモリに保存されるため、再起動すると消えます。追加リクエストを繰り返すとTodoとIDが増えるので、表と同じ結果を確認したい場合はサーバーを再起動してください。

## 全体の確認と修了の目安

コードを変更したら `go-lab/` 内で実行します。`go fmt` はソースの書式を整えます。`go build` はC#の `dotnet build`、`go test` は `dotnet test` に近い役割です。

```sh
go fmt ./...
go build ./...
go test ./...
go vet ./...
```

対応環境では、データ競合も確認します。

```sh
go test -race ./...
go run -race ./cmd/concurrency
```

`-race` は対応OS・アーキテクチャとCツールチェーン等が必要です。実行できない環境では、通常の `go test ./...` と `go vet ./...` まで確認してください。HTTPテストは `httptest` を使うため、APIサーバーを起動しなくても実行できます。

学習を終える目安は、次の内容を確認できることです。

- APIでTodoを追加し、一覧を取得できる。
- 不正入力でTodoが保存されず、保存先の予期しないエラーでも内部の詳細がHTTPレスポンスに出ないことをテストで確認できる。
- health endpointと並行アクセスのテストが成功する。
- 値のコピーと共有、エラー処理、並行処理の注意点を説明できる。理解が曖昧な項目は、対応する章のコードを変更して動作を確かめる。

次に進む場合は、発展課題としてTodoの完了処理、保存失敗のテスト拡張、graceful shutdown、DB保存、同時実行数の制限に取り組んでください。発展課題の完成実装は含めていませんが、保存失敗時の基本的な500レスポンスのテストは第6章のコードに収録済みです。

## 困ったとき

| 症状 | 確認・対処 |
| --- | --- |
| `go` コマンドが見つからない | Goをインストールし、ターミナルを開き直して `go version` を確認します。 |
| `go.mod file not found` や `directory prefix ... does not contain main module` | リポジトリ直下から `cd go-lab` で移動し、`go env GOMOD` がこの教材の `go.mod` を指すことを確認します。 |
| コードを追加すると `redeclared` が出る | 解答例がすでに収録されていないか確認します。同じディレクトリに複数の `main` 関数を置かず、章ごとの `cmd/` 配下で作業します。 |
| API起動時に `address already in use` が出る | 先に起動したサーバーが残っていないか確認し、そのターミナルでCtrl+Cを押してから再実行します。 |
| `curl` で接続できない | サーバー側のターミナルに起動エラーがないか確認し、URLを `http://127.0.0.1:8080` に合わせます。 |
| `error obtaining VCS status` が出る | Gitの作業ディレクトリの状態を確認します。一時的にVCS情報の埋め込みを省く場合は `go build -buildvcs=false ./...` を使えます。 |
