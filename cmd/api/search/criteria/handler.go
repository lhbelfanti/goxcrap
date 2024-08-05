package criteria

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	"goxcrap/internal/broker"
)

// EnqueueHandlerV1 HTTP Handler of the endpoint /criteria/enqueue/v1
func EnqueueHandlerV1(messageBroker broker.MessageBroker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var message IncomingBrokerMessageDTO
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			log.Error().Msg(err.Error())
			http.Error(w, InvalidRequestBody, http.StatusBadRequest)
			return
		}

		err = messageBroker.EnqueueMessage(string(message.Message))
		if err != nil {
			log.Error().Msg(err.Error())
			http.Error(w, FailedToEnqueueTask, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Criteria successfully enqueued"))
	}
}
