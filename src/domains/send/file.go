package send

import "mime/multipart"

type FileRequest struct {
	BaseRequest
	File    *multipart.FileHeader `json:"file" form:"file"`
	Caption string                `json:"caption" form:"caption"`
	TraceCode string `json:"trace_code" form:"trace_code"`
}
