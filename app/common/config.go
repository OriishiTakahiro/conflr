package common

const (
	CONFL_CONTENT_TYPE = "page"
	EXPAND             = "space,history,body.view,metadata.labels"
)

var (
	LABEL_TERMS = map[string][]string{
		"調査":   {"調査", "サーベイ", "レポート", "まとめ", "検証"},
		"試験":   {"試験", "テスト", "検証"},
		"MTG":  {"議事録", "会議", "MTG", "ミーティング", "話し合い"},
		"運用":   {"運用方法", "オペレーション", "使用方法", "進め方", "手順"},
		"Tips": {"Tips", "tips", "幸せになれる", "ノウハウ"},
		"標準":   {"標準", "スタンダード", "規則", "規約", "基準"},
		"仕様":   {"仕様", "エラーコード", "設計"},
		"ホーム":  {"ホーム"},
	}
)
