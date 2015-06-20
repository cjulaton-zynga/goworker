package goworker

type Payload struct {
	Metadata map[string]interface{} `json:"metadata"`
	Class    string                 `json:"class"`
	Args     []interface{}          `json:"args"`
}
