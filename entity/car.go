package entity

import (
	"aeb-cloud/sensor"
)

type Car struct {
	Id        string
	Speed     float32
	EnvSensor *sensor.EnvSensor
	PosSensor *sensor.PosSensor
}
