package payloads

import "encoding/json"

type CreatePayload struct {
	Index string
	DocumentId    string
	Body  json.RawMessage
}