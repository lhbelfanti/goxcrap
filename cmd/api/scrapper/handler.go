package scrapper

import (
	"log/slog"
	"net/http"
	"time"
)

const waitTimeAfterLogin time.Duration = 10

// ExecuteHandlerV1 HTTP Handler of the endpoint /execute-scrapper/v1
func ExecuteHandlerV1(execute Execute) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := execute(waitTimeAfterLogin)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, FailedToRunScrapper, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Scrapper successfully executed"))
	}
}
