# hello-world バージョン差分メモ

## 全体概要

| バージョン | 主な状態 |
| --- | --- |
| v1 | `main` で直接 `"Hello, world"` を表示 |
| v2 | `Hello()` 関数を導入し、最初のテストを追加 |
| v3 | `Hello(name string)` に変更して名前付き挨拶へ対応 |
| v4 | 挨拶プレフィックスを定数化 |
| v5 | 空文字時に `"World"` を使うデフォルト処理を追加 |
| v6 | 言語指定 (`English/Spanish/French`) に対応 |
| v7 | 言語分岐を `if` から `switch` ベースに整理 |
| v8 | プレフィックス選択を `greetingPrefix` 関数へ抽出 |

## vごとの差分（前バージョン比）

### v1
- 初期実装。
- `fmt.Println("Hello, world")` を `main` で直接実行。
- テストなし。

### v2（v1 から）
- `Hello() string` 関数を追加。
- `main` は `Hello()` の戻り値を表示する形に変更。
- `hello_test.go` を追加し、`TestHello` で戻り値を検証。

### v3（v2 から）
- `Hello()` を `Hello(name string)` に変更し、引数名を埋め込んだ挨拶へ変更。
- テストも `Hello("Chris")` を検証する形に更新。

### v4（v3 から）
- `"Hello, "` を `englishHelloPrefix` 定数として切り出し。
- 文字列連結ロジックは同じで、可読性・保守性を向上。

### v5（v4 から）
- `name == ""` のとき `name = "World"` を設定するデフォルト処理を追加。
- テストを `t.Run` で整理し、空文字ケースを追加。
- `assertCorrectMessage` ヘルパーを導入して重複を削減。

### v6（v5 から）
- `Hello(name, language string)` にシグネチャ変更。
- `Spanish` / `French` 定数と各言語プレフィックスを追加。
- 言語に応じて `if` 分岐で `"Hola, "` / `"Bonjour, "` を返す実装を追加。
- テストにスペイン語・フランス語ケースを追加。

### v7（v6 から）
- 言語分岐を連続 `if` から `switch` へ変更。
- `prefix` 変数を英語で初期化し、`case` で上書きする構造に整理。
- 挙動（返却文字列）は v6 と同等。

### v8（v7 から）
- 言語ごとのプレフィックス決定を `greetingPrefix(language string)` へ抽出。
- `Hello` は「名前デフォルト設定 + `greetingPrefix` 呼び出し」の責務に分離。
- `greetingPrefix` は `switch` + `default` で英語へフォールバック。
- テストケース構成は v7 と同様（文言のみ一部調整）。
