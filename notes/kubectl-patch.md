#patch kubectl deploy

1. Create a patch file
patch-priority.yaml:
spec:
  template:
    spec:
      priorityClassName: high-priority


2. Apply the patch file with kubectl
kubectl patch deployment <deployment-name> \
  -n <namespace> \
  --patch-file patch-priority.yaml
  
Replace <deployment-name> with your Deployment name.
Replace <namespace> with the correct namespace (or omit if in default).

3. Verify
kubectl get deployment myapp-deployment -n development -o yaml | grep priorityClassName
priorityClassName: high-priority
