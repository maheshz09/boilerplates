# etcd cluster backup 

question: Take a backup of etcd on cluster1 and save it on the student-node at the path /opt/cluster1.db



If needed, make sure to set the context to cluster1 (on the student-node):

student-node ~ ➜  kubectl config use-context cluster1
Switched to context "cluster1".

student-node ~ ➜  




------------------------------

student-node ~ ➜  k config use-context cluster2
Switched to context "cluster2".
k 
student-node ~ ➜  k get nodes
NAME                    STATUS   ROLES           AGE   VERSION
cluster2-controlplane   Ready    control-plane   93m   v1.29.0
cluster2-node01         Ready    <none>          92m   v1.29.0

student-node ~ ➜  ssh cluster2-controlplane
Welcome to Ubuntu 22.04.5 LTS (GNU/Linux 5.15.0-1077-gcp x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/pro

This system has been minimized by removing packages and content that are
not required on a system that users do not log into.

To restore this content, you can run the 'unminimize' command.
Last login: Wed Mar 12 09:43:43 2025 from 192.168.231.135

cluster2-controlplane ~ ➜  ps -ef | grep etcd
root        3052    2919  0 08:15 ?        00:02:40 kube-apiserver --advertise-address=192.168.193.168 --allow-privileged=true --authorization-mode=Node,RBAC --client-ca-file=/etc/kubernetes/pki/ca.crt --enable-admission-plugins=NodeRestriction --enable-bootstrap-token-auth=true --etcd-cafile=/etc/kubernetes/pki/etcd/ca.pem --etcd-certfile=/etc/kubernetes/pki/etcd/etcd.pem --etcd-keyfile=/etc/kubernetes/pki/etcd/etcd-key.pem --etcd-servers=https://192.168.81.42:2379 --kubelet-client-certificate=/etc/kubernetes/pki/apiserver-kubelet-client.crt --kubelet-client-key=/etc/kubernetes/pki/apiserver-kubelet-client.key --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname --proxy-client-cert-file=/etc/kubernetes/pki/front-proxy-client.crt --proxy-client-key-file=/etc/kubernetes/pki/front-proxy-client.key --requestheader-allowed-names=front-proxy-client --requestheader-client-ca-file=/etc/kubernetes/pki/front-proxy-ca.crt --requestheader-extra-headers-prefix=X-Remote-Extra- --requestheader-group-headers=X-Remote-Group --requestheader-username-headers=X-Remote-User --secure-port=6443 --service-account-issuer=https://kubernetes.default.svc.cluster.local --service-account-key-file=/etc/kubernetes/pki/sa.pub --service-account-signing-key-file=/etc/kubernetes/pki/sa.key --service-cluster-ip-range=172.20.0.0/16 --tls-cert-file=/etc/kubernetes/pki/apiserver.crt --tls-private-key-file=/etc/kubernetes/pki/apiserver.key
root       18014   17951  0 09:49 pts/0    00:00:00 grep etcd

cluster2-controlplane ~ ➜  ssh 192.168.81.42
The authenticity of host '192.168.81.42 (192.168.81.42)' can't be established.
ED25519 key fingerprint is SHA256:v3S0pu0hiGNBalMbcyUZrNARUIuDlUaTabq5wKlI+TQ.
This key is not known by any other names
Are you sure you want to continue connecting (yes/no/[fingerprint])? yes
Warning: Permanently added '192.168.81.42' (ED25519) to the list of known hosts.
root@192.168.81.42's password: 
Permission denied, please try again.
root@192.168.81.42's password: 


cluster2-controlplane ~ ✖ exit
logout
Connection to cluster2-controlplane closed.

student-node ~ ✖ ssh 192.168.81.42
Welcome to Ubuntu 18.04.6 LTS (GNU/Linux 5.15.0-1075-gcp x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage
This system has been minimized by removing packages and content that are
not required on a system that users do not log into.

To restore this content, you can run the 'unminimize' command.

etcd-server ~ ➜  



etcd-server ~ ➜  



etcd-server ~ ➜  

etcd-server ~ ➜  

etcd-server ~ ➜  

etcd-server ~ ➜  ps -ef | grep etcd
etcd         731       1  0 08:14 ?        00:00:48 /usr/local/bin/etcd --name etcd-server --data-dir=/var/lib/etcd-data --cert-file=/etc/etcd/pki/etcd.pem --key-file=/etc/etcd/pki/etcd-key.pem --peer-cert-file=/etc/etcd/pki/etcd.pem --peer-key-file=/etc/etcd/pki/etcd-key.pem --trusted-ca-file=/etc/etcd/pki/ca.pem --peer-trusted-ca-file=/etc/etcd/pki/ca.pem --peer-client-cert-auth --client-cert-auth --initial-advertise-peer-urls https://192.168.81.42:2380 --listen-peer-urls https://192.168.81.42:2380 --advertise-client-urls https://192.168.81.42:2379 --listen-client-urls https://192.168.81.42:2379,https://127.0.0.1:2379 --initial-cluster-token etcd-cluster-1 --initial-cluster etcd-server=https://192.168.81.42:2380 --initial-cluster-state new
root        1054     881  0 09:50 pts/0    00:00:00 grep etcd

etcd-server ~ ➜  vim /etc
etc/                             etcd-v3.4.20-linux-amd64/        etcd-v3.4.20-linux-amd64.tar.gz

etcd-server ~ ➜  vim /etc
etc/                             etcd-v3.4.20-linux-amd64/        etcd-v3.4.20-linux-amd64.tar.gz

etcd-server ~ ➜  vim /etc
etc/                             etcd-v3.4.20-linux-amd64/        etcd-v3.4.20-linux-amd64.tar.gz

etcd-server ~ ➜  rpm -qa | grep etcd 
-bash: rpm: command not found

etcd-server ~ ✖ ETCDCTL_API=3 etcd member list 
2025-03-12 09:51:46.595587 E | etcdmain: error verifying flags, 'member' is not a valid flag. See 'etcd --help'.

etcd-server ~ ✖ ETCDCTL_API=3 

etcd-server ~ ➜   ETCDCTL_API=3 etcdctl \
> --end^C

etcd-server ~ ✖ ^C

etcd-server ~ ✖ ^C
^C

etcd-server ~ ✖ /etcd-v3.4.20-linux-amd64/^C

etcd-server ~ ✖ ^C

etcd-server ~ ✖ ^C

etcd-server ~ ✖ ls -l
total 0

etcd-server ~ ➜  /etc/etcd/pki/
-bash: /etc/etcd/pki/: Is a directory

etcd-server ~ ✖ ls -l /etc/etcd/pki/
total 12
-rwx------ 1 etcd etcd 1334 Mar 12 08:14 ca.pem
-rwx------ 1 etcd etcd 1679 Mar 12 08:14 etcd-key.pem
-rwx------ 1 etcd etcd 1468 Mar 12 08:14 etcd.pem

etcd-server ~ ➜  ETCDCTL_API=3 etcdctl \
> --endpoints=https://127.0.0.1:2379 \
> --cacert=/etc/etcd/pki/ca.pem \
> --cert=/etc/etcd/pki/etcd.pem \
> --key=/etc/etcd/pki/etcd-key.pem \
> member list
512440828fb187d6, started, etcd-server, https://192.168.81.42:2380, https://192.168.81.42:2379, false

etcd-server ~ ➜  exit
logout
Connection to 192.168.81.42 closed.

student-node ~ ➜  ssh 192.168.81.42
Welcome to Ubuntu 18.04.6 LTS (GNU/Linux 5.15.0-1075-gcp x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage
This system has been minimized by removing packages and content that are
not required on a system that users do not log into.

To restore this content, you can run the 'unminimize' command.
Last login: Wed Mar 12 09:49:36 2025 from 192.168.231.135

etcd-server ~ ➜  exit
logout
Connection to 192.168.81.42 closed.

student-node ~ ➜  k config use-context cluster1
Switched to context "cluster1".

student-node ~ ➜  k get nodes
NAME                    STATUS   ROLES           AGE    VERSION
cluster1-controlplane   Ready    control-plane   102m   v1.29.0
cluster1-node01         Ready    <none>          101m   v1.29.0

student-node ~ ➜  etcd
-su: etcd: command not found

student-node ~ ✖ ETCDCTL_API=3

student-node ~ ➜  etcdctl 
NAME:
   etcdctl - A simple command line client for etcd.

WARNING:
   Environment variable ETCDCTL_API is not set; defaults to etcdctl v2.
   Set environment variable ETCDCTL_API=3 to use v3 API or ETCDCTL_API=2 to use v2 API.

USAGE:
   etcdctl [global options] command [command options] [arguments...]
   
VERSION:
   3.3.13
   
COMMANDS:
     backup          backup an etcd directory
     cluster-health  check the health of the etcd cluster
     mk              make a new key with a given value
     mkdir           make a new directory
     rm              remove a key or a directory
     rmdir           removes the key if it is an empty directory or a key-value pair
     get             retrieve the value of a key
     ls              retrieve a directory
     set             set the value of a key
     setdir          create a new directory or update an existing directory TTL
     update          update an existing key with a given value
     updatedir       update an existing directory
     watch           watch a key for changes
     exec-watch      watch a key for changes and exec an executable
     member          member add, remove and list subcommands
     user            user add, grant and revoke subcommands
     role            role add, grant and revoke subcommands
     auth            overall auth controls
     help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug                          output cURL commands which can be used to reproduce the request
   --no-sync                        don't synchronize cluster information before sending request
   --output simple, -o simple       output response in the given format (simple, `extended` or `json`) (default: "simple")
   --discovery-srv value, -D value  domain name to query for SRV records describing cluster endpoints
   --insecure-discovery             accept insecure SRV records describing cluster endpoints
   --peers value, -C value          DEPRECATED - "--endpoints" should be used instead
   --endpoint value                 DEPRECATED - "--endpoints" should be used instead
   --endpoints value                a comma-delimited list of machine addresses in the cluster (default: "http://127.0.0.1:2379,http://127.0.0.1:4001")
   --cert-file value                identify HTTPS client using this SSL certificate file
   --key-file value                 identify HTTPS client using this SSL key file
   --ca-file value                  verify certificates of HTTPS-enabled servers using this CA bundle
   --username value, -u value       provide username[:password] and prompt if password is not supplied.
   --timeout value                  connection timeout per request (default: 2s)
   --total-timeout value            timeout for the command execution (except watch) (default: 5s)
   --help, -h                       show help
   --version, -v                    print the version
   

student-node ~ ➜  ETCDCTL_API=3 etcdctl \
> --endpoints=htttps://127.0.0.1:2379 \
> --cacert=^C

student-node ~ ✖ ^C

student-node ~ ✖ ^C

student-node ~ ✖ ^C

student-node ~ ✖ ^C

student-node ~ ✖ k get pods A
Error from server (NotFound): pods "A" not found

student-node ~ ✖ k get pods -A
NAMESPACE      NAME                                            READY   STATUS    RESTARTS   AGE
kube-flannel   kube-flannel-ds-2p789                           1/1     Running   0          107m
kube-flannel   kube-flannel-ds-6f7j5                           1/1     Running   0          108m
kube-system    coredns-69f9c977-r4hfp                          1/1     Running   0          107m
kube-system    coredns-69f9c977-shqsk                          1/1     Running   0          107m
kube-system    etcd-cluster1-controlplane                      1/1     Running   0          108m
kube-system    kube-apiserver-cluster1-controlplane            1/1     Running   0          108m
kube-system    kube-controller-manager-cluster1-controlplane   1/1     Running   0          108m
kube-system    kube-proxy-6857c                                1/1     Running   0          107m
kube-system    kube-proxy-jdnmb                                1/1     Running   0          108m
kube-system    kube-scheduler-cluster1-controlplane            1/1     Running   0          108m

student-node ~ ➜  k describe pod -n kube-system etcd-cluster1-controlplane
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

student-node ~ ➜  ETCDCTL_API=3 etcdctl --endpoints=htttps://127.0.0.1:2379 \
> --cacert=/etc/kubernetes/pki/etcd/ca.crt \
> --cert=/etc/kubernetes/pki/etcd/server.crt \
> --key=/etc/kubernetes/pki/etcd/server.key \
> snapshot save /opt/cluster1.db
Error: open /etc/kubernetes/pki/etcd/server.crt: no such file or directory

student-node ~ ✖ ETCDCTL_API=3 etcdctl --endpoints=https://127.0.0.1:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt 
--cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/server.key snapshot save /opt/cluster1.db

Error: open /etc/kubernetes/pki/etcd/server.crt: no such file or directory

student-node ~ ✖ ssh etcd-cluster1-controlplane ETCDCTL_API=3 etcdctl --endpoints=https://127.0.0.1:2379 --cacert=/
etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/server.key
 snapshot save /opt/cluster1.db
ssh: Could not resolve hostname etcd-cluster1-controlplane: Name or service not known

student-node ~ ✖ k exec -it etcd-cluster1-controlplane -- ETCDCTL_API=3 etcdctl --endpoints=https://127.0.0.1:2379 
--cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/
server.key snapshot save /opt/cluster1.db
Error from server (NotFound): pods "etcd-cluster1-controlplane" not found

student-node ~ ✖ k exec -it -n kube-system etcd-cluster1-controlplane -- ETCDCTL_API=3 etcdctl --endpoints=https://
127.0.0.1:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kuber
netes/pki/etcd/server.key snapshot save /opt/cluster1.db
error: Internal error occurred: error executing command in container: failed to exec in container: failed to start exec "53756bf046565f45923aeec5ee0d3e5025c6c5ba3b846c5621c1fdaa80748ed0": OCI runtime exec failed: exec failed: unable to start container process: exec: "ETCDCTL_API=3": executable file not found in $PATH: unknown

student-node ~ ✖ ssh controlplane1-^C

student-node ~ ✖ k get nodes
NAME                    STATUS   ROLES           AGE    VERSION
cluster1-controlplane   Ready    control-plane   112m   v1.29.0
cluster1-node01         Ready    <none>          112m   v1.29.0

student-node ~ ➜  ssh cluster1-controlplane
Welcome to Ubuntu 22.04.5 LTS (GNU/Linux 5.15.0-1077-gcp x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/pro

This system has been minimized by removing packages and content that are
not required on a system that users do not log into.

To restore this content, you can run the 'unminimize' command.
Last login: Wed Mar 12 09:38:50 2025 from 192.168.231.135

cluster1-controlplane ~ ➜  ETCDCTL_API=3 etcdctl --endpoints=https://
127.0.0.1:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kuber 
netes/pki/etcd/server.key snapshot save /opt/cluster1.db
NAME:
        etcdctl - A simple command line client for etcd3.

USAGE:
        etcdctl

VERSION:
        3.3.13

API VERSION:
        3.3


COMMANDS:
        get                     Gets the key or a range of keys
        put                     Puts the given key into the store
        del                     Removes the specified key or range of keys [key, range_end)
        txn                     Txn processes all the requests in one transaction
        compaction              Compacts the event history in etcd
        alarm disarm            Disarms all alarms
        alarm list              Lists all alarms
        defrag                  Defragments the storage of the etcd members with given endpoints
        endpoint health         Checks the healthiness of endpoints specified in `--endpoints` flag
        endpoint status         Prints out the status of endpoints specified in `--endpoints` flag
        endpoint hashkv         Prints the KV history hash for each endpoint in --endpoints
        move-leader             Transfers leadership to another etcd cluster member.
        watch                   Watches events stream on keys or prefixes
        version                 Prints the version of etcdctl
        lease grant             Creates leases
        lease revoke            Revokes leases
        lease timetolive        Get lease information
        lease list              List all active leases
        lease keep-alive        Keeps leases alive (renew)
        member add              Adds a member into the cluster
        member remove           Removes a member from the cluster
        member update           Updates a member in the cluster
        member list             Lists all members in the cluster
        snapshot save           Stores an etcd node backend snapshot to a given file
        snapshot restore        Restores an etcd member snapshot to an etcd directory
        snapshot status         Gets backend snapshot status of a given file
        make-mirror             Makes a mirror at the destination etcd cluster
        migrate                 Migrates keys in a v2 store to a mvcc store
        lock                    Acquires a named lock
        elect                   Observes and participates in leader election
        auth enable             Enables authentication
        auth disable            Disables authentication
        user add                Adds a new user
        user delete             Deletes a user
        user get                Gets detailed information of a user
        user list               Lists all users
        user passwd             Changes password of user
        user grant-role         Grants a role to a user
        user revoke-role        Revokes a role from a user
        role add                Adds a new role
        role delete             Deletes a role
        role get                Gets detailed information of a role
        role list               Lists all roles
        role grant-permission   Grants a key to a role
        role revoke-permission  Revokes a key from a role
        check perf              Check the performance of the etcd cluster
        help                    Help about any command

OPTIONS:
      --cacert=""                               verify certificates of TLS-enabled secure servers using this CA bundle
      --cert=""                                 identify secure client using this TLS certificate file
      --command-timeout=5s                      timeout for short running command (excluding dial timeout)
      --debug[=false]                           enable client-side debug logging
      --dial-timeout=2s                         dial timeout for client connections
  -d, --discovery-srv=""                        domain name to query for SRV records describing cluster endpoints
      --endpoints=[127.0.0.1:2379]              gRPC endpoints
      --hex[=false]                             print byte strings as hex encoded strings
      --insecure-discovery[=true]               accept insecure SRV records describing cluster endpoints
      --insecure-skip-tls-verify[=false]        skip server certificate verification
      --insecure-transport[=true]               disable transport security for client connections
      --keepalive-time=2s                       keepalive time for client connections
      --keepalive-timeout=6s                    keepalive timeout for client connections
      --key=""                                  identify secure client using this TLS key file
      --user=""                                 username[:password] for authentication (prompt if password is not supplied)
  -w, --write-out="simple"                      set the output format (fields, json, protobuf, simple, table)

-bash: 127.0.0.1:2379: command not found
-bash: netes/pki/etcd/server.key: No such file or directory

cluster1-controlplane ~ ✖ netes/pki/etcd/server.key snapshot save /opt^Cluster1.db

cluster1-controlplane ~ ✖ ETCDCTL_API=3 etcdctl --endpoints=https://127.0.0.1:2379 
--cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/ 
server.key snapshot save /opt/cluster1.db
NAME:
        etcdctl - A simple command line client for etcd3.

USAGE:
        etcdctl

VERSION:
        3.3.13

API VERSION:
        3.3


COMMANDS:
        get                     Gets the key or a range of keys
        put                     Puts the given key into the store
        del                     Removes the specified key or range of keys [key, range_end)
        txn                     Txn processes all the requests in one transaction
        compaction              Compacts the event history in etcd
        alarm disarm            Disarms all alarms
        alarm list              Lists all alarms
        defrag                  Defragments the storage of the etcd members with given endpoints
        endpoint health         Checks the healthiness of endpoints specified in `--endpoints` flag
        endpoint status         Prints out the status of endpoints specified in `--endpoints` flag
        endpoint hashkv         Prints the KV history hash for each endpoint in --endpoints
        move-leader             Transfers leadership to another etcd cluster member.
        watch                   Watches events stream on keys or prefixes
        version                 Prints the version of etcdctl
        lease grant             Creates leases
        lease revoke            Revokes leases
        lease timetolive        Get lease information
        lease list              List all active leases
        lease keep-alive        Keeps leases alive (renew)
        member add              Adds a member into the cluster
        member remove           Removes a member from the cluster
        member update           Updates a member in the cluster
        member list             Lists all members in the cluster
        snapshot save           Stores an etcd node backend snapshot to a given file
        snapshot restore        Restores an etcd member snapshot to an etcd directory
        snapshot status         Gets backend snapshot status of a given file
        make-mirror             Makes a mirror at the destination etcd cluster
        migrate                 Migrates keys in a v2 store to a mvcc store
        lock                    Acquires a named lock
        elect                   Observes and participates in leader election
        auth enable             Enables authentication
        auth disable            Disables authentication
        user add                Adds a new user
        user delete             Deletes a user
        user get                Gets detailed information of a user
        user list               Lists all users
        user passwd             Changes password of user
        user grant-role         Grants a role to a user
        user revoke-role        Revokes a role from a user
        role add                Adds a new role
        role delete             Deletes a role
        role get                Gets detailed information of a role
        role list               Lists all roles
        role grant-permission   Grants a key to a role
        role revoke-permission  Revokes a key from a role
        check perf              Check the performance of the etcd cluster
        help                    Help about any command

OPTIONS:
      --cacert=""                               verify certificates of TLS-enabled secure servers using this CA bundle
      --cert=""                                 identify secure client using this TLS certificate file
      --command-timeout=5s                      timeout for short running command (excluding dial timeout)
      --debug[=false]                           enable client-side debug logging
      --dial-timeout=2s                         dial timeout for client connections
  -d, --discovery-srv=""                        domain name to query for SRV records describing cluster endpoints
      --endpoints=[127.0.0.1:2379]              gRPC endpoints
      --hex[=false]                             print byte strings as hex encoded strings
      --insecure-discovery[=true]               accept insecure SRV records describing cluster endpoints
      --insecure-skip-tls-verify[=false]        skip server certificate verification
      --insecure-transport[=true]               disable transport security for client connections
      --keepalive-time=2s                       keepalive time for client connections
      --keepalive-timeout=6s                    keepalive timeout for client connections
      --key=""                                  identify secure client using this TLS key file
      --user=""                                 username[:password] for authentication (prompt if password is not supplied)
  -w, --write-out="simple"                      set the output format (fields, json, protobuf, simple, table)

-bash: --cacert=/etc/kubernetes/pki/etcd/ca.crt: No such file or directory
-bash: server.key: command not found

cluster1-controlplane ~ ✖ server.key snapshot save /opt/cluster1.d^C

cluster1-controlplane ~ ✖ ^C

cluster1-controlplane ~ ✖ ETCDCTL_API=3 etcdctl --endpoints=https://127.0.0.1:2379 \
--cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/ \
server.key snapshot save /opt/cluster1.db
Error: unknown command "server.key" for "etcdctl"
Run 'etcdctl --help' for usage.
Error: unknown command "server.key" for "etcdctl"

cluster1-controlplane ~ ✖ ETCDCTL_API=3

cluster1-controlplane ~ ➜  ETCDCTL_API=3 etcdctl --endpoints=https://127.0.0.1:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/ server.key snapshot save /opt/
cluster1.db
Error: unknown command "server.key" for "etcdctl"
Run 'etcdctl --help' for usage.
Error: unknown command "server.key" for "etcdctl"

cluster1-controlplane ~ ✖ ETCDCTL_API=3 etcdctl --endpoints=https://127.0.0.1:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.cr^C--key=/etc/kubernetes/pki/etcd/ server.key snapshot save /opt/c
luster1.db

cluster1-controlplane ~ ✖ ^C

cluster1-controlplane ~ ✖ ^C

cluster1-controlplane ~ ✖ exit
logout
Connection to cluster1-controlplane closed.

student-node ~ ✖ k get pods 
No resources found in default namespace.

student-node ~ ➜  k get pods -A
NAMESPACE      NAME                                            READY   STATUS    RESTARTS   AGE
kube-flannel   kube-flannel-ds-2p789                           1/1     Running   0          113m
kube-flannel   kube-flannel-ds-6f7j5                           1/1     Running   0          114m
kube-system    coredns-69f9c977-r4hfp                          1/1     Running   0          114m
kube-system    coredns-69f9c977-shqsk                          1/1     Running   0          114m
kube-system    etcd-cluster1-controlplane                      1/1     Running   0          114m
kube-system    kube-apiserver-cluster1-controlplane            1/1     Running   0          114m
kube-system    kube-controller-manager-cluster1-controlplane   1/1     Running   0          114m
kube-system    kube-proxy-6857c                                1/1     Running   0          113m
kube-system    kube-proxy-jdnmb                                1/1     Running   0          114m
kube-system    kube-scheduler-cluster1-controlplane            1/1     Running   0          114m

student-node ~ ➜  ETCDCTL_API=3 etcdctl --endpoints=https://127.0.0.1:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/server.key snapshot save /opt/cluster1.db^C

student-node ~ ✖ ^C

student-node ~ ✖ ^C

student-node ~ ✖ k describe etcd-cluster1-controlplane - n kube-system 
error: the server doesn't have a resource type "etcd-cluster1-controlplane"

student-node ~ ✖ k describe pod etcd-cluster1-controlplane - n kube-system 
Error from server (NotFound): pods "etcd-cluster1-controlplane" not found
Error from server (NotFound): pods "-" not found
Error from server (NotFound): pods "n" not found
Error from server (NotFound): pods "kube-system" not found

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

student-node ~ ➜  ETCDCTL_API=3 etcdctl --endpoints=https://192.168.57.53:2379 \
> --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/server.key snapshot save /opt/cluster1.db
Error: open /etc/kubernetes/pki/etcd/server.crt: no such file or directory

student-node ~ ✖ k get nodes 
NAME                    STATUS   ROLES           AGE    VERSION
cluster1-controlplane   Ready    control-plane   119m   v1.29.0
cluster1-node01         Ready    <none>          119m   v1.29.0

student-node ~ ➜  ^Ch 

student-node ~ ✖ ^C

student-node ~ ✖ ^C

student-node ~ ✖ ^C

student-node ~ ✖ k get pods -A
NAMESPACE      NAME                                            READY   STATUS    RESTARTS   AGE
kube-flannel   kube-flannel-ds-2p789                           1/1     Running   0          119m
kube-flannel   kube-flannel-ds-6f7j5                           1/1     Running   0          119m
kube-system    coredns-69f9c977-r4hfp                          1/1     Running   0          119m
kube-system    coredns-69f9c977-shqsk                          1/1     Running   0          119m
kube-system    etcd-cluster1-controlplane                      1/1     Running   0          120m
kube-system    kube-apiserver-cluster1-controlplane            1/1     Running   0          120m
kube-system    kube-controller-manager-cluster1-controlplane   1/1     Running   0          120m
kube-system    kube-proxy-6857c                                1/1     Running   0          119m
kube-system    kube-proxy-jdnmb                                1/1     Running   0          119m
kube-system    kube-scheduler-cluster1-controlplane            1/1     Running   0          120m

student-node ~ ➜  ssh cluster1-controlplane
Welcome to Ubuntu 22.04.5 LTS (GNU/Linux 5.15.0-1077-gcp x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/pro

This system has been minimized by removing packages and content that are
not required on a system that users do not log into.

To restore this content, you can run the 'unminimize' command.
Last login: Wed Mar 12 10:08:17 2025 from 192.168.231.135

cluster1-controlplane ~ ➜  ETCDCTL_API=3 etcdctl --endpoints=https://192.168.57.53:2379 \
--cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/server.crt --key=/etc/kubernetes/pki/etcd/server.key snapshot save /opt/cluster1.db
Snapshot saved at /opt/cluster1.db

cluster1-controlplane ~ ➜  exit
logout
Connection to cluster1-controlplane closed.

student-node ~ ➜  scp cluster1-controlplane:/opt/cluster1.db /opt/
cluster1.db                                                                      100% 2128KB 189.1MB/s   00:00    

student-node ~ ➜  