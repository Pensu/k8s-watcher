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

// Note: the example only works with the code within the same release/branch.
package main

import (
	"context"
	//	"flag"
	"fmt"
	//	"os"
	//	"path/filepath"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {

	http.HandleFunc("/", kdata)
	fmt.Println("Starting the server at port 8080")
	http.ListenAndServe(":8080", nil)

}

func kdata(w http.ResponseWriter, r *http.Request) {

	config, err := clientcmd.BuildConfigFromFlags("", "./config")

	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})

	for i := 0; i < len(nodes.Items); i++ {
		node_name := nodes.Items[i].ObjectMeta.Name
		node_label := nodes.Items[i].ObjectMeta.Labels
		for label, _ := range node_label {
			if label == "arch" {
				fmt.Fprintf(w, "Arch label found in node %s\n", node_name)
				break
			}
		}
		fmt.Fprintf(w, "No arch label found in node %s\n", node_name)
	}
}
