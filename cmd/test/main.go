/*
Copyright 2016 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"time"

	"k8s.io/client-go/pkg/api/unversioned"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/apis/extensions/v1beta1"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeconfig = flag.String("kubeconfig", "/Users/clint/.kube/config", "absolute path to the kubeconfig file")
)

func main() {
	flag.Parse()
	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	var replicas int32
	replicas = 2
	selector := &unversioned.LabelSelector{}
	selector.MatchLabels = make(map[string]string)
	selector.MatchLabels["run"] = "hello-world"
	templateLabels := make(map[string]string)
	templateLabels["run"] = "hello-world"
	containers := make([]v1.Container, 1)
	containers[0] = v1.Container{
		Name:  "hello-world",
		Image: "yeasy/simple-web:latest",
	}

	deplTemplateSpec := v1beta1.DeploymentSpec{
		Replicas: &replicas,
		Template: v1.PodTemplateSpec{
			ObjectMeta: v1.ObjectMeta{
				Labels: templateLabels,
			},
			Spec: v1.PodSpec{
				Containers: containers,
			},
		},
		Selector: selector,
	}

	deplSpec := &v1beta1.Deployment{
		ObjectMeta: v1.ObjectMeta{
			Name:      "ploiotest",
			Namespace: "phone",
			Labels:    templateLabels,
		},
		Spec: deplTemplateSpec,
	}

	_, err = clientset.ExtensionsV1beta1().Deployments("phone").Create(deplSpec)
	if err != nil {
		panic(err.Error())
	}
	nodes, err := clientset.Core().Nodes().List(v1.ListOptions{})
	for _, node := range nodes.Items {
		fmt.Println(node.Name)
	}
	for {
		pods, err := clientset.Core().Pods("").List(v1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		time.Sleep(10 * time.Second)
	}
}
