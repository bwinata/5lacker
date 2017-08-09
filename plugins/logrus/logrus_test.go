package logrus

import (
  "os"
  "testing"

  "github.com/sirupsen/logrus"
)

const (
  URL = "https://hooks.slack.com/services/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"
)

func init () {
  logrus.SetFormatter (&logrus.TextFormatter {
    ForceColors       : false,
    DisableTimestamp  : false,
    DisableColors     : false,
    FullTimestamp     : true,
  })

  logrus.SetOutput (os.Stdout)
  logrus.SetLevel (logrus.InfoLevel)

  // Add log hook into Slack
  hook, err := NewLogHook ("LoghookTest", URL, nil)
  if err != nil { panic (err) }
  logrus.AddHook (hook)
}

// -----------------------------------------------------------------------------

func TestLogWarn (t * testing.T) {
  logrus.Warn ("Warn: This is a test log")
}

// -----------------------------------------------------------------------------

func TestLogError (t * testing.T) {
  logrus.Error ("Error: This is a test log")
}
