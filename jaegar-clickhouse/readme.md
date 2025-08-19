kubectl apply -f https://raw.githubusercontent.com/Altinity/clickhouse-operator/release-0.21.3/deploy/operator/clickhouse-operator-install-bundle.yaml
kc create -f jaegar-clickhouse/jaeger-operator-rbac.yaml
kubectl apply -n observability -f https://raw.githubusercontent.com/jaegertracing/jaeger-operator/v1.27.0/deploy/operator.yaml


Install cert-manager first (Recommended)

# Install cert-manager CRDs
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.3/cert-manager.crds.yaml

# Install cert-manager
helm repo add jetstack https://charts.jetstack.io
helm repo update

helm install cert-manager jetstack/cert-manager \
--namespace cert-manager \
--create-namespace \
--version v1.13.3

# Wait for cert-manager to be ready
kubectl wait --for=condition=ready pod -l app.kubernetes.io/instance=cert-manager -n cert-manager --timeout=120s

# Now install Jaeger Operator
helm install jaeger-operator jaegertracing/jaeger-operator \
--namespace observability \
--create-namespace \
--set rbac.create=true \
--set watchNamespace="observability"