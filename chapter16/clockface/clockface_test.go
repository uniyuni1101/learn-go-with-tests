package clockface

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time   time.Time
		radian float64
	}{
		{simpleTime(0, 0, 0), math.Pi * 0},
		{simpleTime(0, 0, 15), math.Pi / 2},
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 45), math.Pi / 2 * 3},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("radian of second hand at %s", timeFormat(test.time)), func(t *testing.T) {
			got := SecondsInRadians(test.time)

			if !roughlyEqualFloat64(got, test.radian) {
				t.Errorf("got %f, want %f", got, test.radian)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 0, 15), Point{1, 0}},
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("second hand of point at %s", timeFormat(test.time)), func(t *testing.T) {
			got := SecondHandPoint(test.time)

			if !roughlyEqualPoint(got, test.point) {
				t.Errorf("got %v, want %v", got, test.point)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		time   time.Time
		radian float64
	}{
		{simpleTime(0, 0, 0), math.Pi * 0},
		{simpleTime(0, 15, 0), math.Pi / 2},
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 45, 0), math.Pi / 2 * 3},
		{simpleTime(0, 0, 15), math.Pi / 30 / 4},
		{simpleTime(0, 0, 30), math.Pi / 30 / 2},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("radian of minute hand at %s", timeFormat(test.time)), func(t *testing.T) {
			got := MinutesInRadians(test.time)

			if !roughlyEqualFloat64(got, test.radian) {
				t.Errorf("got %f, want %f", got, test.radian)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 15, 0), Point{1, 0}},
		{simpleTime(0, 30, 0), Point{0, -1}},
		{simpleTime(0, 45, 0), Point{-1, 0}},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("minute hand of point at %s", timeFormat(test.time)), func(t *testing.T) {
			got := MinuteHandPoint(test.time)

			if !roughlyEqualPoint(got, test.point) {
				t.Errorf("got %v, want %v", got, test.point)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		time   time.Time
		radian float64
	}{
		{simpleTime(0, 0, 0), 0},
		{simpleTime(3, 0, 0), math.Pi / 2},
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(9, 0, 0), math.Pi / 2 * 3},
		{simpleTime(12, 0, 0), 0},
		{simpleTime(15, 0, 0), math.Pi / 2},
		{simpleTime(18, 0, 0), math.Pi},
		{simpleTime(21, 0, 0), math.Pi / 2 * 3},
		{simpleTime(0, 15, 0), math.Pi / 6 / 4},
		{simpleTime(0, 30, 0), math.Pi / 6 / 2},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("radian of hour hand at %s", timeFormat(test.time)), func(t *testing.T) {
			got := HoursInRadians(test.time)

			if !roughlyEqualFloat64(got, test.radian) {
				t.Errorf("got %f, want %f", got, test.radian)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(3, 0, 0), Point{1, 0}},
		{simpleTime(6, 0, 0), Point{0, -1}},
		{simpleTime(9, 0, 0), Point{-1, 0}},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("minute hand of point at %s", timeFormat(test.time)), func(t *testing.T) {
			got := HourHandPoint(test.time)

			if !roughlyEqualPoint(got, test.point) {
				t.Errorf("got %v, want %v", got, test.point)
			}
		})
	}
}
func simpleTime(hour, min, sec int) time.Time {
	return time.Date(2000, 1, 1, hour, min, sec, 0, time.UTC)
}

func timeFormat(t time.Time) string {
	return fmt.Sprintf("%dh %dm %ds", t.Hour(), t.Minute(), t.Second())
}

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}
