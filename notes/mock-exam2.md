MOCK == 2

# 📘 CKA Practice Questions & Answers

---

## Q1. StorageClass
**Task**  
Create a StorageClass named `local-sc` with the following specifications and set it as the default storage class:  
- Provisioner: `kubernetes.io/no-provisioner`  
- VolumeBindingMode: `WaitForFirstConsumer`  
- Volume expansion enabled  

**Answer**  
```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: local-sc
  annotations:
    storageclass.kubernetes.io/is-default-class: "true"
provisioner: kubernetes.io/no-provisioner
volumeBindingMode: WaitForFirstConsumer
allowVolumeExpansion: true
```

---

## Q2. Logging Deployment with Sidecar
**Task**  
Create a deployment named `logging-deployment` in the namespace `logging-ns` with 1 replica, with:  
- Main container `app-container` (image: busybox) running:  
  ```
  sh -c "while true; do echo 'Log entry' >> /var/log/app/app.log; sleep 5; done"
  ```
- Sidecar container `log-agent` (image: busybox) running:  
  ```
  tail -f /var/log/app/app.log
  ```

**Answer**  
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: logging-deployment
  namespace: logging-ns
spec:
  replicas: 1
  selector:
    matchLabels:
      app: logger
  template:
    metadata:
      labels:
        app: logger
    spec:
      containers:
      - name: app-container
        image: busybox
        command: ["sh", "-c", "while true; do echo 'Log entry' >> /var/log/app/app.log; sleep 5; done"]
        volumeMounts:
        - name: log-volume
          mountPath: /var/log/app
      - name: log-agent
        image: busybox
        command: ["sh", "-c", "tail -f /var/log/app/app.log"]
        volumeMounts:
        - name: log-volume
          mountPath: /var/log/app
      volumes:
      - name: log-volume
        emptyDir: {}
```

---

## Q3. Ingress for WebApp
**Task**  
A Deployment `webapp-deploy` in namespace `ingress-ns` is exposed via Service `webapp-svc`.  
Create an Ingress `webapp-ingress` to:  
- Use `pathType: Prefix`  
- Route `/` → backend service `webapp-svc:80`  
- Host: `kodekloud-ingress.app`  

**Answer**  
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: webapp-ingress
  namespace: ingress-ns
spec:
  rules:
  - host: kodekloud-ingress.app
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: webapp-svc
            port:
              number: 80
```

---

## Q4. Nginx Deployment Rolling Update
**Task**  
Create a new deployment `nginx-deploy` with image `nginx:1.16` and 1 replica. Upgrade to `nginx:1.17` using rolling update.  

**Answer**  
```bash
kubectl create deployment nginx-deploy --image=nginx:1.16 --replicas=1
kubectl set image deployment/nginx-deploy nginx=nginx:1.17
```

---

## Q5. User & Role (CSR)
**Task**  
Create user `john` with CSR `john-developer`. Grant role `developer` in namespace `development` to manage pods. Key: `/root/CKA/john.key`, CSR: `/root/CKA/john.csr`.  

**Answer**  
CSR manifest:  
```yaml
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
  name: john-developer
spec:
  signerName: kubernetes.io/kube-apiserver-client
  request: <BASE64_ENCODED_CSR>
  usages:
  - digital signature
  - key encipherment
  - client auth
```
Approve:  
```bash
kubectl certificate approve john-developer
kubectl create role developer --resource=pods --verb=create,list,get,update,delete --namespace=development
kubectl create rolebinding developer-role-binding --role=developer --user=john --namespace=development
```

---

## Q6. DNS Resolver Test
**Task**  
Create pod `nginx-resolver` with Service `nginx-resolver-service`. Verify DNS resolution using `busybox:1.28`. Save results:  
- `/root/CKA/nginx.svc` (service lookup)  
- `/root/CKA/nginx.pod` (pod lookup)  

**Answer**  
```bash
kubectl run nginx-resolver --image=nginx --restart=Never
kubectl expose pod nginx-resolver --name=nginx-resolver-service --port=80
kubectl run dns-test --image=busybox:1.28 -it --restart=Never -- nslookup nginx-resolver-service > /root/CKA/nginx.svc
kubectl run dns-test2 --image=busybox:1.28 -it --restart=Never -- nslookup nginx-resolver > /root/CKA/nginx.pod
```

---

## Q7. Static Pod
**Task**  
Create a static pod `nginx-critical` on `node01` using image nginx. Place under `/etc/kubernetes/manifests`.  

**Answer**  
```yaml
# /etc/kubernetes/manifests/nginx-critical.yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx-critical
  namespace: kube-system
spec:
  containers:
  - name: nginx
    image: nginx
```

---

## Q8. HPA for Backend
**Task**  
Create HPA `backend-hpa` for Deployment `backend-deployment` in namespace `backend`.  
- Min replicas: 3, Max: 15  
- Scale on memory utilization avg 65%  

**Answer**  
```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: backend-hpa
  namespace: backend
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: backend-deployment
  minReplicas: 3
  maxReplicas: 15
  metrics:
  - type: Resource
    resource:
      name: memory
      target:
        type: Utilization
        averageUtilization: 65
```

---

## Q9. Gateway HTTPS
**Task**  
Modify existing `web-gateway` in namespace `cka5673` to handle HTTPS traffic on port 443 for `kodekloud.com`, using TLS secret `kodekloud-tls`.  

**Answer**  
```yaml
apiVersion: gateway.networking.k8s.io/v1beta1
kind: Gateway
metadata:
  name: web-gateway
  namespace: cka5673
spec:
  listeners:
  - name: https
    protocol: HTTPS
    port: 443
    hostname: kodekloud.com
    tls:
      mode: Terminate
      certificateRefs:
      - kind: Secret
        name: kodekloud-tls
```

---

## Q10. Helm Vulnerable Image
**Task**  
Find release that uses image `kodekloud/webapp-color:v1` and uninstall it.  

**Answer**  
```bash
helm list -A | grep webapp-color
helm uninstall <release-name> -n <namespace>
```

---

## Q11. NetworkPolicy Frontend → Backend
**Task**  
Allow traffic only from `frontend` namespace to backend pods in `backend` namespace. Deny databases namespace. Apply the most restrictive policy from `/root` folder.  

**Answer**  
```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: allow-frontend
  namespace: backend
spec:
  podSelector:
    matchLabels:
      app: backend
  ingress:
  - from:
    - namespaceSelector:
        matchLabels:
          name: frontend
  policyTypes:
  - Ingress
```





-------------
from youtube -- CKA Exam Important question

----

@VenkateshBeemineni
2 months ago
Thanks, Jay, for the great CKA 2025 videos! I do have a question though—do you think simply setting kernel parameters alone isn't enough if the br_netfilter module isn't loaded. That module is crucial for Kubernetes networking, especially when using iptables-based proxies like kube-proxy.


- Even if you set:

sysctl -w net.bridge.bridge-nf-call-iptables=1

it won’t take effect unless the br_netfilter module is actually loaded.

so my answer is #dpkg -i <package name>
systemctl enable cri-docker; systemctl start cri-docker
modprobe br_netfilter
echo "br_netfilter" | sudo tee /etc/modules-load.d/br_netfilter.conf
cd /etc/sysctl.d/<some name>.conf
"put kernel parameters"
systctl --system 
----
---- HPA (video 3)
	for HPA, if you use impartive commands then it genrates template with apps/v1 but to actually use behaviour in HPA you need V2.(V1 does not support behaviour)
	apiVersion: autoscaling/v2
	
---- Ingress
	If your ingress controller doesn’t have a corresponding IngressClass, Kubernetes may ignore the Ingress resource or fail to route traffic. For Kubernetes v1.19 and above, yes,
	in exam ingresss controller should be exposed tru an nodeport service so just check the port once it got in ingress-nginx namespace. 
	and the nodeIP+nodeport/echo will result the page if we do curl on it
	if the curl commands does not provide you the expected output then try to use curl header called -H "Host: example.org" (just example) this should resolve within cluster normally it tries to find the page in google or somewhare
	point is that our ingress relies on host headers
	
---- Network policy support
	 ⁠@ if we need a solution that supports the  network policy enforcement then we can go with calico as flannel doesn’t support the enforcement. will release one soon. 
	if we wanted to test the connection of calico then we can run 2 pods and try to ping from pod1 to pod2 ip 

	THISE SUPPORTS NETWORK policy
	1. Antrea
	2. Calico
	3. Cilium
	4. Kube-router
	5. Romana
	6. Weave Net

---- choose right CNI
	to choose a right cni we can see the exesting cni installed in the cluser by seeing k get pods -A like that 
	Install calico
	By the way, yes, in the exam you’ll typically be given just the calico-operator.yaml URL and no custom resource file. 
	The process remains mostly the same — you just need to replace calico-operator.yaml with the custom-resources.yaml in same url and download files  to your local exam environment and apply them accordingly.😊 
	IMP
	In the CKA exam, if asked to install Calico, you will usually only be given the tigera-operator.yaml URL and not the custom-resources.yaml. The tigera-operator.yaml only installs the operator, but networking will not work until you also apply the custom-resources.yaml. Both are always required. The trick is: take the same base URL where you got tigera-operator.yaml and replace tigera-operator.yaml with custom-resources.yaml. Example: kubectl apply -f https://docs.projectcalico.org/manifests/tigera-operator.yaml && kubectl apply -f https://docs.projectcalico.org/manifests/custom-resources.yaml. In the exam, the cluster usually has no CNI preinstalled, so you must apply both files yourself, otherwise Pods will remain in ContainerCreating state. Final rule: apply operator + custom-resources = working Calico CNI.
	kubectl apply -f https://raw.githubusercontent.com/projectcalico/calico/v3.30.3/manifests/tigera-operator.yaml
	kubectl apply -f https://raw.githubusercontent.com/projectcalico/calico/v3.30.3/manifests/custom-resources.yaml
	
----- install argocd wiout installing CRD's its already installed in cluster (using helm template)
	helm template my-release argo/argo-cd --version 7.7.3 --namespace argocd > argo-helm.yaml
kubectl apply -f argo-helm.yaml

🔹 Step-by-Step Solution

Add Helm repo

helm repo add argo https://argoproj.github.io/argo-helm
helm repo update
kubectl create ns argocd

helm template argocd argo/argo-cd \
  --version 7.7.3 \
  --namespace argocd \
  --set installCRDs=false \
  > /argo-helm.yaml
  
helm install argocd argo/argo-cd \
  --version 7.7.3 \
  --namespace argocd \
  --set installCRDs=false

kubectl get pods -n argocd

----- prority classs
	the exams asks you to patch the deployment after creating priorty class. to patch the deployment they ment to use kubectl patch command otherwise the greader script may reduce the marks
	so betterway writing entire json into single liner command we can use kubectl patch deployment deployment-name --type=merge --patch-file=file
	
vi patch.yaml
`
spec:
  template:
    spec:
      priorityClassName: high-priority

`
kubectl patch deployment busybox-logger -n priority --type=merge --patch-file=patch.yaml


rather than wasting time in this (more error to prone)
kubectl patch deploy busybox-logger -n priority --type=merge -p '{"spec":{"template":{"spec":{"priorityClassName":"high-priority"}}}}'


---- memory/cpu devide accross pods
	eg. memory=2000mi, cpu=2, buffer=15% == 0.15
	echo "2000*0.15" | bc -l == 300.00ki this will get buffer MEMORY value
	echo "2*0.15" | bc -l == .30m this will get buffer cpu value
	-------
	echo "2000-300.00" | bc -l == 1700.00ki this will we can use if we have 3 pods we can devide this using 3
	echo "2-.30" | bc -l == 1.70m this will we can use if we have 3 pods we can devide this using 3
	-------
	echo "1.70/3" | bc -l = 
	
	
	
----- list all service name in the system
systemctl list-unit-files --type service --all | grep -i kube > services.csv



---- testing

Use the CLI tool that will allow you to view the client certificate that the kubelet uses to authenticate to the Kubernetes API. Output the results to a file named “kubelet-config.txt”.

Solution

# view the config that kubelet uses to authenticate to the Kubernetes API
cat /etc/kubernetes/kubelet.conf > kubelet-config.txt

# view the certificate using openssl. Get the certificate file location from the 'kubelet.conf' file above. 
openssl x509 -in /var/lib/kubelet/pki/kubelet-client-current.pem -text -noout

---- gateway api install in local



---------------------------------------

Solve this question on: ssh cka9412

You're asked to extract the following information out of kubeconfig file /opt/course/1/kubeconfig on cka9412:

    Write all kubeconfig context names into /opt/course/1/contexts, one per line

    Write the name of the current context into /opt/course/1/current-context

    Write the client-certificate of user account-0027 base64-decoded into /opt/course/1/cert

-----


Solve this question on: ssh cka7968

Install the MinIO Operator using Helm in Namespace minio. Then configure and create the Tenant CRD:

    Create Namespace minio

    Install Helm chart minio/operator into the new Namespace. The Helm Release should be called minio-operator

    Update the Tenant resource in /opt/course/2/minio-tenant.yaml to include enableSFTP: true under features

    Create the Tenant resource from /opt/course/2/minio-tenant.yaml

    ℹ️ It is not required for MinIO to run properly. Installing the Helm Chart and the Tenant resource as requested is enough

--------------


Solve this question on: ssh cka3962

There are two Pods named o3db-* in Namespace project-h800. The Project H800 management asked you to scale these down to one replica to save resources.


---------

Solve this question on: ssh cka2556

Check all available Pods in the Namespace project-c13 and find the names of those that would probably be terminated first if the nodes run out of resources (cpu or memory).

Write the Pod names into /opt/course/4/pods-terminated-first.txt.

------------

Solve this question on: ssh cka5774

Previously the application api-gateway used some external autoscaler which should now be replaced with a HorizontalPodAutoscaler (HPA). The application has been deployed to Namespaces api-gateway-staging and api-gateway-prod like this:

kubectl kustomize /opt/course/5/api-gateway/staging | kubectl apply -f -
kubectl kustomize /opt/course/5/api-gateway/prod | kubectl apply -f -

Using the Kustomize config at /opt/course/5/api-gateway do the following:

    Remove the ConfigMap horizontal-scaling-config completely
    Add HPA named api-gateway for the Deployment api-gateway with min 2 and max 4 replicas. It should scale at 50% average CPU utilisation
    In prod the HPA should have max 6 replicas
    Apply your changes for staging and prod so they're reflected in the cluster



-----
Solve this question on: ssh cka7968

Create a new PersistentVolume named safari-pv. It should have a capacity of 2Gi, accessMode ReadWriteOnce, hostPath /Volumes/Data and no storageClassName defined.

Next create a new PersistentVolumeClaim in Namespace project-t230 named safari-pvc . It should request 2Gi storage, accessMode ReadWriteOnce and should not define a storageClassName. The PVC should bound to the PV correctly.

Finally create a new Deployment safari in Namespace project-t230 which mounts that volume at /tmp/safari-data. The Pods of that Deployment should be of image httpd:2-alpine.

----

Solve this question on: ssh cka5774

The metrics-server has been installed in the cluster. Write two bash scripts which use kubectl:

    Script /opt/course/7/node.sh should show resource usage of nodes
    Script /opt/course/7/pod.sh should show resource usage of Pods and their containers

----

Solve this question on: ssh cka7968

Create a new PersistentVolume named safari-pv. It should have a capacity of 2Gi, accessMode ReadWriteOnce, hostPath /Volumes/Data and no storageClassName defined.

Next create a new PersistentVolumeClaim in Namespace project-t230 named safari-pvc . It should request 2Gi storage, accessMode ReadWriteOnce and should not define a storageClassName. The PVC should bound to the PV correctly.

Finally create a new Deployment safari in Namespace project-t230 which mounts that volume at /tmp/safari-data. The Pods of that Deployment should be of image httpd:2-alpine.

----
Solve this question on: ssh cka3962

Your coworker notified you that node cka3962-node1 is running an older Kubernetes version and is not even part of the cluster yet.

    Update the node's Kubernetes to the exact version of the controlplane

    Add the node to the cluster using kubeadm

    ℹ️ You can connect to the worker node using ssh cka3962-node1 from cka3962

--------------


Solve this question on: ssh cka9412

There is ServiceAccount secret-reader in Namespace project-swan. Create a Pod of image nginx:1-alpine named api-contact which uses this ServiceAccount.

Exec into the Pod and use curl to manually query all Secrets from the Kubernetes Api.

Write the result into file /opt/course/9/result.json.



---- (killer.sh questions lab1)

Solve this question on: ssh cka3962

Create a new ServiceAccount processor in Namespace project-hamster. Create a Role and RoleBinding, both named processor as well. These should allow the new SA to only create Secrets and ConfigMaps in that Namespace.


----
Solve this question on: ssh cka2556

Use Namespace project-tiger for the following. Create a DaemonSet named ds-important with image httpd:2-alpine and labels id=ds-important and uuid=18426a0b-5f59-4e10-923f-c0e078e82462. The Pods it creates should request 10 millicore cpu and 10 mebibyte memory. The Pods of that DaemonSet should run on all nodes, also controlplanes.

----
Solve this question on: ssh cka2556

Implement the following in Namespace project-tiger:

    Create a Deployment named deploy-important with 3 replicas
    The Deployment and its Pods should have label id=very-important
    First container named container1 with image nginx:1-alpine
    Second container named container2 with image google/pause
    There should only ever be one Pod of that Deployment running on one worker node, use topologyKey: kubernetes.io/hostname for this

    ℹ️ Because there are two worker nodes and the Deployment has three replicas the result should be that the third Pod won't be scheduled. In a way this scenario simulates the behaviour of a DaemonSet, but using a Deployment with a fixed number of replicas

--
Solve this question on: ssh cka7968

The team from Project r500 wants to replace their Ingress (networking.k8s.io) with a Gateway Api (gateway.networking.k8s.io) solution. The old Ingress is available at /opt/course/13/ingress.yaml.

Perform the following in Namespace project-r500 and for the already existing Gateway:

    Create a new HTTPRoute named traffic-director which replicates the routes from the old Ingress
    Extend the new HTTPRoute with path /auto which redirects to mobile if the User-Agent is exactly mobile and to desktop otherwise

The existing Gateway is reachable at http://r500.gateway:30080 which means your implementation should work for these commands:

curl r500.gateway:30080/desktop
curl r500.gateway:30080/mobile
curl r500.gateway:30080/auto -H "User-Agent: mobile" 
curl r500.gateway:30080/auto

----
Solve this question on: ssh cka9412

Perform some tasks on cluster certificates:

    Check how long the kube-apiserver server certificate is valid using openssl or cfssl. Write the expiration date into /opt/course/14/expiration. Run the kubeadm command to list the expiration dates and confirm both methods show the same one
    Write the kubeadm command that would renew the kube-apiserver certificate into /opt/course/14/kubeadm-renew-certs.sh

----
Solve this question on: ssh cka7968

There was a security incident where an intruder was able to access the whole cluster from a single hacked backend Pod.

To prevent this create a NetworkPolicy called np-backend in Namespace project-snake. It should allow the backend-* Pods only to:

    Connect to db1-* Pods on port 1111
    Connect to db2-* Pods on port 2222

Use the app Pod labels in your policy.

    ℹ️ All Pods in the Namespace run plain Nginx images. This allows simple connectivity tests like: k -n project-snake exec POD_NAME -- curl POD_IP:PORT

    ℹ️ For example, connections from backend-* Pods to vault-* Pods on port 3333 should no longer work

----
Solve this question on: ssh cka5774

The CoreDNS configuration in the cluster needs to be updated:

    Make a backup of the existing configuration Yaml and store it at /opt/course/16/coredns_backup.yaml. You should be able to fast recover from the backup
    Update the CoreDNS configuration in the cluster so that DNS resolution for SERVICE.NAMESPACE.custom-domain will work exactly like and in addition to SERVICE.NAMESPACE.cluster.local

Test your configuration for example from a Pod with busybox:1 image. These commands should result in an IP address:

nslookup kubernetes.default.svc.cluster.local
nslookup kubernetes.default.svc.custom-domain

----
Solve this question on: ssh cka2556

In Namespace project-tiger create a Pod named tigers-reunite of image httpd:2-alpine with labels pod=container and container=pod. Find out on which node the Pod is scheduled. Ssh into that node and find the containerd container belonging to that Pod.

Using command crictl:

    Write the ID of the container and the info.runtimeType into /opt/course/17/pod-container.txt

    Write the logs of the container into /opt/course/17/pod-container.log

    ℹ️ You can connect to a worker node using ssh cka2556-node1 or ssh cka2556-node2 from cka2556

------------------- jadamy youtbe

A kubeadm provisioned cluster was migrated to a new machine. Requires configuration changes to run successfully.
Task:
We need fix a single-node cluster that got broken during machine migration.
Identify the broken cluster components and investigate what caused to break those components.
The decommissioned cluster used an external etcd server.
Next, fix the configuration of all broken cluster components.
Ensure to restart all necessary and components for changes to take effect.
Finally, ensure the cluster, single node and all pods are Ready.

---
A user accidentally deleted the MariaDB Deployment in the mariadb namespace,
which was configured with persistent storage. Your responsibility is to reestablish the Deployment while ensuring data is preserved by reusing the available PersistentVolume.
Task:
A PersistentVolume already exists and is retained for reuse. only one pv exist.
Create a PersistentVolumeClaim (PVC) named mariadb in the mariadb NS with the spec: AccessmodeReadWriteOnce and Storage 250Mi
Edit the MariaDB Deploy file located at ~/mariadb-deploy.yaml to use PVC created in the previous step.
Apply the updated Deployment file to the cluster.
Ensure the MariaDB Deployment is running and Stable

----

We have frontend and backend Deploy in separate NS (frontend and backend). They need to communicate.
Analyze: Inspect the frontend and backend Deployments to understand their communication requirements.
Apply: From the NetworkPolicy YAML files in the ~/netpol folder, choose one to apply. It must:
Allow communication between frontend and backend.
Be as restrictive as possible (least permissive)
Do not delete or change the existing "deny-all" netpol's.
Failure to follow these rules may result in a score reduction or zero.

----












---- (killer.sh questions lab2)

