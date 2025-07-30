package response

type ApiResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}

func Success(data interface{}) ApiResponse {
	return ApiResponse{
		Data:    data,
		Message: "Success",
		Status:  "OK",
	}
}

func Error(message string) ApiResponse {
	return ApiResponse{
		Message: message,
		Status:  "ERROR",
	}
}

func PaginatedResponse(data interface{}, total int64, page int, size int, sort string) ApiResponse {
	return ApiResponse{
		Data: map[string]interface{}{
			"items": data,
			"total": total,
			"page":  page,
			"size":  size,
			"sort":  sort,
		},
		Message: "Success",
		Status:  "OK",
	}
}
