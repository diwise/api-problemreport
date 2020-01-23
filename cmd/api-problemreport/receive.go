package main

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"

	"github.com/iot-for-tillgenglighet/api-problemreport/pkg/database"
	"github.com/iot-for-tillgenglighet/api-problemreport/pkg/models"
	"github.com/iot-for-tillgenglighet/messaging-golang/pkg/messaging/telemetry"
)

func receiveProblemreport(msg amqp.Delivery) {

	log.Info("Message received from queue: " + string(msg.Body))

	depth := &telemetry.Problemreport{}
	err := json.Unmarshal(msg.Body, depth)

	if err != nil {
		log.Error("Failed to unmarshal message")
		return
	}

	newdepth := &models.Problemreport{
		Device:    depth.Origin.Device,
		Latitude:  depth.Origin.Latitude,
		Longitude: depth.Origin.Longitude,
		Depth:     depth.Depth,
		Timestamp: depth.Timestamp,
	}

	database.GetDB().Create(newdepth)
}
