package api

import (
  "github.com/gorilla/http"
)

type SlackWeb interface {

}

type Webapi struct {
  name        string
  accessToken string
  client      http.Client
}

// -----------------------------------------------------------------------------
// -- Public Functions
// -----------------------------------------------------------------------------

func NewWebapi (accessToken, name string) (SlackWeb, error) {
  w := new (Webapi)

  w.accessToken = accessToken
  w.name        = name
  w.client      = http.DefaultClient

  return w, nil
}
