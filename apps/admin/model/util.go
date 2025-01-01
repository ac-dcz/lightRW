package model

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type EsModel map[string]any

type M map[string]interface{}

func parseRespToEsModel(resp *esapi.Response) ([]EsModel, error) {
	ret := make(map[string]any)
	data := make([]EsModel, 0)
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}
	if hits, ok := ret["hits"].(map[string]any); ok {
		if items, ok := hits["hits"].([]any); ok {
			for _, item := range items {
				if entry, ok := item.(map[string]any); ok {
					if source, ok := entry["_source"].(map[string]any); ok {
						data = append(data, source)
					}
				}
			}
		}
	}
	return data, nil
}
