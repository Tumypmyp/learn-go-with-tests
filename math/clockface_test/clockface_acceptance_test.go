package clockface_test

import (
  "testing"
  "time"
  //"../clockface"
  "github.com/tumypmyp/learn-go-with-tests/math"
  //"github.com/quii/learn-go-with-tests/math/v1/clockface"
)

func TestSecondHandAtMidnight(t *testing.T) {
  tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

  want := clockface.Point{x: 150, y: 150-90}
  got := clockface.SecondHand(tm)

  if got != want {
    t.Errorf("Got %v, wanted %v", got, want)
  }
}

/*
func TestSecondHandAt30Seconds(t *testing.T) {
  tm := time.Date(1000, time.January, 1, 0, 0, 30, 0, time.UTC)

  want := Point{X: 150, Y: 150 + 90}
  got := SecondHand(tm)

  if got != want {
    t.Errorf("Got %v, wanted %v", got, want)
  }
}
//*/
