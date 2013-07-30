package gauth

import (
  "net/http"
  "bytes"
  "strings"
)

func CanonicalRep(req *http.Request) (rep string, err error) {
  err = req.ParseForm()
  if err != nil {
    return rep, err
  }
  var repBuff bytes.Buffer

  // original URI
  repBuff.WriteString(req.RequestURI)
  // HTTP verb (GET, POST,...) uppercased
  repBuff.WriteString(strings.ToUpper(req.Method))

  rep = repBuff.String()
  return rep, nil
}
