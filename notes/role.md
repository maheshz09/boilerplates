apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: app-role-cka
  namespace: default
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get"]

---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: app-rolebinding-cka
  namespace: default
subjects:
  - kind: ServiceAccount
    name: app-account
    namespace: default
    apiGroup: rbac.authorization.k8s.io
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: app-role-cka





----- NEW


apiVersion: v1
kind: ServiceAccount
metadata:
  name: app-account
  namespace: default
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: app-role-cka
  namespace: default
rules:
- apiGroups: [""]
  resources: ["pods"]
  verbs: ["get"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: app-role-binding-cka
  namespace: default
subjects:
- kind: ServiceAccount
  name: app-account
  namespace: default
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: app-role-cka




---
apiVersion: v1
kind: ConfigMap
metadata:
  name:   
data: 
  APPLICATION=web-app

---


# Patch or edit your deployment like this
apiVersion: apps/v1
kind: Deployment
metadata:
  name: webapp-deployment
spec:
  replicas: 1
  selector:
    matchLabels:
      app: webapp
  template:
    metadata:
      labels:
        app: webapp
    spec:
      containers:
      - name: webapp-container
        image: your-webapp-image
        envFrom:
        - configMapRef:
            name: webapp-deployment-config-map
