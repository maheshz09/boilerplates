# create administrator in kubernetes (TRU API) 

# private key
openssl genrsa -out mahesh.key 2048

#cert sign request(CSR)
openssl req -new -key mahesh.key -subj "/CN=mahesh" -out mahesh.key

# certificate sign request object, it is created like any other object in kubernetes using a manifest file

```yaml

apiVersion: certificates.k8s.io/v1
kind: certificateSigningRequest
metadata:
	name: mahesh
spec:
	expirationSeconds: 600 #seconds
	usage: 
	- digital signature
	- key encipherment
	- server auth
	request:
	
```

request is the feild whare we specify CSR request sent by the user. we dont specify in plain text we must 
encode in base64

cat mahesh.csr | base64

move encoded feild to request section of object, and then submit the request.


# once the object is created, it can seen by the administrator
kubectl get csr

#to approve request
kubectl certificate approve mahesh

it will sign request using private and certificate.

# the cert is created now it extracted and shared with user.
kubectl get csr mahesh -o yaml

the genrated cert is part of output witch is base64 encoded



behind the secne controller-manager from kubernetes handing all those cert related tasks

