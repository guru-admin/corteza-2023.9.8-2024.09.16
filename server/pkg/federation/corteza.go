package federation

import (
	"io"

	"github.com/guru-admin/corteza-2023.9.8-2024.09.16/server/pkg/options"
)

type (
	EncoderAdapterCortezaInternal struct{}

	ResponseWrapper struct {
		Response interface{} `json:"response"`
	}
)

// Build a default Corteza response
func (a EncoderAdapterCortezaInternal) BuildStructure(w io.Writer, o options.FederationOpt, p interface{}) (interface{}, error) {
	return ResponseWrapper{
		Response: listModuleResponseCortezaInternal{
			Filter: p.(ListStructurePayload).Filter,
			Set:    p.(ListStructurePayload).Set,
		}}, nil
}

// Build a default Corteza response
func (a EncoderAdapterCortezaInternal) BuildData(w io.Writer, o options.FederationOpt, p interface{}) (interface{}, error) {
	return ResponseWrapper{
		Response: listRecordResponseCortezaInternal{
			Filter: p.(ListDataPayload).Filter,
			Set:    p.(ListDataPayload).Set,
		}}, nil
}
