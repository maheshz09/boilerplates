# cni 

The task only provides calico-operator url and not the custom-resources yaml file URL. is not the question and selection of the cni is the ask for this question 

By the way, yes, in the exam you’ll typically be given just the calico-operator.yaml URL and no custom resource file. The process remains mostly the same — you just need to replace calico-operator.yaml with the custom-resources.yaml in same url and download files  to your local exam environment and apply them accordingly.😊


Once you choose the right CNI (like Calico, Cilium etc),  it is already clear that suport networkPolicy enforcement therefore I'm assuming the NetworkPolicy setup (step) done was  only for validation purpose.

---- choose right CNI
    to choose a right cni we can see the exesting cni installed in the cluser by seeing k get pods -A like that 
    Install calico
    By the way, yes, in the exam you’ll typically be given just the calico-operator.yaml URL and no custom resource file. 
    The process remains mostly the same — you just need to replace calico-operator.yaml with the custom-resources.yaml in same url and download files  to your local exam environment and apply them accordingly.😊 
    IMP
    In the CKA exam, if asked to install Calico, you will usually only be given the tigera-operator.yaml URL and not the custom-resources.yaml. The tigera-operator.yaml only installs the operator, but networking will not work until you also apply the custom-resources.yaml. Both are always required. The trick is: take the same base URL where you got tigera-operator.yaml and replace tigera-operator.yaml with custom-resources.yaml. Example: kubectl apply -f https://docs.projectcalico.org/manifests/tigera-operator.yaml && kubectl apply -f https://docs.projectcalico.org/manifests/custom-resources.yaml. In the exam, the cluster usually has no CNI preinstalled, so you must apply both files yourself, otherwise Pods will remain in ContainerCreating state. Final rule: apply operator + custom-resources = working Calico CNI.
    kubectl apply -f https://raw.githubusercontent.com/projectcalico/calico/v3.30.3/manifests/tigera-operator.yaml
    kubectl apply -f https://raw.githubusercontent.com/projectcalico/calico/v3.30.3/manifests/custom-resources.yaml
	
	
# CKA 2025 – Install and Configure CNI (Flannel / Calico)

---

## ❓ Question
Install and configure a Container Network Interface (CNI) that:
- Allows pods to communicate with each other
- Supports **NetworkPolicy**
- Must be installed from **manifests** (not Helm)

Options given:
- **Flannel** → ❌ Does not support NetworkPolicy
- **Calico** → ✅ Supports NetworkPolicy

👉 Correct choice: **Calico**

---

## ✅ Solution

### 1. Install Calico Operator
kubectl apply -f https://raw.githubusercontent.com/projectcalico/calico/v3.29.2/manifests/tigera-operator.yaml


----
cat <<EOF | kubectl apply -f -
apiVersion: operator.tigera.io/v1
kind: Installation
metadata:
  name: default
spec:
  cni:
    type: Calico
  calicoNetwork:
    bgp: Disabled
    ipPools:
    - cidr: 192.168.0.0/16
      encapsulation: VXLAN
      natOutgoing: Enabled
      nodeSelector: all()
EOF


# Check operator pods
kubectl get pods -n tigera-operator

# Check calico pods
kubectl get pods -n calico-system

# Nodes should become Ready
kubectl get nodes


---
# Create test namespace and pods
kubectl create ns test-ns
kubectl run pod-a -n test-ns --image=nginx --labels=app=web -- sleep 3600
kubectl run pod-b -n test-ns --image=busybox -- sleep 3600

# Test communication (should work)
kubectl exec -n test-ns pod-b -- wget -qO- pod-a

# Apply deny-all policy
cat <<EOF | kubectl apply -f -
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: deny-all
  namespace: test-ns
spec:
  podSelector: {}
  policyTypes:
  - Ingress
EOF

# Test communication again (should fail)
kubectl exec -n test-ns pod-b -- wget -qO- pod-a