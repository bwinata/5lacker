/* -----------------------------------------------------------------------------
 * Name: Slack Wbehook Logrus Plugin
 *
 * Description:
 *      This plugin allows the user to integrate notifications into Slack via the
 *      logger utility - logrus. <https://github.com/sirupsen/logrus>. The current
 *      implementation allows notification via Slack's provisioned webhook interface.
 *      For more information on webhooks, see - <https://api.slack.com/incoming-webhooks>
 *
 * Author : Barry Winata <barrywinata1@gmail.com>
 * Date   : August 2017
 * -----------------------------------------------------------------------------
 */

package logrus

import (
  "fmt"
  "time"

  "github.com/sirupsen/logrus"
  "github.com/bwinata/5lacker/api"
)

// -----------------------------------------------------------------------------
// -- Public Types
// -----------------------------------------------------------------------------

type Options struct {
  channel     string
  icon        string
  minLogLevel logrus.Level
  async       bool
}

type Loghook struct {
  name        string
  ops         *Options
  webhook     api.SlackWebhook
}

// -----------------------------------------------------------------------------

// NewLogHook returns a hook which is passed into logrus' API function -
// AddHook. Based on the specified min log level, any calls by logrus automatically
// invokes this hook and delivers a notification to Slack.
func NewLogHook (name, url string, ops * Options) (logrus.Hook, error) {
  loghook := new (Loghook)

  loghook.name  = name

  if ops == nil {
    loghook.ops = NewOptions ()
  } else {
    loghook.ops = ops
  }

  webhook, err := api.NewWebhook (name, url)
  if err != nil { return nil, err }

  loghook.webhook = webhook

  return loghook, nil
}

// -----------------------------------------------------------------------------

// NewOptions generates a pointer to Options which sets itself up with default
// values. These can be overriiden by invoking its setter functions.
// NOTE: Depending on the setup of your webhook, if it has been strictly created
// with a channel / user, then it cannot overriden.
func NewOptions () *Options {
  return &Options {
    channel     : "",
    icon        : "",
    minLogLevel : logrus.WarnLevel,
  }
}

// -----------------------------------------------------------------------------

// Set the Slack channel to send the logs. Default is an empty string and is
// dictattd by the wehbook. For certain webhooks, this can be overriden.
func (this * Options) SetChannel (channel string) *Options {
  this.channel = channel
  return this
}

// -----------------------------------------------------------------------------

// Set icon emoji. Default is empty and is dictated by the webhook. For certain
// webhooks, this can be overriden.
func (this * Options) SetIcon (icon string) *Options {
  this.icon = icon
  return this
}

// -----------------------------------------------------------------------------

// SetMinLogLevel sets the minimum level at which logs are sent to Slack.
// By default they are et to the Warning level. This means that based on logrus'
// levels, anything below this level i.e. Error, Fatal and Panic will trigger
// a log to be delivered to Slack.
func (this * Options) SetMinLogLevel (level logrus.Level) *Options {
  this.minLogLevel = level
  return this
}

// -----------------------------------------------------------------------------

// SetAsync sets a boolean to trigger the log events either synchronously or
// asynchronoulsy. By default this is set to true.
func (this * Options) SetAsync () *Options {
  this.async = true
  return this
}

// -----------------------------------------------------------------------------

func (this * Loghook) Levels () []logrus.Level {
  levels := make ([]logrus.Level, 0, len (logrus.AllLevels))
  for _, l := range (logrus.AllLevels) {
    if this.ops.minLogLevel >= l {
      levels = append (levels, l)
    }
  }
  return levels
}

// -----------------------------------------------------------------------------

func (this * Loghook) Fire (e * logrus.Entry) error {
  var err error = nil

  rawText, strErr := e.String ()
  if strErr != nil { return strErr }

  message := fmt.Sprintf ("*APP*: %s | *LEVEL*: %s | *TIME*: %s\n %s",
                           this.name,
                           e.Level.String (),
                           e.Time.Format (time.RFC3339),
                           rawText)

  m := &api.Message { Text : &message }

  if this.ops.async {
    go this.webhook.Push (m)
  } else {
    err = this.webhook.Push (m)
  }

  return err
}
