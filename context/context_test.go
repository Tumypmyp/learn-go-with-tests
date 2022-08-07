package context

import (
  "testing"
  "net/http"
  "net/http/httptest"
  "time"
  "context"
)


func TestServer(t *testing.T) {
  t.Run("returns data from store", func(t *testing.T) {
    data := "hello, world"
    store := &SpyStore{response: data, t: t}
    svr := Server(store)

    request := httptest.NewRequest(http.MethodGet, "/", nil)
    response := httptest.NewRecorder()

    svr.ServeHTTP(response, request)

    if response.Body.String() != data {
      t.Errorf("got %q want %q", response.Body.String(), data)
    }

  })
  t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
    data := "hello, world"
    store := &SpyStore{response: data, t: t}
    svr := Server(store)

    request := httptest.NewRequest(http.MethodGet, "/", nil)

    cancellingCtx, cancel := context.WithCancel(request.Context())
    time.AfterFunc(5 * time.Millisecond, cancel)
    request = request.WithContext(cancellingCtx)

    response := &SpyResponseWriter{}

    svr.ServeHTTP(response, request)
    
    if response.written {
      t.Error("a response should not have been written")
    }
  })
  
}
