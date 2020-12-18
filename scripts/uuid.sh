
token_name=$(kubectl get sa -n kube-system tpaas -o yaml | grep tpaas-token | awk '{print $NF}')
kubectl get secret ${token_name} -n kube-system -o yaml  | grep 'token:' | awk '{print $NF}' | base64 -d
