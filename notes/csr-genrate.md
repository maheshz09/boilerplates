#to genrate csr 
openssl req -new -out dev.s2-eu.capgemini.com.cnf.csr -key dev.s2-eu.capgemini.com.cnf.key -config dev.s2-eu.capgemini.com.cnf


#to validate csr MD5 value
openssl req -noout -modulus -in coconet-svn-fs-01.fs.capgemini.com.csr | openssl md5

#to check key MD5 value
openssl rsa -noout -modulus -in coconet-svn-fs-01.fs.capgemini.com.key | openssl md5

#cert jkey
openssl x509 -noout -modulus -in scm-coconet2.capgemini.com.cer | openssl md5


# NOTE

1. to verify we can validate the md5 value matching with key and the csr that genrated
2. key don't change its the same
3. to verify the csr details and everything is correct 
openssl req -in mycsr.csr -noout -text




For dc1 & pod3
 
----- internal 
FRPR3JPDLB801PR
FRPR3JPDLB802PR 
FRPR3DC1RVPR1
FRPR3DC1RVPR2

-- ext
FRPR3PODSRDMZ
FRPR3JPDRVPR01
FRPR3JPDRVPR02


1.haproxycert--->> 1st-key 2nd-pem(cert) 3rd-(interm-root.pem)

cat s2-eu.capgemini.com.key > s2-eu.capgemini.com.haproxycert
cat s2-eu.capgemini.com.pem >> s2-eu.capgemini.com.haproxycert
cat s2-eu.capgemini.com-intermroot.pem >> s2-eu.capgemini.com.haproxycert

2.withca.pem---->>1st-rootintermimtent 2nd .pem

cat pl.s2-eu.capgemini.com-rootinterm.pem > pl.s2-eu.capgemini.com-withca.pem
cat pl.s2-eu.capgemini.com.pem >> pl.s2-eu.capgemini.com-withca.pem

unlink currentcertificates

ln -s e-3d-jira.capgemini.com.haproxycert s2-eu.capgemini.com.haproxycert
unlink <current folder>
ln -s <new folder> currentcertificates


----------------------------------------------------------------
self singing cert (internally used by k8s)

# to genrate private keys:
openssl genrsa -out ca.key 2048

#openssl request command along with key to genrate a certificate singing req.
openssl req -new ca.key -subj "/CN=KUBERNETES-CA" -out ca.csr

#sign certificate 
openssl x509 -req ca.csr -signkey ca.key -out ca.crt 
## self signed with own private key 


----------------------------------------------------------------------

# genrating client certificate like (admin user) like(kube API)

# to genrate private keys:
openssl genrsa -out admin.key 2048

#openssl request command along with key to genrate a certificate singing req.
openssl req -new admin.key -subj "/CN=kube-admin/o=system:masters" -out admin.csr

#sign certificate 
openssl x509 -req admin.csr -signkey admin.key -out admin.crt 

that admin.crt is self signed certificate that we will use to authonticate with kubernetes cluster
this process is very much similer to like account creatation process, the admin user has special privilages aswell 
that group need to added in the certificate, we can that group detail in the req commandof
openssl -- O means ognization
like this /o=system:masters

so now the certificate that we will get it has admin privilages.

we will do the same process to genrate client certificatates for all other components that access kube API server.



-- usage:

tru admin certificate:
	we can use cert insted of username and password in the rest API call

curl https://kube-apiserver:6443/api/v1/pod \
	--key admin.key --cert admin.crt \
	--cacert ca.crt

-----------------------------------------------------------------------------

server side certificatates

ETCD server:
	



---------------------------------------------------------------------------------

# view and verify exesting cluster certificatates issues.
openssl x509 -in /etc/kubernetes/pki/apiserver.crt -text -noout

eg.
openssl x509 -in /etc/kubernetes/pki/etcd/ca.crt -text -noout

