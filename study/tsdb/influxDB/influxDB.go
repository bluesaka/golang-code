/**
github.com/influxdata/influxdb-client-go/v2
github.com/influxdata/influxdb/client/v2
*/
package influxDB

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb/client/v2"
	"log"
	"time"
)

const (
	ServerUrl = "http://localhost:8086"
)

func Write() {
	cli := influxdb2.NewClient(ServerUrl, "my-token")
	defer cli.Close()

	// 同步阻塞写入
	writeAPI := cli.WriteAPIBlocking("my-org", "test_db")

	// Create point using full params constructor
	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45.0},
		time.Now())
	if err := writeAPI.WritePoint(context.Background(), p); err != nil {
		log.Fatalf("write point err: %v\n", err)
	}

	// Create point using fluent style
	p = influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("avg", 23.2).
		AddField("max", 45.0).
		SetTime(time.Now())
	if err := writeAPI.WritePoint(context.Background(), p); err != nil {
		log.Fatalf("write point err: %v\n", err)
	}

	// write directly line protocol
	line := fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0)
	if err := writeAPI.WriteRecord(context.Background(), line); err != nil {
		log.Fatalf("write line err: %v\n", err)
	}
}

func Query() {
	cli := influxdb2.NewClient(ServerUrl, "my-token")
	defer cli.Close()

	queryAPI := cli.QueryAPI("")
	// Flux语法，必须包含bucket、时间范围、过滤条件
	// 注意influxdb.conf flux-enabled=true开启
	// 否则会报错：403 Forbidden: Flux query service disabled. Verify flux-enabled=true in the [http] section of the InfluxDB config.
	query := `from(bucket:"test_db")
	|> range(start: 2021-06-28T02:52:48.066682Z) 
	|> filter(fn: (r) =>
		r._measurement == "stat"
	)`
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatalf("query error: %v\n", err)
	}

	for result.Next() {
		// Observe when there is new grouping key producing new table
		if result.TableChanged() {
			fmt.Printf("table: %s\n", result.TableMetadata().String())
		}
		// read result
		fmt.Printf("row: %s", result.Record().String())
	}
	if result.Err() != nil {
		fmt.Printf("query result error: %v\n", result.Err())
	}
}

func Write2() {
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: ServerUrl,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	tags := map[string]string{"unit": "temperature"}
	fields := map[string]interface{}{
		"avg": 33.2,
		"max": 46.1,
	}
	point, err := client.NewPoint("stat", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	batchPoints, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "test_db",
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}
	batchPoints.AddPoint(point)
	if err := cli.Write(batchPoints); err != nil {
		log.Fatalf("write error: %v\n", err)
	}
}

func Query2() {
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: ServerUrl,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	q := client.NewQuery("select * from stat", "test_db", "")
	result, err := cli.Query(q)
	if err != nil {
		log.Fatalf("query error: %v\n", err)
	}
	if result.Error() != nil {
		log.Fatalf("query result error: %v\n", result.Error())
	}
	fmt.Printf("result: %+v\n", result.Results)

	for i, row := range result.Results[0].Series[0].Values {
		t, err := time.Parse(time.RFC3339, row[0].(string))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("[%2d] time:%s, value:%s\n", i, t.Format("2006-01-02 15:04:05"), row[2])
	}

}