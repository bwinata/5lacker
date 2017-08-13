# 5lacker

**Status**: In development

**Overview**  
5lacker is a repository dedicated to the Slack API. Its purpose is to provide a clean and simple to use API for the Go language to all developers.

**Source Tree**  
The repo is broken up into two key folders - `api` and `plugins`.

`api` comprises of all the interfaces necessary to interat with Slack's Webhook and Web API for now. Real-Time Messaging and websockets will soon follow...

`plugins` enable third-party libraries to interface with Slack.

```
├── README.md
├── api
│   ├── README.md
│   ├── webapi.go
│   ├── webhook.go
│   └── webhook_test.go
└── plugins
    └── logrus
        ├── README.md
        ├── logrus.go
        └── logrus_test.go
```

**Interfaces**  
See the [api](https://github.com/bwinata/5lacker/tree/master/api) folder for a godoc of all the interfaces
