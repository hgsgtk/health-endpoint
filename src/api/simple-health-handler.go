package api

import "net/http"

// SimpleHealthCheck return 200.
func SimpleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

type SimpleResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Detail  string `json:"string,omitempty"`
}

func SimpleDeepHealthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := NewMySQL(config.DB)
	if err != nil {
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
