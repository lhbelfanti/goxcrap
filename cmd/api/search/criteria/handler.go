package criteria

import (
	"encoding/json"
	"net/http"

	"goxcrap/internal/broker"
	"goxcrap/internal/http/response"
	"goxcrap/internal/log"
)

// EnqueueHandlerV1 HTTP Handler of the endpoint /criteria/enqueue/v1
func EnqueueHandlerV1(messageBroker broker.MessageBroker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var message IncomingBrokerMessageDTO
		err := json.NewDecoder(r.Body).Decode(&message)
		if err != nil {
			response.Send(ctx, w, http.StatusBadRequest, InvalidRequestBody, nil, err)
			return
		}
		ctx = log.With(ctx, log.Param("message", message))

		err = messageBroker.EnqueueMessage(ctx, string(message.Message))
		if err != nil {
			response.Send(ctx, w, http.StatusInternalServerError, FailedToEnqueueTask, nil, err)
			return
		}

		response.Send(ctx, w, http.StatusOK, "Criteria successfully enqueued", nil, nil)
	}
}
