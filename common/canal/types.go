package canal

import "encoding/json"

type EntryType uint32

const (
	UpdateType EntryType = iota
	DeleteType
	InsertType
)

type Column struct {
	Name      string `json:"name"`
	Value     any    `json:"value"`
	MySqlType string `json:"mysql_type"`
	IsUpdate  bool   `json:"is_update"`
	IsNull    bool   `json:"is_null"`
	IsKey     bool   `json:"is_key"`
}

type Record struct {
	DataBase      string    `json:"data_base"`
	Table         string    `json:"table"`
	Type          EntryType `json:"type"`
	BeforeColumns []*Column `json:"before_columns"`
	AfterColumns  []*Column `json:"after_columns"`
}

func (r *Record) Encode() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Record) Decode(data []byte) error {
	return json.Unmarshal(data, r)
}
