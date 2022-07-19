package racer

import (
  "testing"
  "time"
  "net/http"
  "net/http/httptest"
)

func TestRacer(t *testing.T) {
  slowServer := makeDelayedServer(20 * time.Millisecond)
  fastServer := makeDelayedServer(0 * time.Millisecond)

  defer slowServer.Close()
  defer fastServer.Close()

  slowURL := slowServer.URL
  fastURL := fastServer.URL

  want := fastURL
  got := Racer(slowURL, fastURL)

  if got != want {
      t.Errorf("want %q, got %q", want, got)
  }

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
  return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      time.Sleep(delay)
      w.WriteHeader(http.StatusOK)
  }))
}
