package clockface

import (
  "time"
  "math"
)

type Point struct {
  x float64
  y float64
}

func SecondHand(t time.Time) Point {
  return Point{150, 60}
}

func secondsToRadians(t time.Time) float64 {
  return math.Pi / (30 / float64(t.Second()))
}

func secondHandVector(t time.Time) Point {
  angle := secondsToRadians(t)
  x := math.Sin(angle)
  y := math.Cos(angle)
  return Point{x, y}
}
