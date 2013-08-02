gauth
========

HMAC HTTP authentication utils

Usage:
```go
import (
  ....
  "github.com/OnWander/gauth"
)

func handleSomeHttpReq(res http.ResponseWriter, req *http.Request) {
  err := gauth.Authenticate(req)
  if err != nil {
    http.Error(res, err.Error(), http.StatusUnauthorized)
    return
  }

  // do other stuff...
}
```
