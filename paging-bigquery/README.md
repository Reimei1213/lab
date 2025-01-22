# BigQuery のページネーション実装サンプル
- `bigquery_emulator.go`
  - 以下のリポジトリを利用して BigQuery に対するページネーションの実装を検証
  - ただ、2025/01/23 現在保存しているデータサイズちょうどを pageSize に指定しない場合は `googleapi: Error 404: job request-20250122-58368391345000-0001 is not found, notFound` というエラーが発生する
  - https://github.com/goccy/bigquery-emulator
- `real_bigquery.go`
  - 実際の BigQuery に対するページネーションの実装
- `pager/pager.go`
  - ページネーションの実装
