package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"github.com/Reimei1213/lab/paging-bigquery/pager"
)

func main() {
	const (
		projectID = ""
		tableID   = ""
	)

	ctx := context.Background()
	cli, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	p := pager.NewPager(cli)
	if err := p.InitIterator(ctx, tableID); err != nil {
		panic(err)
	}

	pageSize := 2
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
