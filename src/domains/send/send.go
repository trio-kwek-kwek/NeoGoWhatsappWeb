package send

type GenericResponse struct {
	TraceCode string `json:"trace_code"`
	MessageID string `json:"message_id"`
	Sender    string `json:"sender"`
	Status    string `json:"status"`
}
