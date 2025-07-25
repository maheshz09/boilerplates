# node taints and tolerations in k8s

### information
Node affinity is a property of Pods that attracts them to a set of nodes (either as a preference or a hard requirement). ``Taints are the opposite -- they allow a node to repel a set of pods``.

``Tolerations are applied to pods``. Tolerations allow the scheduler to schedule pods with matching taints. Tolerations allow scheduling but don't guarantee scheduling: the scheduler also evaluates other parameters as part of its function.

Taints and tolerations work together to ensure that pods are not scheduled onto inappropriate nodes. One or more taints are applied to a node; this marks that the node should not accept any pods that do not tolerate the taints.

### taints
``Taints applied to nodes``
``tolerations applied to pods``

### how to apply taints to **nodes**
```bash
     kubectl taint nodes <node-name> key=value:taint-effect
```
if you want to specify node to application blue then the key value pair would be **app=blue**
and the ``taint effect`` is use to what happens to PODs that DO NOT Tolerate this taint ? 
there are three taint effects 
    1. ``No scedule``: witch means pods not scedule on nodes
    2. ``PreferNoScedule``: witch means system tries to avoid pod sceduling on that node (no garenteed)
    3. ``NoExecute``: witch means new pods not able to scedule on nodes & exesting pod on the machine if any will be evicted if they do not tolerate that taint.

These pods may have been scheduled
on the node before the taint was applied to the node.

An example command would be to taint node one with the key value pair app=blue and an effect of no schedule.

```bash
    kubectl taint nodes <node1> app=blue:NoSchedule
```

tolerations are added to pods, to add a tolerations to pod first pull up pod defination file. and the specs sections of the pod defination file, add a section called, 'tolerations',
move the same values used while creating the taint.
Under this section, the key is app operator is equal
value is blue, and the effect is no schedule.
And remember, all of these values 

need to be encoded in double codes. When the pods are now created or updated with the new tolerations, they are either not scheduled on nodes or evicted from the existing nodes depending on the effect set.

pod-defination.yaml
```yaml
    apiVersion: v1
    kind: Pod
    metadata:
        name: nginx
        labels:
          env: test
    spec:
      containers:
      - name: nginx
        image: nginx
        imagePullPolicy: IfNotPresent
      tolerations:
      - key: "app"
        operator: "Equal/Exists"
        effect: "NoSchedule"

```

### how it manages the exesting nodes present on the machine
Let us try to understand the no execute taint effect in a bit more depth. In this example, we have three nodes running some workload. We do not have any taints or tolerations at this point, so they're scheduled this way. We then decided to dedicate node one for a special application, and as such, we taint the node with the application name
and add a toleration to the pod that belongs to the application, which happens to be pod D in this case.
While tainting the node, we set the taint effect to no execute, and as such, once the taint on the node takes effect, it evicts pod C from the node, which simply means that the pod is killed. The pod D continues to run on the node as it has a toleration to the taint blue.

Now, going back to our original scenario where we have taints and tolerations configured. Remember taints and tolerations are only meant to restrict nodes from accepting certain pods.
In this case, node one can only accept pod D, but it does not guarantee
that pod D will always be placed on node one.

Since there are no taints or restrictions applied on the other two nodes, pod D may very well be placed
on any of the other two nodes. So remember, taints and tolerations does not tell the pod to go to a particular node.

Instead, it tells the node to only accept pods with certain tolerations. If your requirement is to restrict a pod to certain nodes, it is achieved through another concept called as node affinity, which we will discuss in the next lecture.

### taint on master node 
Finally, while we are on this topic, let us also take a look at an interesting fact.
So far we have only been referring to the worker nodes. But we also have master nodes in the cluster,
which is technically just another node
that has all the capabilities of hosting a pod,plus it runs all the management software.Now, I'm not sure if you noticed the scheduler does notschedule any pods on the master node.Why is that?When the Kubernetes cluster is first set up,a taint is set on the master node automaticallythat prevents any pods from being scheduled on this node.You can see this, as well as modify this behavior,if required.However, a best practice is to not deploy application workloads on a master server.To see this taint, run a kubectl describe node command with cube master as the node name and look for the taint section.You will see a taint set to not schedule any pods on the master node.Well, that's it for this lecture.

```bash
    kubectl describe node kubemaster | grep taint
```

### example

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: bee
spec:
  containers:
  - name: nginx
    image: nginx
  tolerations:
  - key: "spray"
    operator: "Equal"
    value: "mortein"
    effect: "NoSchedule"
```

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: bee
spec:
  containers:
  - name: nginx
    image: nginx
  tolerations:
  - key: "mortein"
    operator: "Exists"
    effect: "NoSchedule"
```


### how to taint from control-plane/any node
I am going to remove taint from controlplane to scedule pod's on controlplane

```bash
    kubectl describe node/controlplane | grep Taint #to get exesting taint information
    # after you run describe command you will output something like this 
    # Taints: master-node-01 spray=mortein:NoSchedule

    # to remove that exesting taint
    kubectl taint nodes master-node-01 spray=mortein:NoSchedule-
```
