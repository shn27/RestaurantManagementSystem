# How to Use Jaeger with OpenTelemetry Operator

---------

## High-Level Steps
Install Dependencies 

Deploy Jaeger via Helm

Deploy OpenTelemetry Operator via Helm

Create OpenTelemetry Collector Config

Instrument Your Golang App

Deploy the Golang App with OTel Sidecar/Agent

Verify Traces in Jaeger UI

-------

##  Deploy Jaeger, OTEL via Helm

```azure

helm repo add jaegertracing https://jaegertracing.github.io/helm-charts
helm repo update

helm upgrade --install jaeger jaegertracing/jaeger \
  --namespace observability --create-namespace \
  --set provisionDataStore.cassandra=false \
  --set collector.enabled=true \
  --set query.enabled=true \
  --set agent.enabled=true \
  --set allInOne.enabled=true \
  --set storage.type=memory

```

```azure
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
    helm repo update

helm upgrade --install opentelemetry-operator open-telemetry/opentelemetry-operator \
    --namespace observability \
    --create-namespace
```
------

## Create OpenTelemetry Collector Config

```yaml
#otel-collector.yaml
apiVersion: opentelemetry.io/v1alpha1
kind: OpenTelemetryCollector
metadata:
  name: otel-collector
  namespace: observability
spec:
  mode: deployment
  config: |
    receivers:
      otlp:
        protocols:
          grpc:
          http:

    exporters:
      otlphttp:
        endpoint: "http://jaeger-collector.observability.svc.cluster.local:4318"
        tls:
          insecure: true

    service:
      pipelines:
        traces:
          receivers: [otlp]
          exporters: [otlphttp]  # <-- Changed from "jaeger" to "otlphttp"
```

-----

## Instrument Your Golang App

-------

##  Deploy Golang App with Env Config

```yaml
env:
- name: OTEL_EXPORTER_OTLP_ENDPOINT
  value: "http://otel-collector.observability.svc.cluster.local:4317"
- name: OTEL_SERVICE_NAME
  value: "my-golang-app"

```
-------

## Verify Integration
`kubectl port-forward svc/jaeger-query -n observability 16686:16686`