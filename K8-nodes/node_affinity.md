# Node affinity to host pods on perticular nodes

### information
In this lecture, we will talk about node affinity feature in Kubernetes. The primary purpose of node affinity feature is to ensure that pods are hosted on particular nodes, in this case to ensure the large data processing pod ends up on node one.

In the previous lecture, we did this easily using node selectors. We discussed that you cannot provide advanced expressions like or or not with node selectors. The node affinity feature provides us with advanced capabilities to limit pod placement on specific nodes. With great power comes great complexity, so the simple node selector specification

### with node-selector
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

### with node-affinity
```yaml
    apiVersion: v1
    kind: Pod
    metadata:
      name: nginx
    spec:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: disktype
                operator: In
                values:
                - ssd            
      containers:
      - name: nginx
        image: nginx
        imagePullPolicy: IfNotPresent
```

will now look like this with node affinity although both does exactly the same thing.
Place the pod on the large node. Let us look at it a bit closer. Under spec you have affinity   and then node affinity under that. And then you have a property that looks like a sentence called,   required during scheduling, ignored during execution.  No description needed for that.   And then you have the node selector terms, that is NRA, and that is where you will specify the key and value pairs.  The key value pairs are in the form   key, operator and value,   where the operator is in.  The in operator ensures that the pod  will be placed on a node  whose label size has any valu in the list of values specified here?   In this case, it is just one called large. If you think your pod could be placed on a large or a medium node, you could simply add the value to the list of values like this.

```yaml
    apiVersion: v1
    kind: Pod
    metadata:
      name: nginx
    spec:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: disktype
                operator: In
                values:
                - Large
                - Medium            
      containers:
      - name: nginx
        image: nginx
        imagePullPolicy: IfNotPresent
```
You could use the node in operator to say something like, size not in small, where node affinity will match the nodes with a size not set to small.

```yaml
    apiVersion: v1
    kind: Pod
    metadata:
      name: nginx
    spec:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: disktype
                operator: NotIn
                values:
                - Small            
      containers:
      - name: nginx
        image: nginx
        imagePullPolicy: IfNotPresent
```

We know that we have only set the label size to large and medium nodes. The smaller nodes don't even have the label set so we don't really have to even check the value of the label.
As long as we are sure we don't set a label size to the smaller node, using the exist operator will give us the same result.

```yaml
    apiVersion: v1
    kind: Pod
    metadata:
      name: nginx
    spec:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: disktype
                operator: Exists            
      
      containers:
      - name: nginx
        image: nginx
        imagePullPolicy: IfNotPresent
```
 
The exist operator will simply check if the label size exists on the nodes, and you don't need the value section for that as it does not compare the values. 
There are a number of other operators as well. 
Check the documentation for specific details. Now, we understand all of this and we're comfortable with creating a pod with specific affinity rules. When the pods are created, these rules are considered and the pods are placed onto the right nodes. But ``what if node affinity could not match a node with a given expression?`` In this case, ``what if there are no nodes with the label called size?`` Say we had the labels and the pods are scheduled.
 
``What if someone changes the label on the node at a future point in time?``
``will the pod continue to stay on the node?``
 
which happens to be the type of node affinity. The type of node affinity defines the behavior of the scheduler with respect to node affinity and the stages in the life cycle of the pod.

``There are currently **two types of node affinity** available,``

Available at kubernetes:  
  1. ``required during scheduling, ignored during execution``
  2. ``and preferred during scheduling, ignored during execution.``

And there are additional types of node affinity planned

planned:
  1. ``Required during scheduling, required during execution.``


We will now break this down to understand further. We will start by looking at the two available affinity types. There are two states in the lifecycle of a pod when considering node affinity, during scheduling and during execution. During scheduling is the state where a pod does not exist and is created for the first time. We have no doubt that when a pod is first created the affinity rules specified are considered to place the pods on the right nodes.

### Question:
  1. ``Now, what if the nodes with matching labels are not available?``
 
For example, we forgot to label the node as large. That is where the type of node affinity used comes into play. If you select the required type, which is the first one, the scheduler will mandate that the pod be placed on a node with a given affinity rules. If it cannot find one, the pod will not be scheduled. This type will be used in cases where the placement of the pod is crucial. If a matching node does not exist the pod will not be scheduled. But let's say the pod placement is less important than running the workload itself. In that case, you could set it to preferred and in cases where a matching node is not found, the scheduler will simply ignore node affinity rules and place the pod on any available node. This is a way of telling the scheduler, ````Hey, try your best to place the pod on matching node." But if you really cannot find one, just place it anywhere.`` The second part of the property or the other state is during execution. During execution is the state where a pod has been running and a change is made in the environment that affects node affinity, such as a change in the label of a node.


For example, say an administrator removed the label
we set earlier called size equals large from the node.

### required during execution if some label changes happens to nodes / available node affinity types 
that are running on the node? As you can see, the two types of node affinity available today has this value set to ignored, which means pods will continue to run and any changes in node affinity will not impact them once they are scheduled.


### planned node affinity type [Required during scheduling, required during execution] 
The new types expected in the future only have a difference in the during execution phase. A new option called required during execution is introduced which will evict any pods that are running on nodes that do not meet affinity rules. In the earlier example, a pod running on the large node will be evicted or terminated if the label large is removed from the node.