## Use the customized configuration to deploy Jaeger:
```yaml
helm upgrade --install jaeger jaegertracing/jaeger \
--history-max 3 \
--values jaeger.yaml
```
## Deploy OpenTelemetry Collector
```yaml
  helm install otel-collector open-telemetry/opentelemetry-collector \
     --values otel-collector-values.yaml \
     --set image.repository="otel/opentelemetry-collector-k8s"
```

## Resource
https://medium.com/@blackhorseya/deploying-opentelemetry-and-jaeger-with-helm-on-kubernetes-d86cc8ba0332 [this works like boom]