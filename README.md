# 🏗️ k8sensus

Leader election process for Kubernetes using the Kubernetes `client-go` client.

![Preview](/static/k8sensus-preview.gif)

## Table of Contents
  - [Why?](#why)
  - [How does it work?](#how-does-it-work)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Local Development](#local-development)
  - [Production Usage](#production-usage)
  - [What did I learn?](#what-did-i-learn)

## Why?

To make systems more fault-tolerant; handling failures in replicas is crucial for higher availability. A leader election process ensures that if the leader fails, the candidate replicas can be elected as the leader.

## How does it work?

An overview on how it works:

- [x] Start by creating a lock object.
- [x] Make leader update/renew lease. (inform other replicas about its leadership)
- [x] Make candidate pods check the lease object.
- [x] If leader fails, re-elect new leader.

## Installation

**By cloning the repo:**

```sh
git clone https://github.com/burntcarrot/k8sensus
cd k8sensus
kubectl apply -f k8s/rbac.yaml
kubectl apply -f k8s/deployment.yaml
```

**By copying the deployment and RBAC definitions:**

```sh
kubectl apply -f k8s/rbac.yaml
kubectl apply -f k8s/deployment.yaml
```

## Usage

A complete example on how to use k8sensus is described [here](usage.md).

## Local Development

There are two commands exposed by the `Makefile`:

For applying definitions:

```sh
make apply
```

For cleaning up k8sensus deployment:

```sh
make clean
```

## Production Usage

*If you like challenges and love debugging on Friday nights, then, please feel free to use it on your production cluster. 😋*

> **Non-satirical note:** Do not use in production.

## What did I learn?

After hours of debugging and opening up 20 tabs of documentation, here's what I learnt:

- Kubernetes has a [leaderelection package](https://pkg.go.dev/k8s.io/client-go/tools/leaderelection) in its client.
- After reading the first line in the documentation, I was a bit disappointed:
>  This implementation does not guarantee that only one client is acting as a leader (a.k.a. fencing).
- This made me write this code, I wanted a single-leader workflow.
- For interacting, we can use `CoordinationV1` to get the client. ([docs](https://pkg.go.dev/k8s.io/client-go/kubernetes@v0.22.3#Clientset.CoordinationV1))
- `leaderelection` (under `client-go`) provides a `LeaseLock` type ([docs](https://pkg.go.dev/k8s.io/client-go@v0.22.3/tools/leaderelection/resourcelock#LeaseLock)), which can be used for the leader election. (leaders renew time in the lease)
- `leaderelection` also provides `LeaderCallbacks` ([docs](https://pkg.go.dev/k8s.io/client-go@v0.22.3/tools/leaderelection/resourcelock#LeaderCallbacks)) which can be used for handling leader events like logging when a new pod/replica gets elected as the new leader, etc.
