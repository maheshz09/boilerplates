curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
helm repo update
helm install vault hashicorp/vault

k exec -it pods/vault-0 -- /bin/sh

/ $ vault operator init
Unseal Key 1: ao3sk3sVIHtPYwM1zOqMNbXooAiaEp5B7vOMRI4FUPzG
Unseal Key 2: 3isvu7wLSpVz5lwO14RyVtx5qLswkzLDNjjV7LeUC+YB
Unseal Key 3: wknWrORZBxn52Zb4YipY1TBg5eNH7dN8pHGEcFwF1riP
Unseal Key 4: e+gXavxDoHcC6RYTSfi4NpZRi3TB/Hu70iKFq+PDHLbT
Unseal Key 5: Zt11G+3PNSSjChkmLTNg+NLRMfyjaPRFy4Ghtw8tFVkM

Initial Root Token: hvs.6XOciSyaDPldg6pP6f3lem3A

Vault initialized with 5 key shares and a key threshold of 3. Please securely
distribute the key shares printed above. When the Vault is re-sealed,
restarted, or stopped, you must supply at least 3 of these keys to unseal it
before it can start servicing requests.

Vault does not store the generated root key. Without at least 3 keys to
reconstruct the root key, Vault will remain permanently sealed!

It is possible to generate new unseal keys, provided you have a quorum of
existing unseal keys shares. See "vault operator rekey" for more information.
/ $


vault status
vault operator unseal

vault login

vault secrets list

vault secrets enable kv-v2

vault kv put kv-v2/vault-demo/mysecret username=mahesh password=passwd

vault kv get kv-v2/vault-demo/mysecret

vault policy list


/ $ vault policy write mysecret - << EOF
> path "kv-v2/data/vault-demo/mysecret" {
> capabilities = ["read"]
> }
> EOF

vault auth enable kubernetes

vault auth list


vault write auth/kubernetes/config \
  kubernetes_host=https://$KUBERNETES_SERVICE_HOST:$KUBERNETES_SERVICE_PORT
  ORRRRRRRRR
vault write auth/kubernetes/config \
  token_reviewer_jwt="$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \
  kubernetes_host="https://$KUBERNETES_SERVICE_HOST:$KUBERNETES_SERVICE_PORT" \
  kubernetes_ca_cert=@/var/run/secrets/kubernetes.io/serviceaccount/ca.crt

k apply -f -<< EOF
apiVersion: v1
kind: ServiceAccount
metadata:
  name: vault-auth
  namespace: default
EOF

vault write auth/kubernetes/role/demo-role \
  bound_service_account_names=vault-auth \
  bound_service_account_namespaces=default \
  policies=mysecret \
  ttl=1h
  
  
printenv

apiVersion: apps/v1
kind: Deployment
metadata:
  name: vault-check
  namespace: default
spec:
  replicas: 1
  selector:
    matchLabels:
      app: vault-check
  template:
    metadata:
      labels:
        app: vault-check
    spec:
      containers:
        - name: vault-check
          image: hashicorp/vault:1.15.0
          command: ["/bin/sh"]
          args: ["-c", "sleep 3600"]
          env:
            - name: VAULT_ADDR
              value: "http://vault.default.svc.cluster.local:8200"
            - name: VAULT_SKIP_VERIFY
              value: "true"



ORRRRRRRRR


apiVersion: v1
kind: Pod
metadata:
  name: vault-demo
spec:
  serviceAccountName: vault-auth
  containers:
    - name: vault-demo
      image: badouralix/curl-jq
      command: [ "sh", "-c" ]
      args:
        - |
          VAULT_ADDR="http://vault.default.svc.cluster.local:8200"
          
          echo "Reading service account token..."
          SA_TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)

          echo "Authenticating to Vault..."
          LOGIN_RESPONSE=$(curl -s --request POST --data "{\"jwt\": \"${SA_TOKEN}\", \"role\": \"demo-role\"}" \
            "${VAULT_ADDR}/v1/auth/kubernetes/login")

          VAULT_TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.auth.client_token')

          echo "Fetching secret from Vault..."
          SECRET=$(curl -s -H "X-Vault-Token: ${VAULT_TOKEN}" \
            "${VAULT_ADDR}/v1/kv-v2/data/vault-demo/mysecret" | jq -r '.data.data')

          echo "✅ Secret retrieved:"
          echo "$SECRET"

          sleep 3600  # keep the pod running to inspect



kubectl exec -it deploy/vault-check -- sh

vault status


