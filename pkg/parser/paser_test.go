package parser

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {

	test := `schema: "1"
id: bart
name: bart
owner: "clint@getweave.com"
repo: "weavelab/bart"

services:
- name: bart-service
  type: node-port 
  ports:
  - name: "grpc"
    protocol: "TCP"
    nodePort: "36011"
    port: "80"
    targetPort: "9000"

env:
- name: "LISTEN_PORT"
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
	a, err := Parse([]byte(test))
	if err != nil {
		fmt.Print(err)
		t.Fail()
	}
	if a.ID != "bart" {
		t.Fail()
	}
	fmt.Printf("%#v", a)

}




