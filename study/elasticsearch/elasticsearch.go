/**
elasticsearch
github.com/elastic/go-elasticsearch/v7
*/
package elasticsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

var (
	Addresses = []string{"http://localhost:9200"}
	IndexName = "go-es-index-test"
	client    *elasticsearch.Client
)

// Info get elasticsearch info
func Info() {
	resp, err := client.Info()
	failOnError(err, "Error getting response")

	defer resp.Body.Close()
	log.Println("res:", resp)
}

// Index index doc to elasticsearch
func Index() {
	var buf bytes.Buffer
	data := map[string]string{
		"title":   "标题一定要响亮",
		"content": "内容一定要精彩",
	}
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		failOnError(err, "Error encoding data")
	}

	resp, err := client.Index(IndexName, &buf)
	failOnError(err, "Error Index response")

	defer resp.Body.Close()
	fmt.Println("body:", resp.String())
}

// Search search document
func Search() {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "响亮",
			},
		},
		"highlight": map[string]interface{}{
			"pre_tags":  []string{"<font color='yellow'>"},
			"post_tags": []string{"</font>"},
			"fields": map[string]interface{}{
				"title": map[string]interface{}{},
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		failOnError(err, "Error encoding query")
	}

	resp, err := client.Search(
		client.Search.WithIndex(IndexName),
		client.Search.WithBody(&buf),
		client.Search.WithPretty(),
	)
	failOnError(err, "Error getting response")

	defer resp.Body.Close()
	fmt.Println(resp.String())
}

// Create create doc to elasticsearch
func Create() {
	var buf bytes.Buffer
	data := map[string]string{
		"title":   "标题一定要响亮",
		"content": "内容一定要精彩",
	}
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		failOnError(err, "Error encoding data")
	}

	// 指定doc_id，如果doc_id已存在返回409 Conflict错误
	resp, err := client.Create(IndexName, "doc_id", &buf)
	failOnError(err, "Error create response")

	defer resp.Body.Close()
	fmt.Println("body:", resp.String())
	if resp.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
			failOnError(err, "Error parsing the response body")
		} else {
			log.Fatalf("[%s] %s: %s",
				resp.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
}

// init
func init() {
	client = newClient()
}

// newClient create elasticsearch client, exit on error
func newClient() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: Addresses,
	}
	client, err := elasticsearch.NewClient(cfg)
	failOnError(err, "Error creating the client")
	return client
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
