package main

import (
  "bytes"
  "encoding/json"
  "net/http"

  "github.com/drone/drone-plugin-go"
)

func main() {
  var repo = plugin.Repo{}
  var build = plugin.Build{}
  var vargs = struct {
    Urls []string `json:"urls"`
  }{}

  plugin.Param("repo": &repo)
  plugin.Param("build", &build)
  plugin.Param("vargs", &vargs)
  plugin.Parse()

  // post build and repo data to webhook urls
  data := struct{
    Repo plugin.Repo `json:"repo"`
    Build plugin.Build `json:"build"`
  }{repo, build}

  payload, _ := json.Marshal(&data)

  for _, url := range vargs.Urls {
    resp, _ := http.Post(url, "application/json", bytes.NewBuffer(payload))
    resp.Body.Close()
  }
}

