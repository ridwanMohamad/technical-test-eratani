package response

import "golangtest/src/properties"

type GlobalResponse struct {
	Status   string                 `json:"status"`
	Message  string                 `json:"message"`
	Data     interface{}            `json:"data,omitempty"`
	Metadata *properties.Pagination `json:"metadata,omitempty"`
}
