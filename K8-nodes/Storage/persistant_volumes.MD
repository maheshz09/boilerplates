Storage: persistant volume

So in previos sections we have configured pod & volume defination into same file only. every configuration information requires by volume goes inside pod defination yaml file.

when we have large cluster each and every user deploying there pods, we need to configure them manually if we go with this method. what ever the storage solution that we are using user needs to configuraer them before deploying pods. when changes has to made user needs to update to all the pods.

insted we like to manage storage centrally. we like to configure in such way that administratr would create a storage pool, and the user take storage from that pool as required

(persistant volumes)
	 a persistant volume is a cluster wide pool of storage volumes, configured by an administratr to be used by users deploying appications on the cluster. so the users can select the storage from this pool using PVC's(persistant volume claims) 
	 
	
pv-def.yml

apiVersion: v1
kind: PersistantVolume
metadata:
	name: pv-vo11
	
spec: 
	accessModes:
		- ReadWriteOnce (it defines how the resuces should be utilized. supported valyes. ReadOnlyMany, ReadWriteOnce, ReadWriteMany)
	capacity:
		Storage: 1Gi (reserverd)
	hostPath: (it uses storage from local nodes dir / not for PR)
		path: /data
	
# k apply -f pv-def.yml

to list pv's

# k get persistentvolume

NOTE: we can also using something like ebs volume and replace the "hostPath" section.
