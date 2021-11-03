# k8sensus

Trying to build a leader election process for Kubernetes using the Kubernetes `client-go` client.

## Why?

To make systems more fault-tolerant; handling failures in replicas is crucial for higher availability.

## How?

- [x] Start by creating a lock object.
- [x] Make leader update/renew lease. (inform other replicas about its leadership)
- [x] Make candidate pods check the lease object.
- [x] If leader fails, re-elect new leader.

## Lease Durations and Renew Deadlines

- Lease duration = 15 seconds
- Renew deadline = 10 seconds
- Retry period = 2 seconds

```go
LeaseDuration: 15 * time.Second,
RenewDeadline: 10 * time.Second,
RetryPeriod:   2 * time.Second,
```

## Usage Example

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

After applying definitions, check the pods:

```sh
❯ kubectl get pods
NAME                      READY   STATUS    RESTARTS   AGE
k8sensus-66459bcf-g2dw7   1/1     Running   0          10m
k8sensus-66459bcf-mfljf   1/1     Running   0          10m
k8sensus-66459bcf-v9df8   1/1     Running   0          10m
```

## Production Usage

*If you like challenges and love debugging on Friday Nights, then, please feel free to use it on your production cluster. 😋*

> **Non-satirical note:** Do not use in production.
