package send

type LinkRequest struct {
	BaseRequest
	Caption string `json:"caption"`
	Link    string `json:"link"`
	TraceCode string `json:"trace_code" form:"trace_code"`
}
