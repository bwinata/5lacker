/* -----------------------------------------------------------------------------
 * Name: Webhook Interface API
 *
 * Description:
 *      This interface allows an authorised user to deliver messages to Slack's servers as
 *      a webhook. In layman's terms - it basically allows you to send notifications
 *      to any channel (public / private) and user you desire, provided you have
 *      the right permissions.
 *
 * Author : Barry Winata <barrywinata1@gmail.com>
 * Date   : August 2017
 * -----------------------------------------------------------------------------
 */

package api

import (
  "fmt"
  "time"
  "bytes"
  "net/http"
  "encoding/json"
)

// -----------------------------------------------------------------------------
// -- Public Type
// -----------------------------------------------------------------------------

type Attachment struct {
  Title       * string
  Pretext     * string
  Text        * string
  Markdown    []*string
}

type Message struct {
  Text        * string
  Username    * string
  Icon        * string
  Channel     * string
  Attach      []*Attachment
}

type Webhook struct {
  name        string
  url         string
  debounceInt time.Duration
  client      http.Client
}

// -----------------------------------------------------------------------------
// -- Interfaces
// -----------------------------------------------------------------------------

type SlackWebhook interface {
  Push (msg * Message) error
  SetDebounce (time.Duration) error
}

// -----------------------------------------------------------------------------
// -- Private Functions
// -----------------------------------------------------------------------------

func MarshalMessage (msg * Message) (map[string]interface{}) {
  intfMap := map[string]interface{} {}

  if msg.Text == nil || len (*msg.Text) == 0 {
    panic ("Invalid field. 'Text' cannot be empty")
  }

  intfMap["text"] = msg.Text

  if msg.Username != nil { intfMap["username"]    = msg.Username}
  if msg.Icon     != nil { intfMap["icon_emoji"]  = msg.Icon    }
  if msg.Channel  != nil { intfMap["channel"]     = msg.Channel }
  if msg.Attach   != nil { intfMap["attachments"] = msg.Attach  }

  return intfMap
}

// -----------------------------------------------------------------------------
// -- Public Functions
// -----------------------------------------------------------------------------

func NewWebhook (name, url string) (SlackWebhook, error) {
  wh := new (Webhook)

  wh.name = name
  wh.url  = url
  wh.debounceInt = time.Millisecond * 1000 // ms

  return wh, nil
}

// -----------------------------------------------------------------------------

func (this * Webhook) Push (msg * Message) (error) {

  // Assemble message into interface map
  intfMap := MarshalMessage (msg)

  // Encode interface map into JSON
  json, err := json.Marshal (intfMap)
  if err != nil { return err }

  // Make a POST request to Slack
  result, err := http.Post (this.url, "application/json", bytes.NewBuffer (json))

  if err != nil { return err }

  if result.StatusCode != http.StatusOK {
    return fmt.Errorf ("Cannot push request. Status=%d, Err=%v", result.StatusCode, result.Status)
  }

  // fmt.Printf ("Result: %v", result)

  return nil
}

// -----------------------------------------------------------------------------

func (this * Webhook) SetDebounce (ms time.Duration) error {
  if ms < time.Millisecond * 1000 {
    return fmt.Errorf ("Invalid interval: %d. Must be >= 1000", ms)
  }
  this.debounceInt = ms
  return nil
}
