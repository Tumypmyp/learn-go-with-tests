package main

import (
  "testing"
  "bytes"
  "reflect"
  "time"
)

func TestCountdown(t *testing.T) {
  t.Run("prints 3 to Go!", func(t *testing.T) {
    buffer := &bytes.Buffer{}

    Countdown(buffer, &spyCountdownOperations{})

    got := buffer.String()
    want := "3\n2\n1\nGo!"

    if got != want {
      t.Errorf("got %q want %q", got, want)
    }
  })

  t.Run("sleep before every print", func(t *testing.T) {
    spySleepPrinter := &spyCountdownOperations{}

    Countdown(spySleepPrinter, spySleepPrinter)

    want := []string{write, sleep, write, sleep, write, sleep, write}

    if !reflect.DeepEqual(want, spySleepPrinter.calls) {
      t.Errorf("wanted calls %v got %v", want, spySleepPrinter.calls)
    }
  })
}

type spyCountdownOperations struct {
  calls []string
}

func (s *spyCountdownOperations) Sleep() {
  s.calls = append(s.calls, sleep)
}

func (s *spyCountdownOperations) Write(p []byte) (n int, err error) {
  s.calls = append(s.calls, write)
  return
}

const write = "write"
const sleep = "sleep"

type SpyTime struct {
  durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
  s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
  sleepTime := 5 * time.Second

  spyTime := &SpyTime{}
  sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
  sleeper.Sleep()

  if spyTime.durationSlept != sleepTime {
    t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
  }
}

