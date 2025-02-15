package main

import (
	"context"
	"fmt"
	"path/filepath"

	"cloud.google.com/go/bigquery"
	"github.com/goccy/bigquery-emulator/server"
	"google.golang.org/api/option"

	"github.com/Reimei1213/lab/paging-bigquery/pager"
)

func main() {
	const (
		projectID = "test"
		tableName = "test.dataset1.table_a"
	)

	// bigquery-emulator サーバを起動

	ctx := context.Background()
	bqServer, err := server.New(server.TempStorage)
	if err != nil {
		panic(err)
	}
	if err := bqServer.SetProject(projectID); err != nil {
		panic(err)
	}
	if err := bqServer.Load(server.YAMLSource(filepath.Join("testdata", "data.yaml"))); err != nil {
		panic(err)
	}

	testServer := bqServer.TestServer()
	defer testServer.Close()

	// ----------------------------

	cli, err := bigquery.NewClient(
		ctx,
		projectID,
		option.WithEndpoint(testServer.URL),
		option.WithoutAuthentication(),
	)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	p := pager.NewPager(cli)
	if err := p.InitIterator(ctx, tableName); err != nil {
		panic(err)
	}

	// エミュレータを使用する場合、保存しているデータピッタリの値を設定しないとエラーになる
	pageSize := 1
	for {
		rows, err := p.GetDataPage(ctx, pageSize)
		if err != nil {
			panic(err)
		}
		for _, r := range rows {
			fmt.Println(r)
		}
		if p.GetPageToken() == "" {
			break
		}
	}
}
