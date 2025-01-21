package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"cloud.google.com/go/bigquery"
	"github.com/goccy/bigquery-emulator/server"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const projectID = "test"

func main() {
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

	fmt.Println("=======================")

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

	pageSize := 1
	pageToken := ""
	_, _, err = getBigQueryPage(ctx, cli, pageToken, pageSize)
	if err != nil {
		log.Fatal(err)
	}
	//for {
	//	rows, nextPageToken, err := getBigQueryPage(ctx, cli, pageToken, pageSize)
	//	if err != nil {
	//		log.Fatalf("Error fetching BigQuery data: %v", err)
	//	}
	//	for _, r := range rows {
	//		fmt.Println(r)
	//	}
	//
	//	if nextPageToken == "" {
	//		break
	//	}
	//
	//	pageToken = nextPageToken
	//}

	log.Println("End")

}

// https://github.com/googleapis/google-cloud-go/blob/main/bigquery/integration_test.go#L856-L892
func getBigQueryPage(ctx context.Context, cli *bigquery.Client, pageToken string, pageSize int) ([][]bigquery.Value, string, error) {
	query := cli.Query(`SELECT * FROM test.dataset1.table_a`)
	iter, err := query.Read(ctx)
	if err != nil {
		return nil, "", err
	}

	//iter.PageInfo().Token = pageToken
	//var results [][]bigquery.Value
	//for {
	//	var r []bigquery.Value
	//	if err := iter.Next(&r); err != nil {
	//		if errors.Is(err, iterator.Done) {
	//			break
	//		}
	//		return nil, "", err
	//	}
	//	results = append(results, r)
	//}
	//return results, iter.PageInfo().Token, nil

	pager := iterator.NewPager(iter, pageSize, pageToken)
	var results [][]bigquery.Value
	nextPageToken, err := pager.NextPage(&results)
	if err != nil {
		return nil, "", err
	}

	for _, r := range results {
		fmt.Println(r)
	}
	fmt.Println(nextPageToken)
	return nil, "", nil
	//return results, nextPageToken, nil
}
