## Till now
```yaml
kc create -f deploy/mysql-deployment.yaml
docker exec -it db /bin/sh -c "mysql --host localhost -u<user-name> -p<password> -e 'CREATE DATABASE IF NOT EXISTS food_delivery;'"
kc create -f deploy/redis-deployment.yaml
kc create -f deploy/go-app-deployment.yaml
```

----------------

## Use the customized configuration to deploy Jaeger:
```yaml
helm upgrade --install jaeger jaegertracing/jaeger \
--history-max 3 \
--values jaeger.yaml
```


## For Jaegar
```yaml
helm upgrade --install jaeger jaegertracing/jaeger \
  --set provisionDataStore.cassandra=false \
  --set provisionDataStore.elasticsearch=true \
  --set allInOne.enabled=true \
  --set storage.type=elasticsearch \
  --set storage.elasticsearch.host=elasticsearch \
  --set storage.elasticsearch.port=9200 \
  --set esIndexCleaner.enabled=true \
  --set esIndexCleaner.numberOfDays=3 \
  --set esIndexCleaner.schedule="55 23 * * *" \
  --set provisionDataStore.elasticsearch=true \
  --set elasticsearch.replicas=1 \
  --set elasticsearch.minimumMasterNodes=1 \
  --set elasticsearch.volumeClaimTemplate.resources.requests.storage=5Gi \
  --set elasticsearch.resources.requests.memory=512Mi \
  --set elasticsearch.resources.requests.cpu=200m \
  --set collector.enabled=false \
  --set query.enabled=false \
  --set agent.enabled=false
```

The previous helm upgrade has issue that is it's work fine with storage type = memory but got hanged with elasticsearch because of storage issue. We solved it using this which taking minimum memory.

--------


## Deploy OpenTelemetry Collector
```yaml
  helm install otel-collector open-telemetry/opentelemetry-collector \
     --values otel-collector-values.yaml \
     --set image.repository="otel/opentelemetry-collector-k8s"
```
--------------------

## Resource
[https://medium.com/@blackhorseya/deploying-opentelemetry-and-jaeger-with-helm-on-kubernetes-d86cc8ba0332 [this works like boom]]()

----------------
