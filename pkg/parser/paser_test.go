package parser

import (
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {

	test := `schema: "1"

id: bart
name: bart
owner: "clint@getweave.com"
repo: "weavelab/bart"

services:
- id: bart-service
  name: "Bart Service"
  type: node-port 
  ports:
  - id: "bart-service-port-grpc"
    name: "grpc"
    protocol: "TCP"
    nodePort: "36011"
    port: "80"
    targetPort: "9000"

env:
- id: "bart-env-listen-port"
  name: "LISTEN_PORT"
  type: "value"
  value: "9000"

pipeline:
  name: "Deploy to Dev"
  template: deploy-manual-check
  variables:
    clusterID: dev-ut
    namespace: phone
`
	fmt.Print(test)
	a, err := Marshal([]byte(test))
	if err != nil {
		fmt.Print(err)
		t.Fail()
	}
	if a.ID != "bart" {
		t.Fail()
	}
	fmt.Printf("%#v", a)

}




