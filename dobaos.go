package main

import "fmt"
import "encoding/json"
import "github.com/go-redis/redis"

var redisOpts = redis.Options{
    Addr: "localhost:6379",
    Password: "",
    DB: 0,
  }

var sub = redis.NewClient(&redisOpts)
var pubsub = sub.PSubscribe("dobaosgo_*");

var pub = redis.NewClient(&redisOpts)

func commonRequest(channel string, method string, payload interface{}) map[string] interface{} {
  var req map[string] interface{}
  req = make(map[string] interface{})
  req["method"] = method
  req["payload"] = payload
  req["response_channel"] = "dobaosgo_42"
  var res map[string] interface{}
  b, err := json.Marshal(req)
  
  if err != nil {
    return nil
  }

  pub.Publish(channel, b)

  {
    msgi, err := pubsub.ReceiveMessage()
    if err != nil {
      return nil
    }
    err = json.Unmarshal([]byte(msgi.Payload), &res)
  }


  return res
}

func getValue(id interface {}) map[string] interface{} {
  return commonRequest("dobaos_req", "get value", id)
}

func setValue(payload interface {}) map[string] interface{} {
  return commonRequest("dobaos_req", "set value", payload)
}

func main() {
  fmt.Println("hello, friend");
  defer pubsub.Close()

  //fmt.Println(getValue(42))

  var value map[string] interface{}
  value = make(map[string] interface{})
  value["id"] = 1
  value["value"] = true

  values := getValue([2]uint16 {2, 3})
  payload := values["payload"].([] interface{})
  var newValues []interface{}
  newValues = make([]interface{}, len(payload))

  for i := 0; i < len(payload); i += 1 {
    element := payload[i].(map[string] interface{})
    element["value"] = !element["value"].(bool)
    delete(element, "raw")
    newValues[i] = element
  }
  
  fmt.Println("new values: ", newValues)
  fmt.Println(setValue(newValues))
}
