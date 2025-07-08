Storage: Persistant volume claims.

	claim the extesting storge available on the node.
	
Persistantvolumes and Persistantvolumesclaims are two separate objects in kubernetes. later we can able to bind this together to use.

based on request and propertes set on the volume kubernetes tries to mount volume.

every Persistant volume claim is mounted to Persistantvolume, during binding process kubernetes tries to find a Persistantvolume that has sufficient capacity as provided by user.

and any other request properties such as Sufficient capacity, Access Modes, Volume mounts,  storage class. (on this aswell we can use labels and selectors to assign specific volume.)

NOTE: a smaller claim bound to a larger volume if all the other critaria matches and there are no better options. there is a one to one relation to claims and volumes. so no other claims can utilize remaining capacity in the volume.


if there is no volume available the Persistantvolumesclaims remain in the pending state until new volume made available to cluster once newer volumes are available the claims automatically bound to newly available volume.

pv-def.yml

apiVersion: v1
kind: PersistantvolumeClaim
matadata:
	name: my-claim
spec:
	accessModes:
		- ReadWriteOnce
	resources: 
		requests: 
			Storage: 500Mi
			
# kubectl create -f pv-def.yml

to view extesting claims in the cluster
# kubectl get PersistantvolumeClaim

we can see the claim into pending state.

	when the claim is created kubernetes will look at the volumes created previously. the access modes are match, the capacity match, since there is no other volumes the PersistantvolumeClaim will be mounted to that volume.
	
pvc-def.yml

apiVersion: v1
kind: Persistantvolume
metadata:
	name: pv-011
spec:
	accessModes:
		- ReadWriteOnce
	resources: 
		requests: 
			Storage: 1Gi
	awsElasticBlockStore:
		volumeID: <volumeID>
		fsType: ext4

now the Persistant volume claim is mounted to Persistant volume.

# kubectl get PersistantvolumeClaim


to delete pvc 
# kubectl delete PersistantvolumeClaim my-claim
	so suppose if we deleted the pvc the pv remain, we can choose this behaviour what is to happen to that pv. by default it set to retain,
	1. PersistantvolumeReclaimPolicy: retain
		the Persistant volume is remain until its deleted manually. it is not be resusable by any other claims, or it can be deleted automatically by the policy.
	2. PersistantvolumeReclaimPolicy: deleted
		
	also there is 3rd option.
	
	3. PersistantvolumeReclaimPolicy: Recycle
			with this the data will be automatically scarb before it provided to oher PersistantvolumeClaim
			
