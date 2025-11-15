# namespace

k create ns web

# create role for new user

k -n web create role pod-reader --verb=get,list --resource=pods

#role binding for new user (--user resource does not existes in k8s) k
-n web create rolebinding pod-reader-binding --role=pod-reader
--user=carlton

\`\` \$ k -n web get role,rolebinding NAME CREATED AT
role.rbac.authorization.k8s.io/pod-reader 2022-03-23T21:45:51Z

NAME ROLE AGE rolebinding.rbac.authorization.k8s.io/pod-reader-binding
Role/pod-reader 8s \`\`

------------------------------------------------------------------------

Since Kubernetes doesn't have a "user" resource, all that's required is
a client certificate and key with the common name (CN) to In our case,
when we created the RoleBinding, we assigned it to the user "carlton",
so that user will assume the permissions from the Role for that
resource.

As long as the CN in the key is "carlton", we will be able to use this
to access the Kubernetes API.

------------------------------------------------------------------------

To create a private key, we can use the openssl command-line tool. We'll
use 2048 bit encryption and we'll name it carlton.key openssl genrsa
-out carlton.key 2048

`controlplane:~$ openssl genrsa -out carlton.key 2048 controlplane:~$ ls -l  total 4 -rw------- 1 root root 1704 Aug 26 14:54 carlton.key lrwxrwxrwx 1 root root    1 Aug 19 08:57 filesystem -> / controlplane:~$`
Kubernetes itself is a certificate authority, therefore, it can approve
and generate certificates

Let's create a Certificate Signing Request (CSR) for the Kubernetes API
using our private key and insert the common name and output that to a
file named carlton.csr with the following command

`controlplane:~$ openssl req -new -key carlton.key -subj "/CN=carlton" -out carlton.csr  controlplane:~$ ll total 88`
-rw-r--r-- 1 root root 887 Aug 26 14:57 carlton.csr -rw------- 1 root
root 1704 Aug 26 14:54 carlton.key lrwxrwxrwx 1 root root 1 Aug 19 08:57
filesystem -\> // controlplane:\~\$ \`\`

------------------------------------------------------------------------

Now that we have a CSR, we can submit it to the Kubernetes API for
approval.

before that create a base64 encoded string of our csr
`controlplane:~$ cat carlton.csr | base64 -w 0 <base64-csr>`

Then, we can create a YAML manifest and sumbit it to the Kubernetes API.

to get the template you can use this below mentioned

`controlplane:~$ k get csr NAME        AGE    SIGNERNAME                                    REQUESTOR                  REQUESTEDDURATION   CONDITION csr-c5vhd   7d5h   kubernetes.io/kube-apiserver-client-kubelet   system:node:controlplane   <none>              Approved,Issued controlplane:~$ k get csr -o yaml > csr.yml controlplane:~$ vi csr.yml  controlplane:~$`

ORRRR

`apiVersion: certificates.k8s.io/v1 kind: CertificateSigningRequest metadata:   name: carlton spec:   groups:   - system:authenticated   request: <reqest the csr in base64 encoded>   signerName: kubernetes.io/kube-apiserver-client   usages:   - client auth`

`controlplane:~$ k apply -f csr.yml  certificatesigningrequest.certificates.k8s.io/carlton created controlplane:~$ k get csr NAME        AGE    SIGNERNAME                                    REQUESTOR                  REQUESTEDDURATION   CONDITION carlton     2s     kubernetes.io/kube-apiserver-client           kubernetes-admin           <none>              Pending csr-c5vhd   7d6h   kubernetes.io/kube-apiserver-client-kubelet   system:node:controlplane   <none>              Approved,Issued controlplane:~$`

In order to get our client certificate that we can use in our
kubeconfig, we'll approve the CSR we submitted to the Kubernetes API k
certificate approve carlton
`controlplane:~$ k certificate approve carlton  certificatesigningrequest.certificates.k8s.io/carlton approved controlplane:~$ k get csr NAME        AGE    SIGNERNAME                                    REQUESTOR                  REQUESTEDDURATION   CONDITION carlton     113s   kubernetes.io/kube-apiserver-client           kubernetes-admin           <none>              Approved,Issued csr-c5vhd   7d6h   kubernetes.io/kube-apiserver-client-kubelet   system:node:controlplane   <none>              Approved,Issued controlplane:~$`

We can extract the client certificate out from the "k get csr" command,
decode it and save it to a file named carlton.crt

    controlplane:~$ k get csr carlton -o jsonpath='{.status.certificate}' | base64 -d > carlton.crt
    ``
    controlplane:~$ cat carlton.crt 
    -----BEGIN CERTIFICATE-----
    ...
    -----END CERTIFICATE-----
    controlplane:~$ 

Now that we have the key and certificate, we can set the credentials in
our kubeconfig and embed the certs within

    controlplane:~$ k config set-credentials carlton --client-key=carlton.key --client-certificate=carlton.crt --embed-certs 
    User "carlton" set.

The output of k config view will now show carlton as one of the users

    controlplane:~$ k config view 
    ...
    users:
    - name: carlton
      user:
        client-certificate-data: DATA+OMITTED
        client-key-data: DATA+OMITTED
    ...

Next, we'll set and use the context in which kubectl uses to access the
Kubernetes API

    controlplane:~$ k config set-context carlton --user=carlton --cluster=kubernetes
    Context "carlton" created.
    controlplane:~$  k config use-context carlton

so now carlton can able to access the pods in web namespace but not for
another namespaces

    controlplane:~$ k get pods -n kube-system
    Error from server (Forbidden): pods is forbidden: User "carlton" cannot list resource "pods" in API group "" in the namespace "kube-system"
    controlplane:~$ k get pods -n web        
    NAME   READY   STATUS    RESTARTS   AGE
    pod1   1/1     Running   0          23m
