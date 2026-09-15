# Kubernetes Basics

## Why Kubernetes Exists

Imagine you have a backend application.

For example:

```text
Java / Go Backend
```

You package that application inside a Docker container.

Now imagine your application is important and must stay available.

You may want:

```text
3 copies of the backend running
```

Why multiple copies?

Because if one crashes, users should still be able to use the application.

Without Kubernetes, somebody or some script would have to:

- Start containers
- Restart failed containers
- Check if enough copies are running
- Move workloads between machines
- Expose the application to users
- Handle application updates

Kubernetes automates much of this work.

---

# Core Idea

You tell Kubernetes:

```text
"I want 3 copies of my backend running."
```

Kubernetes tries to maintain that desired state.

Example:

```text
Backend Pod 1 ✅
Backend Pod 2 ✅
Backend Pod 3 ✅
```

If one crashes:

```text
Backend Pod 1 ✅
Backend Pod 2 ❌
Backend Pod 3 ✅
```

Kubernetes can create another one:

```text
Backend Pod 1 ✅
Backend Pod 3 ✅
Backend Pod 4 ✅
```

Back to:

```text
3 running copies
```

This is one of the main reasons Kubernetes is useful.

---

# First Understand the Layers

The important pieces are:

```text
Cluster
  ↓
Nodes
  ↓
Pods
  ↓
Containers
```

And separately:

```text
Deployment = manages Pods

Service = gives network access to Pods
```

We will understand each one.

---

# Container

A container is the packaged application.

Example:

```text
Java Backend
+
Java Runtime
+
Required Libraries
+
Configuration
```

Docker packages these together into a container image.

When it runs:

```text
Docker Container
    ↓
Your Backend Application
```

Example:

```text
Container
└── Spring Boot application
```

or:

```text
Container
└── Go backend application
```

---

# Pod

A **Pod** is the basic unit Kubernetes runs.

A Pod usually contains one main application container.

Example:

```text
Pod
└── Java Backend Container
```

So the mental model is:

```text
Kubernetes
    ↓
runs Pod
    ↓
Pod contains
    ↓
Application Container
```

For beginner understanding:

```text
Pod = wrapper around your running application container
```

Important:

A Pod is not the same as the physical computer.

The Pod runs on a computer called a **Node**.

---

# Node

A **Node** is one machine that participates in the Kubernetes cluster.

It can be:

- A physical server
- A virtual machine
- A cloud VM

For example:

```text
Node 1
```

might simply be a Linux computer.

That machine can run multiple Pods:

```text
Node 1
├── Pod A
├── Pod B
├── Pod C
└── Pod D
```

So:

```text
Node = computer/machine

Pod = application workload running on that computer
```

---

# Node vs Client

A Node is NOT a client.

A client is something sending requests to your application.

Examples:

```text
Browser
Mobile App
Another Backend
Postman
```

Example:

```text
Browser
  ↓
Backend
```

The browser is the client.

The Node is the computer where the backend is physically running.

---

# Node vs Pod

Think of it like:

```text
Computer
   ↓
runs applications
```

In Kubernetes:

```text
Node
   ↓
runs Pods
```

Example:

```text
Node 1
├── Pod 1
│   └── Backend Container
│
├── Pod 2
│   └── Backend Container
│
└── Pod 3
    └── Backend Container
```

One Node can run many Pods depending on:

- CPU
- Memory
- Resource limits
- Kubernetes configuration

---

# Cluster

A **Cluster** is a group of Nodes managed together by Kubernetes.

Example:

```text
Kubernetes Cluster
├── Node 1
├── Node 2
├── Node 3
├── Node 4
└── Node 5
```

If you have five Kubernetes machines:

```text
5 machines
=
5 Nodes
```

All of them together:

```text
Kubernetes Cluster
```

---

# Cluster Mental Model

```text
Kubernetes Cluster
│
├── Node 1
│   ├── Pod
│   └── Pod
│
├── Node 2
│   ├── Pod
│   ├── Pod
│   └── Pod
│
└── Node 3
    └── Pod
```

So:

```text
Cluster = group of Nodes

Node = one machine

Pod = workload running on a Node

Container = actual packaged application
```

---

# Deployment

Now suppose you want:

```text
3 copies of your backend
```

You do not usually manually create three separate Pods and manage them yourself.

You create a **Deployment**.

The Deployment says:

```text
"I want 3 replicas of this application."
```

Replica means:

```text
one copy of the application
```

So:

```text
Deployment
    ↓
desired replicas = 3
    ↓
Pod 1
Pod 2
Pod 3
```

---

# What Happens If a Pod Crashes?

Suppose:

```text
Deployment says:
desired replicas = 3
```

Currently:

```text
Pod 1 ✅
Pod 2 ✅
Pod 3 ✅
```

Then:

```text
Pod 2 crashes ❌
```

Now only two are running:

```text
Pod 1 ✅
Pod 3 ✅
```

Kubernetes sees:

```text
Desired = 3

Actual = 2
```

So it creates another Pod:

```text
Pod 4 ✅
```

Now:

```text
Pod 1 ✅
Pod 3 ✅
Pod 4 ✅
```

Actual becomes 3 again.

That is the idea of maintaining the **desired state**.

---

# Deployment Mental Model

```text
Deployment
    ↓
"I want 3 Pods"
    ↓
Pod 1
Pod 2
Pod 3
```

If one disappears:

```text
Deployment
    ↓
notices only 2 exist
    ↓
creates replacement Pod
```

So remember:

```text
Deployment = manages desired number of Pods
```

---

# Service

Now we have another problem.

Suppose clients directly call a Pod:

```text
Client
  ↓
Pod 1
```

The Pod has an IP address.

But Pods are temporary.

Pod 1 may crash.

Kubernetes creates Pod 4.

Pod 4 may have a different IP address.

The client should not need to know:

```text
Pod 1 IP
Pod 2 IP
Pod 3 IP
```

So Kubernetes provides a **Service**.

---

# What a Service Does

A Service gives a stable network entry point to a set of Pods.

Instead of:

```text
Client
  ↓
Specific Pod IP
```

we use:

```text
Client
  ↓
Service
  ↓
Pods
```

Example:

```text
             Service
            /   |   \
           ↓    ↓    ↓
        Pod 1 Pod 2 Pod 3
```

The Service can route traffic to healthy Pods.

---

# If a Pod Is Replaced

Suppose:

```text
Client
  ↓
Service
  ↓
Pod 1
Pod 2
Pod 3
```

Pod 2 crashes:

```text
Pod 2 ❌
```

Kubernetes creates:

```text
Pod 4 ✅
```

Now:

```text
Client
  ↓
Service
  ↓
Pod 1
Pod 3
Pod 4
```

The client does not need to know that Pod 2 disappeared.

That is why the Service is useful.

---

# Deployment vs Service

This distinction is extremely important.

## Deployment

Deployment answers:

```text
"How many Pods should exist?"
```

Example:

```text
Deployment
Desired replicas = 3
```

Its job is:

```text
manage Pods
```

---

## Service

Service answers:

```text
"How should clients reach those Pods?"
```

Its job is:

```text
provide stable network access
```

---

# Simple Comparison

```text
Deployment = manages Pods

Service = routes traffic to Pods
```

or:

```text
Deployment
    ↓
creates/manages
    ↓
Pods
```

while:

```text
Client
    ↓
Service
    ↓
Pods
```

---

# Complete Mental Model

```text
                    Client
                      |
                      v
                   Service
                      |
            ---------------------
            |         |         |
            v         v         v
          Pod 1     Pod 2     Pod 3
            |         |         |
            --------------------------------
                  running on Nodes
            --------------------------------
            |                   |
            v                   v
          Node 1              Node 2

All Nodes together
        ↓
Kubernetes Cluster
```

And separately:

```text
Deployment
    |
    | manages desired replicas
    v
   Pods
```

---

# Full Layer Model

```text
Kubernetes Cluster
│
├── Node 1
│   ├── Pod 1
│   │   └── Container
│   │       └── Backend Application
│   │
│   └── Pod 2
│       └── Container
│           └── Backend Application
│
└── Node 2
    └── Pod 3
        └── Container
            └── Backend Application
```

Clients do not normally care which Node or Pod handles the request.

They interact through:

```text
Service
```

---

# Example: Java Backend

Imagine a Spring Boot backend.

First:

```text
Spring Boot Application
```

Package it:

```text
Docker Image
```

Run that image:

```text
Container
```

Kubernetes runs that container inside:

```text
Pod
```

Pods run on:

```text
Nodes
```

Nodes together form:

```text
Cluster
```

Deployment manages how many Pods exist:

```text
Deployment
    ↓
3 Pods
```

Service lets clients access them:

```text
Client
  ↓
Service
  ↓
Pod 1 / Pod 2 / Pod 3
```

---

# Example Failure

Suppose your backend receives thousands of requests.

You run:

```text
3 Pods
```

Current state:

```text
Service
   ↓
Pod 1 ✅
Pod 2 ✅
Pod 3 ✅
```

Pod 2 crashes.

Deployment sees:

```text
Desired = 3

Running = 2
```

It creates:

```text
Pod 4
```

Now:

```text
Service
   ↓
Pod 1 ✅
Pod 3 ✅
Pod 4 ✅
```

Clients continue using the Service.

They do not need to know that the Pod changed.

---

# Desired State

A key Kubernetes idea is **desired state**.

You tell Kubernetes what you want.

Example:

```text
I want 3 backend Pods.
```

Kubernetes continuously tries to make reality match that request.

```text
Desired = 3
Actual = 3
→ everything is fine
```

If:

```text
Desired = 3
Actual = 2
```

Kubernetes tries to create another one.

This is a fundamental Kubernetes idea.

---

# Self-Healing

When Kubernetes replaces failed application instances automatically, this is often described as **self-healing**.

Example:

```text
Pod crashes
    ↓
Kubernetes detects missing replica
    ↓
Replacement Pod created
```

---

# Scaling

Suppose three Pods are not enough.

You can change:

```text
replicas = 3
```

to:

```text
replicas = 5
```

Then:

```text
Deployment
    ↓
Pod 1
Pod 2
Pod 3
Pod 4
Pod 5
```

This is scaling the application horizontally.

Horizontal scaling means:

```text
more application instances
```

instead of making one machine larger.

---

# Kubernetes and Docker Relationship

Docker and Kubernetes solve different problems.

## Docker

Docker helps package and run an application inside a container.

```text
Application
    ↓
Docker Container
```

## Kubernetes

Kubernetes manages many containerized applications.

```text
Kubernetes
    ↓
Pods
    ↓
Containers
```

Simple mental model:

```text
Docker = package/run the application

Kubernetes = manage many running application instances
```

---

# Important Terms

## Container

```text
Packaged running application
```

## Pod

```text
Kubernetes unit containing application container(s)
```

## Node

```text
Computer that runs Pods
```

## Cluster

```text
Group of Nodes managed together
```

## Deployment

```text
Manages desired number of Pods
```

## Service

```text
Stable network access to Pods
```

## Replica

```text
One copy of the application
```

---

# Common Interview Questions

## What is Kubernetes?

Kubernetes is a container orchestration platform used to deploy, scale, and manage containerized applications.

In simple terms:

```text
It helps keep the required application containers running.
```

---

## What is a Pod?

A Pod is the smallest deployable unit in Kubernetes.

It normally contains one main application container.

---

## What is a Node?

A Node is a machine in the Kubernetes cluster where Pods run.

---

## What is a Cluster?

A Cluster is a collection of Nodes managed together by Kubernetes.

---

## What is a Deployment?

A Deployment manages the desired state and number of application Pods.

Example:

```text
replicas = 3
```

means Kubernetes should try to keep three Pods running.

---

## What is a Service?

A Service provides stable network access to a set of Pods.

Clients communicate with the Service instead of depending on individual Pod IP addresses.

---

## What happens when a Pod crashes?

If the Pod is controlled through a Deployment, Kubernetes notices that the desired replica count is not satisfied and creates a replacement Pod.

---

## Why should a client not directly call a Pod IP?

Pods are temporary and can be replaced.

A replacement Pod may receive a different IP address.

A Service provides a stable endpoint.

---

## Deployment vs Service?

```text
Deployment = manages Pods

Service = gives network access to Pods
```

---

# Interview Answer

Kubernetes is a container orchestration platform.

Applications run inside containers, which Kubernetes normally runs inside Pods.

Pods run on Nodes, and multiple Nodes form a Kubernetes Cluster.

A Deployment manages the desired number of Pods and can replace failed Pods.

A Service provides stable network access to those Pods so clients do not depend on individual Pod IP addresses.

---

# Quick Revision

```text
Container
= packaged application

Pod
= Kubernetes unit containing the application container

Node
= machine running Pods

Cluster
= group of Nodes

Deployment
= manages desired number of Pods

Service
= stable network entry point for Pods

Replica
= one copy of an application
```

---

# One-Line Mental Model

```text
Client
  ↓
Service
  ↓
Pods
  ↓
Nodes
  ↓
Cluster
```

And:

```text
Deployment
    ↓
manages Pods
```

---

# Final Picture

```text
                       Kubernetes Cluster
                               |
                --------------------------------
                |                              |
              Node 1                         Node 2
                |                              |
         ----------------                ----------------
         |              |                |              |
       Pod 1          Pod 2            Pod 3          Pod 4
         |              |                |              |
     Container      Container        Container      Container
         |              |                |              |
      Backend        Backend          Backend        Backend


Client
   |
   v
Service
   |
   +----------> Pod 1
   +----------> Pod 2
   +----------> Pod 3
   +----------> Pod 4


Deployment
   |
   └----------> ensures required number of Pods exist
```