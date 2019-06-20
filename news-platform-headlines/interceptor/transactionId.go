package transactionId

import (
	"context"
	"github.com/gocql/gocql"
	"net/http"
)

// To prevent golint complaint "should not use basic type string as key in context.WithValue"
type key string

// Key is a identifier for Context
const Key key = "TransactionID"

// FromContext extracts X-Request-ID value from HTTP Request' Context
func FromContext(ctx context.Context) (s string) {
	s, _ = ctx.Value(Key).(string)
	return
}

// AddIDRequestMiddleware sets TransactionID to header and context
func AddIDRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var reqID = r.Header.Get(string(Key))
		if reqID == "" {
			reqID = gocql.TimeUUID().String()
			r.Header.Add("TransactionID", reqID)
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), Key, reqID)))
	})
}
