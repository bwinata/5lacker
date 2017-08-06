package api

import (
  "testing"
)

const (
  URL = "https://hooks.slack.com/services/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"
)

// -----------------------------------------------------------------------------

func TestSingleMsgWebhook (t *testing.T) {
  wh, err := NewWebhook ("WebhookTest", URL)
  if err != nil {
    t.Error (err)
  }

  text := "Your test if finished! Head to <https://strava.com>"
  m := &Message { Text : &text }

  err = wh.Push (m)
  if err != nil {
    t.Error (err)
  }
}

// -----------------------------------------------------------------------------

func TestMultiMsgWebhook (t * testing.T) {

}
