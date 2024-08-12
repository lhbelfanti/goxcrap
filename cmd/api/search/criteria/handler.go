package criteria

import (
	"encoding/json"
	"net/http"

	"goxcrap/internal/broker"
	"goxcrap/internal/log"
)

// EnqueueHandlerV1 HTTP Handler of the endpoint /criteria/enqueue/v1
func EnqueueHandlerV1(messageBroker broker.MessageBroker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var message IncomingBrokerMessageDTO
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, InvalidRequestBody, http.StatusBadRequest)
			return
		}
		ctx = log.With(ctx, log.Param("message", message))

		err = messageBroker.EnqueueMessage(ctx, string(message.Message))
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, FailedToEnqueueTask, http.StatusInternalServerError)
			return
		}

		log.Info(ctx, "Criteria successfully enqueued")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Criteria successfully enqueued"))
	}
}
