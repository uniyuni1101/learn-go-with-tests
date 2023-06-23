package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

func SecondsInRadians(t time.Time) float64 {
	return math.Pi / (secondsInHalfClock / float64(t.Second()))
}

func SecondHandPoint(t time.Time) Point {
	radian := SecondsInRadians(t)
	x := math.Sin(radian)
	y := math.Cos(radian)
	return Point{x, y}
}

func MinutesInRadians(t time.Time) float64 {
	secAddRad := SecondsInRadians(t) / secondsInClock
	return math.Pi/(minutesInHalfClock/float64(t.Minute())) + secAddRad
}

func MinuteHandPoint(t time.Time) Point {
	radian := MinutesInRadians(t)
	x := math.Sin(radian)
	y := math.Cos(radian)
	return Point{x, y}
}

func HoursInRadians(t time.Time) float64 {
	minAddRad := MinutesInRadians(t) / hoursInClock
	return math.Pi/(hoursInHalfClock/float64(t.Hour()%hoursInClock)) + minAddRad
}

func HourHandPoint(t time.Time) Point {
	radian := HoursInRadians(t)
	x := math.Sin(radian)
	y := math.Cos(radian)
	return Point{x, y}
}
