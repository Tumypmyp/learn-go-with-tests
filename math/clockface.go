package clockface

import (
  "time"
  "math"
)

type Point struct {
  X float64
  Y float64
}

const secondHandLength = 90
const clockCenterX = 150
const clockCenterY = 150
func SecondHand(t time.Time) Point {
  p := secondHandVector(t)
  p = Point{p.X * secondHandLength, -p.Y * secondHandLength}
  p.X += clockCenterX
  p.Y += clockCenterY
  return p
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
