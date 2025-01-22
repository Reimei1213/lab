package pager

import (
	"context"
	"fmt"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

type Pager struct {
	cli       *bigquery.Client
	iterator  *bigquery.RowIterator
	pageToken string
}

func NewPager(cli *bigquery.Client) *Pager {
	return &Pager{
		cli: cli,
	}
}

func (p *Pager) InitIterator(ctx context.Context, tableName string) error {
	q := p.cli.Query(fmt.Sprintf("SELECT * FROM `%s` LIMIT 10", tableName))
	it, err := q.Read(ctx)
	if err != nil {
		return err
	}
	p.iterator = it
	p.pageToken = ""
	return nil
}

/*
	$ export GODEBUG=http2debug=2
	を設定するとネットワーク上のやり取りを観察でき、
	NextPage の実行ごとにリクエストを送ってそう
*/
func (p *Pager) GetDataPage(_ context.Context, pageSize int) ([][]bigquery.Value, error) {
	var rows [][]bigquery.Value
	pager := iterator.NewPager(p.iterator, pageSize, p.pageToken)
	nextToken, err := pager.NextPage(&rows)
	if err != nil {
		return nil, err
	}
	p.pageToken = nextToken
	return rows, nil
}

func (p *Pager) GetPageToken() string {
	return p.pageToken
}
