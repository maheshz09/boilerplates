Storage: container storage interfaces (CSI)

(Container Runtime interface)
	So early days kubernetes only relay on docker as runtime engine to host application. as kubernetes 
	sccale and widely used they have intruduced something called as CSI driver that makes other provides like rocket, cri-o and many other docker runtime integration with kubernetes easily.

(Container Network Interface)
	Same like CSI, also with the CNI standards any company can integrate there network solution on the kubernetes like 
	flannel, calico, cilium. to extend network functionality of kubernetes.
	
(Container Storage Interface)
	With CSI we can write our own drivers to work with our storage systems, use that storage in kubernetes pods, deployments etc.
	like, portworx, Amazon EBS, DEll EMC, NETApp, GlusterFS like storages can also be implemnted by using this.
	

NOTE: CSI not an kubernetes specific standurd 

	it is ment to be a universal standard, and if implemnted it allows any container orchestration tool to work with any storage  vendor with a supported plugins 


so some of them have already onborded the CSI. 

here is how CSI looks like. it defines set of RPC's (Remote procedure calls) that will be called by the container orchestratior and this shuould be implemnted by storage drivers.

like (RPC)
	1. SHOULD call to provision a new volume
	2. SHOULD call to Delete a volume
	3. SHOULD call to place a workloads that uses a volume onto a node
	4. SHOULD provision a volume on storage 
	5. SHOULD decomission a volume
	6. SHOULD make volume available to nodes.
	
RPC's should handle this kind's of requests & return the results of operations. similerly orchestrator should call delete volume RPCS when a volume is use to be deleted. 