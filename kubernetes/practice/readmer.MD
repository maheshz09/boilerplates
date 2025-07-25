      Welcome to the KodeKloud Hands-On lab                                                    
    __ ______  ____  ________ __ __    ____  __  ______ 
   / //_/ __ \/ __ \/ ____/ //_// /   / __ \/ / / / __ \
  / ,< / / / / / / / __/ / ,<  / /   / / / / / / / / / /
 / /| / /_/ / /_/ / /___/ /| |/ /___/ /_/ / /_/ / /_/ / 
/_/ |_\____/_____/_____/_/ |_/_____/\____/\____/_____/  
                                                        
          All rights reserved                                                                  

controlplane ~ ➜  ls -l /
total 68
lrwxrwxrwx    1 root   root       7 Sep 11  2024 bin -> usr/bin
drwxr-xr-x    1 root   root    4096 Apr 18  2022 boot
drwxr-xr-x    6 root   root     440 Mar 29 19:10 dev
drwxr-xr-x    1 root   root    4096 Mar 29 19:10 etc
drwxr-xr-x    1 root   root    4096 Dec 28 17:51 home
lrwxrwxrwx    1 root   root       7 Sep 11  2024 lib -> usr/lib
lrwxrwxrwx    1 root   root       9 Sep 11  2024 lib32 -> usr/lib32
lrwxrwxrwx    1 root   root       9 Sep 11  2024 lib64 -> usr/lib64
lrwxrwxrwx    1 root   root      10 Sep 11  2024 libx32 -> usr/libx32
drwxr-xr-x    1 root   root    4096 Sep 11  2024 media
drwxr-xr-x    1 root   root    4096 Sep 11  2024 mnt
drwxr-xr-x    1 root   root    4096 Mar 29 19:10 opt
dr-xr-xr-x 1530 root   root       0 Mar 29 19:10 proc
drwx------    1 root   root    4096 Mar 29 19:15 root
drwxr-xr-x   18 root   root     500 Mar 29 19:15 run
lrwxrwxrwx    1 root   root       8 Sep 11  2024 sbin -> usr/sbin
drwxr-xr-x    1 root   root    4096 Sep 11  2024 srv
dr-xr-xr-x   13 nobody nogroup    0 Mar 29 19:10 sys
drwxrwxrwt    7 root   root    4096 Mar 29 19:15 tmp
drwxr-xr-x    1 root   root    4096 Sep 11  2024 usr
drwxr-xr-x    1 root   root    4096 Mar 29 19:15 var

controlplane ~ ➜  ls -l /root/
total 8
-rw-r--r-- 1 root root  887 Mar 29 19:15 akshay.csr
-rw------- 1 root root 1704 Mar 29 19:15 akshay.key

controlplane ~ ➜  ls -l
total 8
-rw-r--r-- 1 root root  887 Mar 29 19:15 akshay.csr
-rw------- 1 root root 1704 Mar 29 19:15 akshay.key

controlplane ~ ➜  cat akshay.csr | base64
LS0tLS1CRUdJTiBDRVJUSUZJQ0FURSBSRVFVRVNULS0tLS0KTUlJQ1ZqQ0NBVDRDQVFBd0VURVBN
QTBHQTFVRUF3d0dZV3R6YUdGNU1JSUJJakFOQmdrcWhraUc5dzBCQVFFRgpBQU9DQVE4QU1JSUJD
Z0tDQVFFQWp0aERDak1JUVNUTi83UUdzQ1ZkalFBSzZMdnZmU21LYzVrNjRQamU2QXhUCm1MT0l1
SzFNeElhUGx0WmVrMGIvKzhHU0JQMGtoL2VVQUNaUTVwRDRKRTZ1ODVNYWxtc0tNYUxsQTRRcGdN
NXEKY2d6NzJOUjR0QU0yTkpGYXpkbjYrTjFWM1Fudk9KK3U3d2lQYmg4c21QdVVWOTlBVHY5K1JB
RzJJMFlXckpsQQpYSlh3UGRJalJvcTBIYlZTVlhhUUJkVVpmbnRubUNPd2VlNm9CLzdwVjdic2lu
d1FrajU0ZXg2enAxenRxNkQ5CkJBZ3Y4QmlSdmNzNmQrNEhtMFJXVXJuVms3V1FDM2c1OFRkNHNw
a2IzMlJoTjdRVW5aOTVzMlFYUnVoaThETjYKSW9kVFVZaFRaNnN4OWovWGVMODlIeXBpb212UUh0
TXB4RnJPbFloOGlRSURBUUFCb0FBd0RRWUpLb1pJaHZjTgpBUUVMQlFBRGdnRUJBRm5wRnB1Z3NH
cFlZQTYvVjNGTllDUTdpSDV0bVBaU3U3dzlOQ21vc0c3bWN1MnY4SlBSCjRaTUNWV29lbzRpWmlG
dG1hZFRrN1pBSENUS0ovUU1hbDRERXFCc0EvTE42K3ZnMjFzVnBjczgvN05hTjZXdncKN3A4b09l
NTQzRUs4R1lFZ1Noa1VGVkZ4UFE5Ry8zK2FUY2pKeklOMEFLWGZ6U2FKd3JUY1FVaWo5aFFWUkdn
cApOcURQMCtxOTBGdlVLWGh0dmZEMDZQU3BzbVgvZXpYNmc4K1RhZDJ6Rmp3d2s3MnRzRHIzL0I5
Ujc1dmtLNXFkCnM2TzdJWUpyQ0pqelU2ZVpkNW81UHkzT0VORk9xS0VTNkloeUZnSUJpZXEvYmY1
QzgzblhaWlEwNGNxN0hnakQKYmtaMWI2aUpGaDJjTkNSd1QyRmd4YnpHc3NGbEcyTHMxVGM9Ci0t
LS0tRU5EIENFUlRJRklDQVRFIFJFUVVFU1QtLS0tLQo=

controlplane ~ ➜  ls -0l
ls: invalid option -- '0'
Try 'ls --help' for more information.

controlplane ~ ✖ ls- l
-bash: ls-: command not found

controlplane ~ ✖ ls -l
total 8
-rw-r--r-- 1 root root  887 Mar 29 19:15 akshay.csr
-rw------- 1 root root 1704 Mar 29 19:15 akshay.key

controlplane ~ ➜  vim csr.yml

controlplane ~ ➜  cat csr.yml 
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
  name: akshay
spec:
  request:
    LS0tLS1CRUdJTiBDRVJUSUZJQ0FURSBSRVFVRVNULS0tLS0KTUlJQ1ZqQ0NBVDRDQVFBd0VURVBN
    QTBHQTFVRUF3d0dZV3R6YUdGNU1JSUJJakFOQmdrcWhraUc5dzBCQVFFRgpBQU9DQVE4QU1JSUJD
    Z0tDQVFFQWp0aERDak1JUVNUTi83UUdzQ1ZkalFBSzZMdnZmU21LYzVrNjRQamU2QXhUCm1MT0l1
    SzFNeElhUGx0WmVrMGIvKzhHU0JQMGtoL2VVQUNaUTVwRDRKRTZ1ODVNYWxtc0tNYUxsQTRRcGdN
    NXEKY2d6NzJOUjR0QU0yTkpGYXpkbjYrTjFWM1Fudk9KK3U3d2lQYmg4c21QdVVWOTlBVHY5K1JB
    RzJJMFlXckpsQQpYSlh3UGRJalJvcTBIYlZTVlhhUUJkVVpmbnRubUNPd2VlNm9CLzdwVjdic2lu
    d1FrajU0ZXg2enAxenRxNkQ5CkJBZ3Y4QmlSdmNzNmQrNEhtMFJXVXJuVms3V1FDM2c1OFRkNHNw
    a2IzMlJoTjdRVW5aOTVzMlFYUnVoaThETjYKSW9kVFVZaFRaNnN4OWovWGVMODlIeXBpb212UUh0
    TXB4RnJPbFloOGlRSURBUUFCb0FBd0RRWUpLb1pJaHZjTgpBUUVMQlFBRGdnRUJBRm5wRnB1Z3NH
    cFlZQTYvVjNGTllDUTdpSDV0bVBaU3U3dzlOQ21vc0c3bWN1MnY4SlBSCjRaTUNWV29lbzRpWmlG
    dG1hZFRrN1pBSENUS0ovUU1hbDRERXFCc0EvTE42K3ZnMjFzVnBjczgvN05hTjZXdncKN3A4b09l
    NTQzRUs4R1lFZ1Noa1VGVkZ4UFE5Ry8zK2FUY2pKeklOMEFLWGZ6U2FKd3JUY1FVaWo5aFFWUkdn
    cApOcURQMCtxOTBGdlVLWGh0dmZEMDZQU3BzbVgvZXpYNmc4K1RhZDJ6Rmp3d2s3MnRzRHIzL0I5
    Ujc1dmtLNXFkCnM2TzdJWUpyQ0pqelU2ZVpkNW81UHkzT0VORk9xS0VTNkloeUZnSUJpZXEvYmY1
    QzgzblhaWlEwNGNxN0hnakQKYmtaMWI2aUpGaDJjTkNSd1QyRmd4YnpHc3NGbEcyTHMxVGM9Ci0t
    LS0tRU5EIENFUlRJRklDQVRFIFJFUVVFU1QtLS0tLQo=
  usages:
  - digital signature
  - key encipherment
  - server auth
  - client auth
  signerName: kubernetes.io/kube-apiserver-client


controlplane ~ ➜  k apply -f csr.yml 
Error from server (BadRequest): error when creating "csr.yml": CertificateSigningRequest in version "v1" cannot be handled as a CertificateSigningRequest: illegal base64 data at input byte 76

controlplane ~ ✖ ^C

controlplane ~ ✖ ^C

controlplane ~ ✖ ^C

controlplane ~ ✖ vi csr.yml 

controlplane ~ ➜  k apply -f csr.yml 
certificatesigningrequest.certificates.k8s.io/akshay created


controlplane ~ ➜  k get csr
NAME          AGE     SIGNERNAME                                    REQUESTOR                  REQUESTEDDURATION   CONDITION
agent-smith   14s     kubernetes.io/kube-apiserver-client           agent-x                    <none>              Pending
akshay        2m12s   kubernetes.io/kube-apiserver-client           kubernetes-admin           <none>              Approved,Failed
csr-df8th     13m     kubernetes.io/kube-apiserver-client-kubelet   system:node:controlplane   <none>              Approved,Issued

controlplane ~ ➜  k get csr agent-smith -o yaml
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
  creationTimestamp: "2025-03-29T19:24:27Z"
  name: agent-smith
  resourceVersion: "1509"
  uid: 4204d0ba-aea6-484a-9e01-dc0721dc8b45
spec:
  extra:
    authentication.kubernetes.io/credential-id:
    - X509SHA256=e20e0483f114618220370a6fb68ef56b46354e25aaa557e64a7ed593880c8123
  groups:
  - system:masters
  - system:authenticated
  request: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURSBSRVFVRVNULS0tLS0KTUlJQ1dEQ0NBVUFDQVFBd0V6RVJNQThHQTFVRUF3d0libVYzTFhWelpYSXdnZ0VpTUEwR0NTcUdTSWIzRFFFQgpBUVVBQTRJQkR3QXdnZ0VLQW9JQkFRRE8wV0pXK0RYc0FKU0lyanBObzV2UklCcGxuemcrNnhjOStVVndrS2kwCkxmQzI3dCsxZUVuT041TXVxOTlOZXZtTUVPbnJEVU8vdGh5VnFQMncyWE5JRFJYall5RjQwRmJtRCs1eld5Q0sKeTNCaWhoQjkzTUo3T3FsM1VUdlo4VEVMcXlhRGtuUmwvanYvU3hnWGtvazBBQlVUcFdNeDRCcFNpS2IwVSt0RQpJRjVueEF0dE1Wa0RQUTdOYmVaUkc0M2IrUVdsVkdSL3o2RFdPZkpuYmZlek90YUF5ZEdMVFpGQy93VHB6NTJrCkVjQ1hBd3FDaGpCTGt6MkJIUFI0Sjg5RDZYYjhrMzlwdTZqcHluZ1Y2dVAwdEliT3pwcU52MFkwcWRFWnB3bXcKajJxRUwraFpFV2trRno4MGxOTnR5VDVMeE1xRU5EQ25JZ3dDNEdaaVJHYnJBZ01CQUFHZ0FEQU5CZ2txaGtpRwo5dzBCQVFzRkFBT0NBUUVBUzlpUzZDMXV4VHVmNUJCWVNVN1FGUUhVemFsTnhBZFlzYU9SUlFOd0had0hxR2k0CmhPSzRhMnp5TnlpNDRPT2lqeWFENnRVVzhEU3hrcjhCTEs4S2czc3JSRXRKcWw1ckxaeTlMUlZyc0pnaEQ0Z1kKUDlOTCthRFJTeFJPVlNxQmFCMm5XZVlwTTVjSjVURjUzbGVzTlNOTUxRMisrUk1uakRRSjdqdVBFaWM4L2RoawpXcjJFVU02VWF3enlrcmRISW13VHYybWxNWTBSK0ROdFYxWWllKzBIOS9ZRWx0K0ZTR2poNUw1WVV2STFEcWl5CjRsM0UveTNxTDcxV2ZBY3VIM09zVnBVVW5RSVNNZFFzMHFXQ3NiRTU2Q0M1RGhQR1pJcFVibktVcEF3a2ErOEUKdndRMDdqRytocGtueG11RkFlWHhnVXdvZEFMYUo3anUvVERJY3c9PQotLS0tLUVORCBDRVJUSUZJQ0FURSBSRVFVRVNULS0tLS0K
  signerName: kubernetes.io/kube-apiserver-client
  usages:
  - digital signature
  - key encipherment
  - server auth
  username: agent-x
status: {}

controlplane ~ ➜  k certificate reject agent-smith
Modify certificate resources.

Available Commands:
  approve       Approve a certificate signing request
  deny          Deny a certificate signing request

Usage:
  kubectl certificate SUBCOMMAND [options]

Use "kubectl certificate <command> --help" for more information about a given
command.
Use "kubectl options" for a list of global command-line options (applies to all
commands).

controlplane ~ ➜  k certificate deny agent-smith
certificatesigningrequest.certificates.k8s.io/agent-smith denied

controlplane ~ ➜  k get csr
NAME          AGE     SIGNERNAME                                    REQUESTOR                  REQUESTEDDURATION   CONDITION
agent-smith   104s    kubernetes.io/kube-apiserver-client           agent-x                    <none>              Denied
akshay        3m42s   kubernetes.io/kube-apiserver-client           kubernetes-admin           <none>              Approved,Failed
csr-df8th     15m     kubernetes.io/kube-apiserver-client-kubelet   system:node:controlplane   <none>              Approved,Issued

controlplane ~ ➜  k delete csr ^C

controlplane ~ ✖ k get csr
NAME          AGE     SIGNERNAME                                    REQUESTOR                  REQUESTEDDURATION   CONDITION
agent-smith   2m10s   kubernetes.io/kube-apiserver-client           agent-x                    <none>              Denied
akshay        4m8s    kubernetes.io/kube-apiserver-client           kubernetes-admin           <none>              Approved,Failed
csr-df8th     15m     kubernetes.io/kube-apiserver-client-kubelet   system:node:controlplane   <none>              Approved,Issued

controlplane ~ ➜  k delete csr agent-smith
certificatesigningrequest.certificates.k8s.io "agent-smith" deleted

controlplane ~ ➜  \