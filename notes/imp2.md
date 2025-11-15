ANsible VM: 10.24.220.34 

# with doc root define
[root@FRPR3FRNT02PR sites-enabled]# cat  ai4u.s2.capgemini.com.conf
<VirtualHost *:80>
    ProxyPreserveHost on
    SSLProxyEngine on
    SSLProxyVerify none
    SSLProxyCheckPeerCN off
    SSLProxyCheckPeerName off
    ServerName ai4u.s2.capgemini.com
    ErrorLog /var/log/httpd/ai4u.s2.capgemini.com-error.log
    CustomLog /var/log/httpd/ai4u.s2.capgemini.com-access.log combined
    DocumentRoot "/var/www/html/ai4u"
	
	<Directory "/www/hub.test.com/www/root">
       allow from all
       Options +Indexes
    </Directory>
 
</VirtualHost>
	 
# kubernetes 

controlplane ~ ✖ ETCDCTL_API=3 etcdctl --endpoints=https://[127.0.0.1]:2379  --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/server.key sna
pshot save /opt/snapshot-pre-boot.db
Snapshot saved at /opt/snapshot-pre-boot.db


controlplane ~ ✖ ETCDCTL_API=3 etcdctl --data-dir=/var/lib/etcd-from-backup \
> snapshot restore /opt/snapshot-pre-boot.db 
2025-03-06 10:26:58.206664 I | mvcc: restore compact to 881
2025-03-06 10:26:58.210447 I | etcdserver/membership: added member 8e9e05c52164694d [http://localhost:2380] to cluster cdf818194e3a8c32



etcd-server ~ ➜  ETCDCTL_API=3 etcdctl \
> --endpoints=https://127.0.0.1:2379 \
> --cacert=/etc/etcd/pki/ca.pem \
> --cert=/etc/etcd/pki/etcd.pem \
> --key=/etc/etcd/pki/etcd-key.pem \
> member list

### refactor this command down below
ETCDCTL_API=3 etcdctl --endpoints=https://127.0.0.1:2379 \
--cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/ \
server.key snapshot save /opt/cluster1.db





student-node ~ ✖ k describe pod etcd-cluster1-controlplane -n kube-system 
Name:                 etcd-cluster1-controlplane
Namespace:            kube-system
Priority:             2000001000
Priority Class Name:  system-node-critical
Node:                 cluster1-controlplane/192.168.57.53
Start Time:           Wed, 12 Mar 2025 08:15:24 +0000
Labels:               component=etcd
                      tier=control-plane
Annotations:          kubeadm.kubernetes.io/etcd.advertise-client-urls: https://192.168.57.53:2379
                      kubernetes.io/config.hash: bb2fb60dd86dfcba6816f6a9f8bc2e1b
                      kubernetes.io/config.mirror: bb2fb60dd86dfcba6816f6a9f8bc2e1b
                      kubernetes.io/config.seen: 2025-03-12T08:15:24.001655930Z
                      kubernetes.io/config.source: file
Status:               Running
SeccompProfile:       RuntimeDefault
IP:                   192.168.57.53
IPs:
  IP:           192.168.57.53
Controlled By:  Node/cluster1-controlplane
Containers:
  etcd:
    Container ID:  containerd://f0a7acb3fbd43e65aa824ce522044b5bea929e605712de0f9fc621fb54801d8a
    Image:         registry.k8s.io/etcd:3.5.10-0
    Image ID:      registry.k8s.io/etcd@sha256:22f892d7672adc0b9c86df67792afdb8b2dc08880f49f669eaaa59c47d7908c2
    Port:          <none>
    Host Port:     <none>
    Command:
      etcd
      --advertise-client-urls=https://192.168.57.53:2379
      --cert-file=/etc/kubernetes/pki/etcd/server.crt
      --client-cert-auth=true
      --data-dir=/var/lib/etcd
      --experimental-initial-corrupt-check=true
      --experimental-watch-progress-notify-interval=5s
      --initial-advertise-peer-urls=https://192.168.57.53:2380
      --initial-cluster=cluster1-controlplane=https://192.168.57.53:2380
      --key-file=/etc/kubernetes/pki/etcd/server.key
      --listen-client-urls=https://127.0.0.1:2379,https://192.168.57.53:2379
      --listen-metrics-urls=http://127.0.0.1:2381
      --listen-peer-urls=https://192.168.57.53:2380
      --name=cluster1-controlplane
      --peer-cert-file=/etc/kubernetes/pki/etcd/peer.crt
      --peer-client-cert-auth=true
      --peer-key-file=/etc/kubernetes/pki/etcd/peer.key
      --peer-trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt
      --snapshot-count=10000
      --trusted-ca-file=/etc/kubernetes/pki/etcd/ca.crt
    State:          Running
      Started:      Wed, 12 Mar 2025 08:15:14 +0000
    Ready:          True
    Restart Count:  0
    Requests:
      cpu:        100m
      memory:     100Mi
    Liveness:     http-get http://127.0.0.1:2381/health%3Fexclude=NOSPACE&serializable=true delay=10s timeout=15s period=10s #success=1 #failure=8
    Startup:      http-get http://127.0.0.1:2381/health%3Fserializable=false delay=10s timeout=15s period=10s #success=1 #failure=24
    Environment:  <none>
    Mounts:
      /etc/kubernetes/pki/etcd from etcd-certs (rw)
      /var/lib/etcd from etcd-data (rw)
Conditions:
  Type                        Status
  PodReadyToStartContainers   True 
  Initialized                 True 
  Ready                       True 
  ContainersReady             True 
  PodScheduled                True 
Volumes:
  etcd-certs:
    Type:          HostPath (bare host directory volume)
    Path:          /etc/kubernetes/pki/etcd
    HostPathType:  DirectoryOrCreate
  etcd-data:
    Type:          HostPath (bare host directory volume)
    Path:          /var/lib/etcd
    HostPathType:  DirectoryOrCreate
QoS Class:         Burstable
Node-Selectors:    <none>
Tolerations:       :NoExecute op=Exists
Events:            <none>



ETCDCTL_API=3 etcdctl --endpoints=https://192.168.57.53:2379 \
--cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/server.key snapshot save /opt/cluster1.db



cluster1-controlplane ~ ➜  ETCDCTL_API=3 etcdctl --endpoints=https://192.160.244.10:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/server.key snapshot save /opt/cluster1.db
Snapshot saved at /opt/cluster1.db

cluster1-controlplane ~ ➜





# haptio ark to take backup of entire k8s cluster
https://github.com/shubheksha/ark





 k create secret docker-registry private-reg-cred --dock
er-username=dock_user --docker-password=dock_password --docker-server=mypriva
teregistry.com:5000 --docker-email=dock_user@myprivateregistry.com
secret/private-reg-cred created



HA Proxy(Louis and Gabin):
FRPR3HAPOCPROD1 ---10.70.88.13
FRPR3HAPOCPROD2 ---10.70.88.11
Reverse Proxy(Louis and gabin):
FRPR3FRNTPLPR01 ---10.70.88.110
FRPR3FRNTPLPR02 ---10.70.88.111
FRPR3FRNTPLPR03 ---10.247.228.194
FRPR3FRNTPLPR04	---10.70.88.54
HA Proxy(Manwe): 
DEFR2HAP01PR    ---10.42.225.119
DEFR2HAP02PR    ---10.42.225.120
Reverse Proxy(Manwe):
DEFR2FRNTPLPR01 ---10.42.225.237
DEFR2FRNTPLPR02 ---10.42.225.226

https://s2plmonitoring.s2-eu.capgemini.com/

backend:
 louis:
	 BalancerMember "http://10.247.243.26:80" connectionTimeout=12000 timeout=12000
       BalancerMember "http://10.247.243.27:80" connectionTimeout=12000 timeout=12000
       BalancerMember "http://10.247.243.28:80" connectionTimeout=12000 timeout=12000
       BalancerMember "http://10.247.243.144:80" connectionTimeout=12000 timeout=12000
       BalancerMember "http://10.247.243.145:80" connectionTimeout=12000 timeout=12000

 gabin:
	ProxyPass / https://10.36.252.38/ nocanon
	ProxyPassReverse / https://10.36.252.38/
	    server oc-manwe-master1 10.42.225.127:6443 check
        server oc-manwe-master2 10.42.225.128:6443 check
        server oc-manwe-master3 10.42.225.129:6443 check



******************************************************************
what are types of env in k8s 
fieldfrom fiedref secret ref etc how many are there explain with example full detaild
*********************************************************************

In Kubernetes, environment variables (env) can be injected into containers in several ways. These are defined in the Pod specification under the env field in a container. Kubernetes provides multiple mechanisms to populate these variables, including:

🔹 Types of env in Kubernetes
1. Static Environment Variables
You can define environment variables with fixed values directly in the Pod spec.

Example:

env:
  - name: APP_MODE
    value: "production"


2. Environment Variables from ConfigMap (configMapRef)
You can inject values from a ConfigMap either as individual keys or all keys.

a. Single Key from ConfigMap
env:
  - name: DB_HOST
    valueFrom:
      configMapKeyRef:
        name: my-config
        key: database_host


b. All Keys from ConfigMap

envFrom:
  - configMapRef:
      name: my-config


3. Environment Variables from Secret (secretRef)
Similar to ConfigMap, you can inject secrets either individually or all at once.

a. Single Key from Secret
env:
  - name: DB_PASSWORD
    valueFrom:
      secretKeyRef:
        name: db-secret
        key: password


b. All Keys from Secret

envFrom:
  - secretRef:
      name: db-secret

4. Environment Variables from Pod Fields (fieldRef)
You can reference fields from the Pod itself, such as its name, namespace, IP, etc.

Example:
env:
  - name: POD_NAME
    valueFrom:
      fieldRef:
        fieldPath: metadata.name
		
Other common fieldPath values:
metadata.namespace
status.podIP
spec.nodeName


5. Environment Variables from Resource Fields (resourceFieldRef)
You can expose container resource limits/requests as environment variables.

Example:
env:
  - name: CPU_LIMIT
    valueFrom:
      resourceFieldRef:
        resource: limits.cpu


Other options:

limits.memory
requests.cpu
requests.memory



Static	value	Literal	"production"
ConfigMap (single)	configMapKeyRef	ConfigMap	database_host
ConfigMap (all)	configMapRef	ConfigMap	All keys
Secret (single)	secretKeyRef	Secret	password
Secret (all)	secretRef	Secret	All keys
Pod Field	fieldRef	Pod metadata	metadata.name
Resource Field	resourceFieldRef	Container spec	limits.cpu



****************************************************
inject secret into pods
*********************************************

To inject a Kubernetes Secret into a Pod using a volume, you mount the secret as a file or directory inside the container. This is useful when applications expect credentials or configuration as files rather than environment variables.

🔹 Step-by-Step Guide
1. Create a Secret
You can create a secret using kubectl:


kubectl create secret generic my-secret \
  --from-literal=username=admin \
  --from-literal=password=secret123


This creates a secret named my-secret with two keys: username and password.

2. Mount the Secret as a Volume in a Pod
Here’s a full example of a Pod manifest that mounts the secret:

apiVersion: v1
kind: Pod
metadata:
  name: secret-volume-pod
spec:
  containers:
    - name: my-container
      image: busybox
      command: [ "sleep", "3600" ]
      volumeMounts:
        - name: secret-volume
          mountPath: "/etc/secret-data"
          readOnly: true
  volumes:
    - name: secret-volume
      secret:
        secretName: my-secret


3. What Happens Inside the Pod
The secret keys (username, password) become files inside /etc/secret-data.
You can read them like this

cat /etc/secret-data/username
cat /etc/secret-data/password


Optional Configurations
a. Specify File Permissions

secret:
  secretName: my-secret
  defaultMode: 0400


b. Map Specific Keys to Specific Paths
secret:
  secretName: my-secret
  items:
    - key: username
      path: user.txt
    - key: password
      path: pass.txt


This will create:

/etc/secret-data/user.txt
/etc/secret-data/pass.txt

