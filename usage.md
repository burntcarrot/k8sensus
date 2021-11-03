# Usage

After applying the definitions, check the pods:

```sh
❯ kubectl get pods
NAME                        READY   STATUS    RESTARTS   AGE
k8sensus-67798d9cf6-gqfjf   1/1     Running   0          2m2s
k8sensus-67798d9cf6-qwxj6   1/1     Running   0          2m2s
k8sensus-67798d9cf6-vtlx6   1/1     Running   0          15s
```

**To simulate leader election, delete a pod:**

```sh
❯ kubectl delete pod k8sensus-67798d9cf6-n4sgl
```

**Check the logs of all pods:**

```sh
❯ kubectl logs k8sensus-67798d9cf6-vtlx6
🚢🏗️ k8sensus is running!
Creating lease using the following metadata:
Lease Name: k8sensus-lease
Pod Name: k8sensus-67798d9cf6-vtlx6
Namespace: default
I1103 10:34:56.358804       1 leaderelection.go:248] attempting to acquire leader lease default/k8sensus-lease...
I1103 10:34:56.374900       1 main.go:60] New leader is: k8sensus-67798d9cf6-n4sgl
I1103 10:35:07.734547       1 main.go:60] New leader is: k8sensus-67798d9cf6-qwxj6
```

```sh
❯ kubectl logs k8sensus-67798d9cf6-qwxj6
🚢🏗️ k8sensus is running!
Creating lease using the following metadata:
Lease Name: k8sensus-lease
Pod Name: k8sensus-67798d9cf6-qwxj6
Namespace: default
I1103 10:33:37.924516       1 leaderelection.go:248] attempting to acquire leader lease default/k8sensus-lease...
I1103 10:33:37.938187       1 main.go:60] New leader is: k8sensus-67798d9cf6-n4sgl
I1103 10:35:07.562421       1 leaderelection.go:258] successfully acquired lease default/k8sensus-lease
I1103 10:35:07.562627       1 main.go:57] I'm the new leader! 😋
I1103 10:35:07.562845       1 main.go:70] k8sensus is running sample task.
I1103 10:35:17.563958       1 main.go:70] k8sensus is running sample task.
I1103 10:35:27.566011       1 main.go:70] k8sensus is running sample task.
```

**Check the lease object:**

```md
Lease Durations and Renew Deadlines

- Lease duration = 15 seconds
- Renew deadline = 10 seconds
- Retry period = 2 seconds
```

```sh
❯ kubectl describe lease k8sensus-lease
Name:         k8sensus-lease
Namespace:    default
Labels:       <none>
Annotations:  <none>
API Version:  coordination.k8s.io/v1
Kind:         Lease
Metadata:
  Creation Timestamp:  2021-11-03T10:33:30Z
  Managed Fields:
    API Version:  coordination.k8s.io/v1
    Fields Type:  FieldsV1
    fieldsV1:
      f:spec:
        f:acquireTime:
        f:holderIdentity:
        f:leaseDurationSeconds:
        f:leaseTransitions:
        f:renewTime:
    Manager:         k8sensus
    Operation:       Update
    Time:            2021-11-03T10:33:30Z
  Resource Version:  15947
  UID:               7b1aac4c-baa8-4c3d-a6c8-18deff2f0c2f
Spec:
  Acquire Time:            2021-11-03T10:35:07.497077Z
  Holder Identity:         k8sensus-67798d9cf6-qwxj6
  Lease Duration Seconds:  15
  Lease Transitions:       1
  Renew Time:              2021-11-03T10:39:26.731022Z
Events:                    <none>
```

## Cleanup

```
❯ kubectl delete deployment k8sensus
deployment.apps "k8sensus" deleted
```
