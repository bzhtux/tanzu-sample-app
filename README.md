# TSA aka Tanzu Sample Application

## Checks

[ ] PVC/PV (read and write)

## Notes

For NS add the following to the container.spec using the `downward API` :

```yaml
spec:
  containers:
   - env:
      - name: POD_NAMESPACE
        valueFrom:
          fieldRef:
            fieldPath: metadata.namespace
      - name: K8S_DEPLOYMENT
        valueFrom:
            fieldRef:
                fieldPath: metadata.name
```

And then get the value of the env `POD_NAMESPACE`
