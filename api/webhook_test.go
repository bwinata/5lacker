package api

import (
  "fmt"
  "time"
  "testing"
)

const (
  URL = "https://hooks.slack.com/services/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"
)

var wh SlackWebhook

func init () {
  wh, _ = NewWebhook ("WebhookTest", URL)
}

// -----------------------------------------------------------------------------

func TestSingleMsgWebhook (t *testing.T) {
  var err error = nil

  text := fmt.Sprintf ("Pass: TestSingleMsgWebhook. Time: %s", time.Now().Format(time.RFC3339))
  m := &Message { Text : &text }

  err = wh.Push (m)
  if err != nil {
    t.Error (err)
  }
}

// -----------------------------------------------------------------------------

func TestMultiMsgWebhook (t * testing.T) {
  var err error = nil

  for i := 0; i < 3; i++ {
    text := fmt.Sprintf ("Pass: TestMultiMsgWebhook. Time: %s", time.Now().Format(time.RFC3339))

    m := &Message { Text : &text }
    err = wh.Push (m)
    if err != nil {
      t.Error (err)
    }
  }
}

// -----------------------------------------------------------------------------

func TestSingleAttachment (t * testing.T) {
  var err error = nil

  text := fmt.Sprintf ("Pass: TestSingleAttachment. Time: %s", time.Now().Format(time.RFC3339))

  aTitle    := "Test title attachment"
  aPretext  := "Test pretext attachment"
  aText     := "Text text attachement"

  a := &Attachment {
    Title    : &aTitle,
    Pretext  : &aPretext,
    Text     : &aText,
  }

  m := &Message {
    Text    : &text,
    Attach  : []*Attachment{a},
  }

  err = wh.Push (m)
  if err != nil {
    t.Error (nil)
  }
}

// -----------------------------------------------------------------------------
