package federation

import (
	"io"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/options"
)

type (
	EncoderAdapter interface {
		BuildStructure(io.Writer, options.FederationOpt, interface{}) (interface{}, error)
		BuildData(io.Writer, options.FederationOpt, interface{}) (interface{}, error)
	}
)
