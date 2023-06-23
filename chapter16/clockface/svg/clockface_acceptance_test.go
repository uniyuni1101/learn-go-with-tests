package svg_test

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"
	"time"

	"github.com/uniyuni1101/learn-go-with-tests/chapter16/clockface/svg"
)

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 60}},
		{simpleTime(0, 0, 30), Line{150, 150, 150, 240}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("write second hand in svg format at %s", timeFormat(c.time)), func(t *testing.T) {
			b := bytes.Buffer{}

			svg.SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the second hand line of %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}

}
func TestSVGWriterMinutedHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 70}},
		{simpleTime(0, 30, 0), Line{150, 150, 150, 230}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("write minute hand in svg format at %s", timeFormat(c.time)), func(t *testing.T) {
			b := bytes.Buffer{}

			svg.SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the minute hand line of %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}

}
func TestSVGWriterHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 100}},
		{simpleTime(6, 0, 0), Line{150, 150, 150, 200}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("write minute hand in svg format at %s", timeFormat(c.time)), func(t *testing.T) {
			b := bytes.Buffer{}

			svg.SVGWriter(&b, c.time)

			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the hour hand line of %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}

}

func containsLine(line Line, lines []Line) bool {
	for _, l := range lines {
		if line == l {
			return true
		}
	}

	return false
}

func simpleTime(hour, min, sec int) time.Time {
	return time.Date(2000, 1, 1, hour, min, sec, 0, time.UTC)
}

func timeFormat(t time.Time) string {
	return fmt.Sprintf("%dh %dm %ds", t.Hour(), t.Minute(), t.Second())
}
