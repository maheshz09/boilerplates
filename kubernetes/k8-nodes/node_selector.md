# Node selectors  

### information 
we will talk about node selectors in Kubernetes. Let us start with a simple example. You have a three node cluster of which two are smaller nodes with lower hardware resources, and one of them is a larger node configured with higher resources. You have different kinds of workloads running in your cluster. ``You would like to dedicate the data processing workloads that require higher horsepower to the larger node as that is the only node that will not run out of resources in case the job demands extra resources.`` However, in the current default setup, any pods can go to any nodes. So, pod C in this case, may very well end up on nodes two or three, which is not desired.

### node selectors
To solve this,we can ``set a limitation on the pods`` so that they only run on particular nodes.There are two ways to do this.The first is using node selectorswhich is the simple and easier method.For this,we look at the pod definition file we created earlier.This file has a simple definition to create a podwith a data processing image.To limit this pod to run on the larger node,we add a new property called ``node selector to the spec section and specify the size as large.``

```yaml
apiVersion: v1
    kind: Pod
    metadata:
        name: myapp-prod
    spec:
      containers:
      - name: data-processor
        image: data-processor
    nodeSelector:
        size: Large
```
But wait a minute, ``where did the size large come from`` ``and how does Kubernetes know which is the large node?`` The key value pair of size and large are in fact labels assigned to the nodes. The scheduler uses these labels to match and identify the right node to place the pods on.

### label to node
Labels and selectors are a topic we have seen many timesthroughout this Kubernetes course such as with services, replica sets, and deployments.To use labels in a node selector like this,you must have first labeled your nodesprior to creating this pod.So, let us go back and see how we can label the nodes.
 
To label a node, use the command cube control label nodes followed by the name of the node and the label in a key value pair format.

```bash
    kubectl label nodes <node-name> <label-key>=<label-key>
```

In this case, it would be cube control label nodes node one followed by the label in a key value format such as size=large.

```bash
    kubectl label nodes <node-1> size=Large
```

### now place nodes to pod

Now that we have labeled the node, we can get back to creating the pod. This time with the node selector set to a size of large. When the pod is now created, it is placed on node one as desired.

```yaml
apiVersion: v1
    kind: Pod
    metadata:
        name: myapp-prod
    spec:
      containers:
      - name: data-processor
        image: data-processor
    nodeSelector:
        size: Large
```

```bash
    kubectl apply -f pod-defination.yaml
```

Node selectors served our purpose, but it has limitations. We used a single label and selector to achieve our goal here. But what if our requirement is much more complex? For example, we would like to say something like place the pod on a large or medium node, or something like place the pod on any nodes that are not small. You cannot achieve this using node selectors. For this, ``node affinity`` and ``anti-affinity`` features
