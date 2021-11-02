# k8sensus

Trying to build a leader election process for Kubernetes using the Kubernetes `client-go` client.

## Why?

To make systems more fault-tolerant; handling failures in replicas is crucial for higher availability.

## How?

- [x] Start by creating a lock object.
- [ ] Make leader update/renew lease. (inform other replicas about its leadership)
- [ ] Make candidate pods check the lease object.
- [ ] If leader fails, re-elect new leader.

## Production Usage

*If you like challenges and love debugging on Friday Nights, then, please feel free to use it on your production cluster. 😋*

> **Non-satirical note:** Do not use in production.
