# Kubernetes Basics

## Why Kubernetes?

Kubernetes helps keep applications running in the desired state.

Example:
- I want 3 copies of my backend running.
- If one copy crashes, Kubernetes creates a replacement.

## Main Concepts

### Cluster

A cluster is a group of Kubernetes machines working together.

### Node

A node is one machine/computer in the Kubernetes cluster.

A node can run many Pods.

### Pod

A Pod is the basic unit Kubernetes runs.

A Pod usually contains the application container.

### Deployment

A Deployment manages Pods.

Example:
- Desired replicas = 3
- Kubernetes tries to keep 3 Pods running.

### Service

A Service gives clients a stable way to reach Pods.

The client does not need to know individual Pod IP addresses.

## Mental Model

```text
Kubernetes Cluster
├── Node 1
│   ├── Pod
│   └── Pod
├── Node 2
│   └── Pod
└── Node 3
    ├── Pod
    └── Pod