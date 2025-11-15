# namespace 
k create ns web

# create role for new user
k -n web create role pod-reader --verb=get,list --resource=pods

#role binding for new user (--user resource does not existes in k8s)
k -n web create rolebinding pod-reader-binding --role=pod-reader --user=carlton

``
$ k -n web get role,rolebinding
NAME                                        CREATED AT
role.rbac.authorization.k8s.io/pod-reader   2022-03-23T21:45:51Z

NAME                                                       ROLE              AGE
rolebinding.rbac.authorization.k8s.io/pod-reader-binding   Role/pod-reader   8s
``

--------------------


Since Kubernetes doesn't have a "user" resource, 
all that's required is a client certificate and key with the common name (CN) to
In our case, when we created the RoleBinding, we assigned it to the user "carlton", 
so that user will assume the permissions from the Role for that resource.

As long as the CN in the key is "carlton", we will be able to use this to access the Kubernetes API.


-----------------------------

To create a private key, we can use the openssl command-line tool. We'll use 2048 bit encryption and we'll name it carlton.key
openssl genrsa -out carlton.key 2048

``
controlplane:~$ openssl genrsa -out carlton.key 2048
controlplane:~$ ls -l 
total 4
-rw------- 1 root root 1704 Aug 26 14:54 carlton.key
lrwxrwxrwx 1 root root    1 Aug 19 08:57 filesystem -> /
controlplane:~$ 
``
Kubernetes itself is a certificate authority, therefore, it can approve and generate certificates

Let's create a Certificate Signing Request (CSR) for the Kubernetes API using our private key and insert the common name and output that to a file named carlton.csr with the following command

``
controlplane:~$ openssl req -new -key carlton.key -subj "/CN=carlton" -out carlton.csr 
controlplane:~$ ll
total 88
``
-rw-r--r--  1 root root  887 Aug 26 14:57 carlton.csr
-rw-------  1 root root 1704 Aug 26 14:54 carlton.key
lrwxrwxrwx  1 root root    1 Aug 19 08:57 filesystem -> //
controlplane:~$ 
``


------------------------


Now that we have a CSR, we can submit it to the Kubernetes API for approval.

before that create a base64 encoded string of our csr
``
controlplane:~$ cat carlton.csr | base64 -w 0
LS0tLS1CRUdJTiBDRVJUSUZJQ0FURSBSRVFVRVNULS0tLS0KTUlJQ1Z6Q0NBVDhDQVFBd0VqRVFNQTRHQTFVRUF3d0hZMkZ5YkhSdmJqQ0NBU0l3RFFZSktvWklodmNOQVFFQgpCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFORDRtNmx5M1o3UXBoSWVwOTdkVmZyM2dIWWtrOEJzRzBEZjNObHJ3cm9LCkF2M2RCYVVnVWhaUjQwa21Sc1N2WnowM3dzOVFFZkJQWThkYmU1WVRtbGJzVHVucHQ1RVNKWE83RUZxZ1k0bUQKdUFoeDhTeWdwWVdKTUxWZ05ZdURBQ3NSeE9ZZ2EwZ0pYZHBITk03d1pEdzBkMjJXTGtZcWU1YUlyL1pIaEFzWQp3RmRtNGFHYlJobVZSUEMrM2Q0QzM2YUVlbHB6VlVDdk5GQ094YTZVbDhxd20wYk1kdmNJZ0hrQlg1YzNHejZBCklKRDMzc3JBbkxrL2hrZTU0Ym1xa2VGcWdHeVJGS2oyVEdkMkVjQkJMeU12N0Qyak9RaElNT0o5Qyt0aXlYWmwKdnBibGpYTW9MZXVjd3BnL3BQeXh2ZU1xS2kvUlpjcEwvZVhJMXVxWjkrY0NBd0VBQWFBQU1BMEdDU3FHU0liMwpEUUVCQ3dVQUE0SUJBUUJ1UlNXVzVuYzBoZFVxOWVuT1VlRzl1bE1xa2R0Q3NpRGJQb0JiTjJQN01Bd2ErdVlVClhHN3lpcmVqTTRQZTZFVWNsSzZaS0lDNFdhejl2dWw5WTdqN28rdHJvcXVnMlArTzM5ZTNpTElxT3pLV25xVkMKSW1Ia3BwM3RPMUVnTnhFbmRoYTMyTVZsdWVxKzMzemcySTE4eVJMd2phUFltLzJ1cXhWelZQSEYzc00zVlF4OQptaVdtbiszUzlEbURvWFJURUMydVlGQkIwbEpjYUVUVnNoTG0yT3F0RGNwSStBYnFRa0VFVkNvUExqRkVlZVZmClFaM1pwd1c2U01pMGdUTW5vT0pKN2k0VGNlRGxBa0grS2hYeVZaUjlDVGlGWURMUDFXbGhoSkV6L1ZtYzMyTzMKUHNMS0lrcGJhOUNVQVlhbi9BT1NUN2JETm5GQlIzVWtYS2tnCi0tLS0tRU5EIENFUlRJRklDQVRFIFJFUVVFU1QtLS0tLQo=
``

Then, we can create a YAML manifest and sumbit it to the Kubernetes API. 

to get the template you can use this below mentioned

``
controlplane:~$ k get csr
NAME        AGE    SIGNERNAME                                    REQUESTOR                  REQUESTEDDURATION   CONDITION
csr-c5vhd   7d5h   kubernetes.io/kube-apiserver-client-kubelet   system:node:controlplane   <none>              Approved,Issued
controlplane:~$ k get csr -o yaml > csr.yml
controlplane:~$ vi csr.yml 
controlplane:~$ 

``

ORRRR

``
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
  name: carlton
spec:
  groups:
  - system:authenticated
  request: <reqest the csr in base64 encoded>
  signerName: kubernetes.io/kube-apiserver-client
  usages:
  - client auth
``

``
controlplane:~$ k apply -f csr.yml 
certificatesigningrequest.certificates.k8s.io/carlton created
controlplane:~$ k get csr
NAME        AGE    SIGNERNAME                                    REQUESTOR                  REQUESTEDDURATION   CONDITION
carlton     2s     kubernetes.io/kube-apiserver-client           kubernetes-admin           <none>              Pending
csr-c5vhd   7d6h   kubernetes.io/kube-apiserver-client-kubelet   system:node:controlplane   <none>              Approved,Issued
controlplane:~$ 

``


In order to get our client certificate that we can use in our kubeconfig, we'll approve the CSR we submitted to the Kubernetes API
k certificate approve carlton
``
controlplane:~$ k certificate approve carlton 
certificatesigningrequest.certificates.k8s.io/carlton approved
controlplane:~$ k get csr
NAME        AGE    SIGNERNAME                                    REQUESTOR                  REQUESTEDDURATION   CONDITION
carlton     113s   kubernetes.io/kube-apiserver-client           kubernetes-admin           <none>              Approved,Issued
csr-c5vhd   7d6h   kubernetes.io/kube-apiserver-client-kubelet   system:node:controlplane   <none>              Approved,Issued
controlplane:~$ 
``

We can extract the client certificate out from the "k get csr" command, decode it and save it to a file named carlton.crt
```
controlplane:~$ k describe csr carlton -o json
error: unknown shorthand flag: 'o' in -o
See 'kubectl describe --help' for usage.
controlplane:~$ k get csr carlton -o json
{
    "apiVersion": "certificates.k8s.io/v1",
    "kind": "CertificateSigningRequest",
    "metadata": {
        "annotations": {
            "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"certificates.k8s.io/v1\",\"kind\":\"CertificateSigningRequest\",\"metadata\":{\"annotations\":{},\"name\":\"carlton\"},\"spec\":{\"groups\":[\"system:authenticated\"],\"request\":\"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURSBSRVFVRVNULS0tLS0KTUlJQ1Z6Q0NBVDhDQVFBd0VqRVFNQTRHQTFVRUF3d0hZMkZ5YkhSdmJqQ0NBU0l3RFFZSktvWklodmNOQVFFQgpCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFORDRtNmx5M1o3UXBoSWVwOTdkVmZyM2dIWWtrOEJzRzBEZjNObHJ3cm9LCkF2M2RCYVVnVWhaUjQwa21Sc1N2WnowM3dzOVFFZkJQWThkYmU1WVRtbGJzVHVucHQ1RVNKWE83RUZxZ1k0bUQKdUFoeDhTeWdwWVdKTUxWZ05ZdURBQ3NSeE9ZZ2EwZ0pYZHBITk03d1pEdzBkMjJXTGtZcWU1YUlyL1pIaEFzWQp3RmRtNGFHYlJobVZSUEMrM2Q0QzM2YUVlbHB6VlVDdk5GQ094YTZVbDhxd20wYk1kdmNJZ0hrQlg1YzNHejZBCklKRDMzc3JBbkxrL2hrZTU0Ym1xa2VGcWdHeVJGS2oyVEdkMkVjQkJMeU12N0Qyak9RaElNT0o5Qyt0aXlYWmwKdnBibGpYTW9MZXVjd3BnL3BQeXh2ZU1xS2kvUlpjcEwvZVhJMXVxWjkrY0NBd0VBQWFBQU1BMEdDU3FHU0liMwpEUUVCQ3dVQUE0SUJBUUJ1UlNXVzVuYzBoZFVxOWVuT1VlRzl1bE1xa2R0Q3NpRGJQb0JiTjJQN01Bd2ErdVlVClhHN3lpcmVqTTRQZTZFVWNsSzZaS0lDNFdhejl2dWw5WTdqN28rdHJvcXVnMlArTzM5ZTNpTElxT3pLV25xVkMKSW1Ia3BwM3RPMUVnTnhFbmRoYTMyTVZsdWVxKzMzemcySTE4eVJMd2phUFltLzJ1cXhWelZQSEYzc00zVlF4OQptaVdtbiszUzlEbURvWFJURUMydVlGQkIwbEpjYUVUVnNoTG0yT3F0RGNwSStBYnFRa0VFVkNvUExqRkVlZVZmClFaM1pwd1c2U01pMGdUTW5vT0pKN2k0VGNlRGxBa0grS2hYeVZaUjlDVGlGWURMUDFXbGhoSkV6L1ZtYzMyTzMKUHNMS0lrcGJhOUNVQVlhbi9BT1NUN2JETm5GQlIzVWtYS2tnCi0tLS0tRU5EIENFUlRJRklDQVRFIFJFUVVFU1QtLS0tLQo=\",\"signerName\":\"kubernetes.io/kube-apiserver-client\",\"usages\":[\"client auth\"]}}\n"
        },
        "creationTimestamp": "2025-08-26T15:05:43Z",
        "name": "carlton",
        "resourceVersion": "4831",
        "uid": "397927db-be5f-429d-ac32-bb731906a049"
    },
    "spec": {
        "extra": {
            "authentication.kubernetes.io/credential-id": [
                "X509SHA256=5b6a51b9e44dfe3021ef089d089fb1c09dc801391343dee2370a3f990604445d"
            ]
        },
        "groups": [
            "kubeadm:cluster-admins",
            "system:authenticated"
        ],
        "request": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURSBSRVFVRVNULS0tLS0KTUlJQ1Z6Q0NBVDhDQVFBd0VqRVFNQTRHQTFVRUF3d0hZMkZ5YkhSdmJqQ0NBU0l3RFFZSktvWklodmNOQVFFQgpCUUFEZ2dFUEFEQ0NBUW9DZ2dFQkFORDRtNmx5M1o3UXBoSWVwOTdkVmZyM2dIWWtrOEJzRzBEZjNObHJ3cm9LCkF2M2RCYVVnVWhaUjQwa21Sc1N2WnowM3dzOVFFZkJQWThkYmU1WVRtbGJzVHVucHQ1RVNKWE83RUZxZ1k0bUQKdUFoeDhTeWdwWVdKTUxWZ05ZdURBQ3NSeE9ZZ2EwZ0pYZHBITk03d1pEdzBkMjJXTGtZcWU1YUlyL1pIaEFzWQp3RmRtNGFHYlJobVZSUEMrM2Q0QzM2YUVlbHB6VlVDdk5GQ094YTZVbDhxd20wYk1kdmNJZ0hrQlg1YzNHejZBCklKRDMzc3JBbkxrL2hrZTU0Ym1xa2VGcWdHeVJGS2oyVEdkMkVjQkJMeU12N0Qyak9RaElNT0o5Qyt0aXlYWmwKdnBibGpYTW9MZXVjd3BnL3BQeXh2ZU1xS2kvUlpjcEwvZVhJMXVxWjkrY0NBd0VBQWFBQU1BMEdDU3FHU0liMwpEUUVCQ3dVQUE0SUJBUUJ1UlNXVzVuYzBoZFVxOWVuT1VlRzl1bE1xa2R0Q3NpRGJQb0JiTjJQN01Bd2ErdVlVClhHN3lpcmVqTTRQZTZFVWNsSzZaS0lDNFdhejl2dWw5WTdqN28rdHJvcXVnMlArTzM5ZTNpTElxT3pLV25xVkMKSW1Ia3BwM3RPMUVnTnhFbmRoYTMyTVZsdWVxKzMzemcySTE4eVJMd2phUFltLzJ1cXhWelZQSEYzc00zVlF4OQptaVdtbiszUzlEbURvWFJURUMydVlGQkIwbEpjYUVUVnNoTG0yT3F0RGNwSStBYnFRa0VFVkNvUExqRkVlZVZmClFaM1pwd1c2U01pMGdUTW5vT0pKN2k0VGNlRGxBa0grS2hYeVZaUjlDVGlGWURMUDFXbGhoSkV6L1ZtYzMyTzMKUHNMS0lrcGJhOUNVQVlhbi9BT1NUN2JETm5GQlIzVWtYS2tnCi0tLS0tRU5EIENFUlRJRklDQVRFIFJFUVVFU1QtLS0tLQo=",
        "signerName": "kubernetes.io/kube-apiserver-client",
        "usages": [
            "client auth"
        ],
        "username": "kubernetes-admin"
    },
    "status": {
        "certificate": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUMrRENDQWVDZ0F3SUJBZ0lSQVBkRDd1ZUFsSzhOV2htckJtT1hKUFF3RFFZSktvWklodmNOQVFFTEJRQXcKRlRFVE1CRUdBMVVFQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TlRBNE1qWXhOVEF5TXpOYUZ3MHlOakE0TWpZeApOVEF5TXpOYU1CSXhFREFPQmdOVkJBTVRCMk5oY214MGIyNHdnZ0VpTUEwR0NTcUdTSWIzRFFFQkFRVUFBNElCCkR3QXdnZ0VLQW9JQkFRRFErSnVwY3QyZTBLWVNIcWZlM1ZYNjk0QjJKSlBBYkJ0QTM5elphOEs2Q2dMOTNRV2wKSUZJV1VlTkpKa2JFcjJjOU44TFBVQkh3VDJQSFczdVdFNXBXN0U3cDZiZVJFaVZ6dXhCYW9HT0pnN2dJY2ZFcwpvS1dGaVRDMVlEV0xnd0FyRWNUbUlHdElDVjNhUnpUTzhHUThOSGR0bGk1R0tudVdpSy8yUjRRTEdNQlhadUdoCm0wWVpsVVR3dnQzZUF0K21oSHBhYzFWQXJ6UlFqc1d1bEpmS3NKdEd6SGIzQ0lCNUFWK1hOeHMrZ0NDUTk5N0sKd0p5NVA0Wkh1ZUc1cXBIaGFvQnNrUlNvOWt4bmRoSEFRUzhqTCt3OW96a0lTRERpZlF2cllzbDJaYjZXNVkxegpLQzNybk1LWVA2VDhzYjNqS2lvdjBXWEtTLzNseU5icW1mZm5BZ01CQUFHalJqQkVNQk1HQTFVZEpRUU1NQW9HCkNDc0dBUVVGQndNQ01Bd0dBMVVkRXdFQi93UUNNQUF3SHdZRFZSMGpCQmd3Rm9BVW9xQWpMbW8yVXB3SFMxbVIKTUpWc2x6K3lEUDB3RFFZSktvWklodmNOQVFFTEJRQURnZ0VCQURGUStuVmhFYUpSYVRWNVppNWlOQXZhc2JHegoyRnZHOUxQcnFCdGNqaFR6SldhNk01MGQzaHN6Nm9SQlpmbkV0OWpVSmx5dUYydFdTQlduUG80clloMTZYdUtCClNXdVdyZW1OdFRid21QYTZiZVhMN2hXRjN6RlF0U1hnOFNrTmVIUEFEZC9PdDA4SWFNV0VCWE5lTSt5UnhmaUkKV25hZ0dqeVJjQkJ3VWNiUDI1SURZSDhncG12di8vbCs3RHgzcER2eXViMW5xcC8xak9QbFo1U2I2aU5LRGVCVwozUGRkV25nMlVHbXc5OWRCT0ZwZjhCaW8xb2p6ekJCdENNclVtaGJZT1kyeDNsRkhDQWZuT243NHY0dXZ4bmZkCjdvci9uekk4MllSMUdTaGZ0WW81U2RCekMwdktKODRZR2R6U2dzcHBieUNaS2xOT1FOZzNCcmMzN1dFPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==",
        "conditions": [
            {
                "lastTransitionTime": "2025-08-26T15:07:33Z",
                "lastUpdateTime": "2025-08-26T15:07:33Z",
                "message": "This CSR was approved by kubectl certificate approve.",
                "reason": "KubectlApprove",
                "status": "True",
                "type": "Approved"
            }
        ]
    }
}
```


controlplane:~$ k get csr carlton -o jsonpath='{.status.certificate}' | base64 -d > carlton.crt
``
controlplane:~$ cat carlton.crt 
-----BEGIN CERTIFICATE-----
MIIC+DCCAeCgAwIBAgIRAPdD7ueAlK8NWhmrBmOXJPQwDQYJKoZIhvcNAQELBQAw
FTETMBEGA1UEAxMKa3ViZXJuZXRlczAeFw0yNTA4MjYxNTAyMzNaFw0yNjA4MjYx
NTAyMzNaMBIxEDAOBgNVBAMTB2Nhcmx0b24wggEiMA0GCSqGSIb3DQEBAQUAA4IB
DwAwggEKAoIBAQDQ+Jupct2e0KYSHqfe3VX694B2JJPAbBtA39zZa8K6CgL93QWl
IFIWUeNJJkbEr2c9N8LPUBHwT2PHW3uWE5pW7E7p6beREiVzuxBaoGOJg7gIcfEs
oKWFiTC1YDWLgwArEcTmIGtICV3aRzTO8GQ8NHdtli5GKnuWiK/2R4QLGMBXZuGh
m0YZlUTwvt3eAt+mhHpac1VArzRQjsWulJfKsJtGzHb3CIB5AV+XNxs+gCCQ997K
wJy5P4ZHueG5qpHhaoBskRSo9kxndhHAQS8jL+w9ozkISDDifQvrYsl2Zb6W5Y1z
KC3rnMKYP6T8sb3jKiov0WXKS/3lyNbqmffnAgMBAAGjRjBEMBMGA1UdJQQMMAoG
CCsGAQUFBwMCMAwGA1UdEwEB/wQCMAAwHwYDVR0jBBgwFoAUoqAjLmo2UpwHS1mR
MJVslz+yDP0wDQYJKoZIhvcNAQELBQADggEBADFQ+nVhEaJRaTV5Zi5iNAvasbGz
2FvG9LPrqBtcjhTzJWa6M50d3hsz6oRBZfnEt9jUJlyuF2tWSBWnPo4rYh16XuKB
SWuWremNtTbwmPa6beXL7hWF3zFQtSXg8SkNeHPADd/Ot08IaMWEBXNeM+yRxfiI
WnagGjyRcBBwUcbP25IDYH8gpmvv//l+7Dx3pDvyub1nqp/1jOPlZ5Sb6iNKDeBW
3PddWng2UGmw99dBOFpf8Bio1ojzzBBtCMrUmhbYOY2x3lFHCAfnOn74v4uvxnfd
7or/nzI82YR1GShftYo5SdBzC0vKJ84YGdzSgsppbyCZKlNOQNg3Brc37WE=
-----END CERTIFICATE-----
controlplane:~$ 
``

Now that we have the key and certificate, we can set the credentials in our kubeconfig and embed the certs within

controlplane:~$ k config set-credentials carlton --client-key=carlton.key --client-certificate=carlton.crt --embed-certs 
User "carlton" set.
controlplane:~$ 

The output of k config view will now show carlton as one of the users

``
controlplane:~$ k config view 
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: DATA+OMITTED
    server: https://172.30.1.2:6443
  name: kubernetes
contexts:
- context:
    cluster: kubernetes
    user: kubernetes-admin
  name: kubernetes-admin@kubernetes
current-context: kubernetes-admin@kubernetes
kind: Config
preferences: {}
users:
- name: carlton
  user:
    client-certificate-data: DATA+OMITTED
    client-key-data: DATA+OMITTED
- name: kubernetes-admin
  user:
    client-certificate-data: DATA+OMITTED
    client-key-data: DATA+OMITTED
controlplane:~$ 
``

The output of k config view will now show carlton as one of the users

Next, we'll set and use the context in which kubectl uses to access the Kubernetes API
``
controlplane:~$ k config set-context carlton --user=carlton --cluster=kubernetes
Context "carlton" created.
controlplane:~$  k config use-context carlton
``

so now carlton can able to access the pods in web  namespace but not for another namespaces 

``
controlplane:~$ k get pods -n kube-system
Error from server (Forbidden): pods is forbidden: User "carlton" cannot list resource "pods" in API group "" in the namespace "kube-system"
controlplane:~$ k get pods -n web        
NAME   READY   STATUS    RESTARTS   AGE
pod1   1/1     Running   0          23m
controlplane:~$ k get all -A
Error from server (Forbidden): pods is forbidden: User "carlton" cannot list resource "pods" in API group "" at the cluster scope
Error from server (Forbidden): replicationcontrollers is forbidden: User "carlton" cannot list resource "replicationcontrollers" in API group "" at the cluster scope
Error from server (Forbidden): services is forbidden: User "carlton" cannot list resource "services" in API group "" at the cluster scope
Error from server (Forbidden): daemonsets.apps is forbidden: User "carlton" cannot list resource "daemonsets" in API group "apps" at the cluster scope
Error from server (Forbidden): deployments.apps is forbidden: User "carlton" cannot list resource "deployments" in API group "apps" at the cluster scope
Error from server (Forbidden): replicasets.apps is forbidden: User "carlton" cannot list resource "replicasets" in API group "apps" at the cluster scope
Error from server (Forbidden): statefulsets.apps is forbidden: User "carlton" cannot list resource "statefulsets" in API group "apps" at the cluster scope
Error from server (Forbidden): horizontalpodautoscalers.autoscaling is forbidden: User "carlton" cannot list resource "horizontalpodautoscalers" in API group "autoscaling" at the cluster scope
Error from server (Forbidden): cronjobs.batch is forbidden: User "carlton" cannot list resource "cronjobs" in API group "batch" at the cluster scope
Error from server (Forbidden): jobs.batch is forbidden: User "carlton" cannot list resource "jobs" in API group "batch" at the cluster scope
controlplane:~$ 

``