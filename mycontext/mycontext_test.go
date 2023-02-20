package mycontext

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	// 	ctx := r.Context()
	//
	// 	data := make(chan string, 1)
	//
	// 	go func() {
	// 		data <- store.Fetch()
	// 	}()
	//
	// 	select {
	// 	case d := <-data:
	// 		fmt.Fprint(w, d)
	// 	case <-ctx.Done():
	// 		store.Cancel()
	// 	}
	// }
	}
}

type SpyStore struct {
	response  string
	t *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <- ctx.Done():
               s.t.Log("spy store was cancelled")
			   return 
			default:
		        time.Sleep(10 * time.Millisecond)
			    result += string(c)
			}
		}

		data <- result
	}()

	select {
	case <- ctx.Done():
	    return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestHandler(t *testing.T) {
	data := "hello"
	t.Run("request with no cancel", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got '%s', expected '%s'", response.Body.String(), data)
		}

		store.assertWasNotCancelled()
	})

	t.Run("request with no cancel", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)

		request = request.WithContext(cancellingCtx)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		store.aassertWasCancelled()
	})
}

func (s *SpyStore) aassertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("shouldn't have cancelled the request")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("shouldn't have cancelled the request")
	}
}
