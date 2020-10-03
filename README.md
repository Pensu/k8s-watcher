# k8s-watcher

This is a sample repo to create k8s-watcher. Currently, the code takes a kubeconfig file and uses Kuberenetes APIs to pull labels on all kubernetes cluster nodes.
Then it checks against a database on required labels and prints out if the node has that label or not.
