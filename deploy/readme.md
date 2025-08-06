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
  --set provisionDataStore.cassandra=false \
  --set collector.enabled=true \
  --set query.enabled=true \
  --set agent.enabled=true \
  --set allInOne.enabled=true \
  --set storage.type=memory

```

```yaml
helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts
    helm repo update
helm upgrade --install opentelemetry-operator open-telemetry/opentelemetry-operator \
    --set admissionWebhooks.certManager.enabled=false
```
------

## Create OpenTelemetry Collector Config

```yaml
#otel-collector.yaml
apiVersion: opentelemetry.io/v1alpha1
kind: OpenTelemetryCollector
metadata:
  name: otel-collector
spec:
  mode: deployment
  config: |
    receivers:
      otlp:
        protocols:
          grpc:
          http:

    exporters:
      otlp:
        endpoint: "jaeger-collector.default:14250"
        tls:
          insecure: true

    service:
      pipelines:
        traces:
          receivers: [otlp]
          exporters: [otlp]
```

-----

## Instrument Your Golang App

-------

##  Deploy Golang App with Env Config

```yaml
env:
- name: OTEL_EXPORTER_OTLP_ENDPOINT
  value: "http://otel-collector.default.svc.cluster.local:4317"
- name: OTEL_SERVICE_NAME
  value: "my-golang-app"

```
-------

## Verify Integration
`kubectl port-forward svc/jaeger-query -n default 16686:16686`