#helm
helm template argocd argo/argo-cd --version 7.7.3 --set crds.install=false  > /tmp/argocdTemplate

helm install argocd argo/argo-cd -f /tmp/argocdTemplate
helm install argocd argo/argo-cd -f ~/argo-helm.yaml --skip-schema-validation --force --namespace argocd

k get pods 
NAME                                                READY   STATUS    RESTARTS   AGE                                                                                                                                                                                                            130 ↵
argocd-application-controller-0                     1/1     Running   0          101s
argocd-applicationset-controller-55b9696ff6-mdgkt   1/1     Running   0          99s
argocd-dex-server-5b7b96ddf9-ddvtq                  1/1     Running   0          102s
argocd-notifications-controller-85774c5568-9jcq7    1/1     Running   0          102s
argocd-redis-7f855978d6-8mh2k                       1/1     Running   0          102s
argocd-repo-server-7c8867c847-t72rn                 1/1     Running   0          102s
argocd-server-775846f8bc-zzwdf                      1/1     Running   0          102s

----
"helm template argocd argo/argo-cd --version 7.7.3 --set crds.install=false -n argocd >argo-helm.yaml" 
  "helm install argocd argo/argo-cd --version 7.7.3 --set crds.install=false -n argocd" .

----
another way...

controlplane:~$ helm pull argo/argo-cd --version 7.7.3
controlplane:~$ ls
argo-cd-7.7.3.tgz  filesystem
controlplane:~$ tar -zxf argo-cd-7.7.3.tgz

controlplane:~$ k create ns argocd

namespace/argocd created

helm install argocd ./argo-cd -n argocd --set crds.install=false


