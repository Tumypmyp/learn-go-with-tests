package clockface

import (
  "time"
  "testing"
  "math"
)

func TestSecondsToRadians(t *testing.T) {
  cases := []struct {
    time time.Time
    angle float64
  }{
    {sampleTime(0, 0, 30), math.Pi},
    {sampleTime(0, 0, 0), 0},
    {sampleTime(0, 0, 45), math.Pi / 2 * 3},
    {sampleTime(0, 0, 7), math.Pi / 30 * 7},
  }

  for _, c := range cases {
    t.Run(testName(c.time), func(t *testing.T) {
      got := secondsToRadians(c.time)

      if c.angle != got {
        t.Fatalf("wanted %v radians but go %v", c.angle, got)
      }
    })
  }
}

func sampleTime(h, m, s int) time.Time {
  return time.Date(1337, time.October, 28, h, m, s, 0, time.UTC)
}

func testName(t time.Time) string {
  return t.Format("15:04:05")
}

func TestSecondHandVector(t *testing.T) {
  cases := []struct {
    time  time.Time
    point Point
  }{
    {sampleTime(0, 0, 30), Point{0, -1}},
    {sampleTime(0, 0, 45), Point{-1, 0}},
  }

  for _, c := range cases {
    t.Run(testName(c.time), func(t *testing.T) {
      got := secondHandVector(c.time)

      if !roughlyEqualPoint(c.point, got) {
        t.Fatalf("wanted %v point but go %v", c.point, got)
      }
    })
  }
}

func roughlyEqualPoint(a, b Point) bool {
  return roughlyEqualFloat64(a.X, b.X) &&
    roughlyEqualFloat64(a.Y, b.Y)
}

func roughlyEqualFloat64(a, b float64) bool {
  const equalityThreshold = 1e-9
  return math.Abs(a - b) < equalityThreshold
}
