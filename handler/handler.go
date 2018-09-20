package handler

import (
  "fmt"
  "log"
  "strings"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "net/http/cookiejar"
  "net/url"
)

func CorsProxy(w http.ResponseWriter, req *http.Request) {
  fmt.Println("Handling request from " + req.Header.Get("Referer"))
  type ReqBody struct {
    Url    string `json:url`
    Data   string `json:data`
    Method string `json:method`
  }

  var reqBody ReqBody

  if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  jar, err := cookiejar.New(nil)
  if err != nil {
    log.Fatal(err)
  }

  reqUrl, _ := url.Parse(reqBody.Url)
  jar.SetCookies(reqUrl, req.Cookies())

  remoteReq, _ := http.NewRequest(reqBody.Method, reqBody.Url, strings.NewReader(reqBody.Data))
  remoteReq.Header.Set("Content-Type", "application/json")
  client := &http.Client{ Jar: jar }
  remoteRes, err := client.Do(remoteReq)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  response, _ := ioutil.ReadAll(remoteRes.Body)

  for _, cookie := range jar.Cookies(reqUrl) {
    http.SetCookie(w, cookie)
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(remoteRes.StatusCode)
  w.Write([]byte(response))
}
