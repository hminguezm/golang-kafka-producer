package service

import (
	"github.com/google/uuid"
	"os"
	"time"
	"wrk-connector/src/shared/domain/constant"
	"wrk-connector/src/shared/domain/entity/event"
)

func MessageGenerator(
	eventName string,
	eventDataFormat string,
	eventType string,
	version string,
	eventOrigin string,
) event.MessageBase {
	message := event.MessageBase{
		EventId:         uuid.New().String(),
		EventName:       eventName,
		EventDataFormat: eventDataFormat,
		Type:            eventType,
		Timestamp:       time.Now().Format(constant.TimeFormat),
		Version:         version,
		Country:         os.Getenv("ENVIRONMENT_COUNTRY"),
		Origin:          eventOrigin,
	}

	return message
}
