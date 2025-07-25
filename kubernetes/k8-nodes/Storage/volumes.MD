Storage: volumes.


	So on Docker terminiology docker is efemral in nature that means they can leave short amount time. 
once the process complete they can be deleted automatically. to persist data inside from the volume we can use this terminiology 

same as like into kubernetes.
	pods are transiant in nauture when pod is created we can assing a volume to a pod to store data, and our data remain secure. if the pod also went down we have our data safe in a volume.
	

eg. 

we have an application that genrates a multiple numbers, while defining pod specfication we can add volume and some datails as well as mentioned below.

volumes:
	- name: data-volume
	  hostPath:
		path: /data
		type: Directory

so this way any data genrated by my pod kept into that data dir in host system.

and to mount the pod inside the conatiner we use volumeMounts like.

apiVersion: v1
kind: pod 
metadata:
	name: random-pod
spec:
	containers:
		- image: alpine
		  name: alpine
		  command: ["/bin/sh", "c"]
		  args: ["shuf -i 0-100 -n 1 >> /opt/numbers.out;"]
		  volumeMounts:
			- mountPath: /opt/
			  name: data-volume
			  
			  
so that mount point will be mounted inside my conatiner, in /opt. 

	
	so thats works fine in single node. what if we have multi-node cluster, in that case the data will mess into all the nodes in the cluster and expects that it is on single node.
	
insted, you can use external storage solutions like. NFS, GlusterFS, Amazonwebservices, EBS etc.

suppose if we have to replace this opentioned insted we can use EBS volme like 


volume:
- name: data-volume
  awsElasticBlockStore:
	VolumeID: <volme-id>
	fsType: ext4
	
	
so volume are stored into ebs volumes in this way.