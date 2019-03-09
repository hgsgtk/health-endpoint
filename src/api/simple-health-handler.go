package api

import "net/http"

type SimpleResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Detail  string `json:"string,omitempty"`
}

func SimpleHealthCheck(w http.ResponseWriter, r *http.Request) {
	rs := SimpleResponse{
		Status: http.StatusOK,
	}
	respondJson(w, rs, http.StatusOK)
}

func SimpleDeepHealthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := NewMySQL(config.DB)
	if err != nil {
		// DBへのコネクションにてエラーが発生した場合は503レスポンス
		rs := SimpleResponse{
			Status:  http.StatusServiceUnavailable,
			Message: "failed to get connection database",
			Detail:  err.Error(),
		}
		respondJson(w, rs, rs.Status)
		return
	}

	rs := SimpleResponse{
		Status:  http.StatusOK,
		Message: "success to connect server",
	}
	respondJson(w, rs, rs.Status)
}
