package bench

import(
  "net/http"
  "fmt"
  "bytes"
  "encoding/json"
)

type chaincode struct{
  Jsonrpc string `json:"jsonrpc"`
  Method  string `json:"method"`
  Params  Param  `json:"params"`
  ID      int    `json:"id"`
}

type Param struct{
  Type        int      `json:"type"`
  ChaincodeID ChaincodeID `json:"chaincodeID"`
  CtorMsg     CtorMsg     `json:"ctorMsg"`
}

type ChaincodeID struct{
  Name string `json:"name"`
}

type CtorMsg struct{
  Args []string `json:"args"`
}

var ReqCount int

var TransactionType = map[string]int{
      "deploy": 2,
      "invoke": 4,
      "query": 5,
    }

func postJSON(url string, data []byte) {
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
  if err != nil {
    fmt.Println(err)
    panic(err)
  }
  req.Header.Set("Content-Type", "application/json")

  client := http.Client{}
  resp, err2 := client.Do(req)
  if err2 != nil {
    fmt.Println(err2)
    panic(err2)
  }

  defer resp.Body.Close()
}

func createChainReq(action string, msg CtorMsg) []byte{
  chaincodeID := ChaincodeID{
   // blockchain一つしか使わない予定なので
    Name: "mycc",
  }
  params := Param{
    Type: TransactionType[action],
    ChaincodeID: chaincodeID,
    CtorMsg: msg,
  }

  post_data := chaincode{
    Jsonrpc: "2.0",
    Method:  action,
    Params:  params,
    ID:      ReqCount,
  }
  data, err := json.Marshal(post_data)
  if err != nil {
    fmt.Print("JSON marshaling failed: %s", err)
  }
  return data
}
