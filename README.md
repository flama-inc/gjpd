# Go JP Date

- JIS X 0301 に基づく日本語での日付処理
  - ただしスラッシュ区切りを許容する

## Usage

- time.Time -> 元号 -> フォーマット

```go
date, _ := time.Parse(time.DateOnly, "2025-01-01")
g, err := gjpd.GetGengoByTime(date)
if err != nil {
    // ...
}
fmt.Println(g.Name) // "令和"

s := g.Format("ggge年m月d日(aaa)") // 下記参照
fmt.Println(s) // "令和7年1月1日(水曜日)"
```

- 日付の指定と対応する元号への自動切り替え

```go
g := gjpd.GetGengoByName("令和")
t, err := g.SetDate("1", "1", "1") // 令和元年1月1日 -> 平成31年1月1日
if err != nil {
    // ...
}
fmt.Println(t) // 2019-01-01T00:00:00Z
fmt.Println(g.Name) // "平成"
```

- 曜日

```go
date, _ := time.Parse(time.DateOnly, "2025-01-01")
s := gjpd.TimeToWeekday(date, "aaa") // 下記参照
fmt.Println(s) // "水曜日"
```

## レイアウトフォーマット指定子

- Excel準拠
  - <https://learn.microsoft.com/ja-jp/office/client-developer/visio/about-format-pictures>

|指定子|意味|例|
|--|--|--|
|g|元号略称|`R`, `H`|
|gg|元号省略形|`令`, `平`
|ggg|元号|`令和`, `平成`|
|ee|和暦（2桁）|`02`, `10`|
|e|和暦|`2`, `10`|
|mm|月（2桁）|`03`, `12`|
|m|月|`3`, `12`|
|dd|日（2桁）|`07`, `31`|
|d|日|`7`, `31`|
|aaaa|曜日（日本語）|`火曜日`|
|aaa|曜日（日本語省略形）|`火`|
|dddd|曜日（英語）|`Friday`|
|ddd|曜日（英語省略形）|`Fri`|
