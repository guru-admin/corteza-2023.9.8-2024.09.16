package geolocation

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/sql"
)

type (
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	}

	Properties struct {
		Name string `json:"name"`
	}

	Full struct {
		Geometry   Geometry   `json:"geometry"`
		Properties Properties `json:"properties"`
	}
)

func Parse(ss []string) (m Full, err error) {
	if len(ss) == 0 {
		return
	}

	err = json.Unmarshal([]byte(ss[0]), &m)
	return
}

func (set *Full) Scan(src any) error          { return sql.ParseJSON(src, set) }
func (set Full) Value() (driver.Value, error) { return json.Marshal(set) }
