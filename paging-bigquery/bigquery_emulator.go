package main

import (
	"context"
	"fmt"
	"path/filepath"

	"cloud.google.com/go/bigquery"
	"github.com/Reimei1213/lab/paging-bigquery/pager"
	"github.com/goccy/bigquery-emulator/server"
	"google.golang.org/api/option"
)

func main() {
	const projectID = "test"

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
	if err := p.InitIterator(ctx, "test.dataset1.table_a"); err != nil {
		panic(err)
	}

	// エミュレータを使用する場合、保存しているデータピッタリの値を設定しないとエラーになる
	pageSize := 3
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
