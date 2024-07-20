package payloads

import "encoding/json"

type UpdatePayload struct {
	Index       string
	DocumentId  string
	Body        json.RawMessage
}