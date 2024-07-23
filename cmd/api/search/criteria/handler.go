package criteria

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"goxcrap/internal/broker"
)

// EnqueueHandlerV1 HTTP Handler of the endpoint /criteria/enqueue/v1
func EnqueueHandlerV1(messageBroker broker.MessageBroker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var message IncomingBrokerMessageDTO
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, InvalidRequestBody, http.StatusBadRequest)
			return
		}

		err = messageBroker.EnqueueMessage(string(message.Message))
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, FailedToEnqueueTask, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Criteria successfully enqueued"))
	}
}
