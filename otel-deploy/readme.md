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
--values otel-deploy/jaeger.yaml
```


## For Jaegar elasticsearch as DB
Working
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

May be with NodePort
```yaml
helm upgrade --install jaeger jaegertracing/jaeger \
--set provisionDataStore.cassandra=false \
--set provisionDataStore.elasticsearch=true \
--set allInOne.enabled=false \
--set storage.type=elasticsearch \
--set storage.elasticsearch.host=elasticsearch \
--set storage.elasticsearch.port=9200 \
--set esIndexCleaner.enabled=true \
--set esIndexCleaner.numberOfDays=3 \
--set esIndexCleaner.schedule="55 23 * * *" \
--set collector.enabled=true \
--set collector.otlp.enabled=true \
--set collector.otlp.protocols.grpc.port=4317 \
--set collector.service.type=NodePort \
--set collector.replicas=1 \
--set query.enabled=true \
--set query.replicas=1 \
--set agent.enabled=false \
--set elasticsearch.replicas=1 \
--set elasticsearch.minimumMasterNodes=1 \
--set elasticsearch.volumeClaimTemplate.resources.requests.storage=5Gi \
--set elasticsearch.resources.requests.memory=512Mi \
--set elasticsearch.resources.requests.cpu=200m


```

The previous helm upgrade has issue that is it's work fine with storage type = memory but got hanged with elasticsearch because of storage issue. We solved it using this which taking minimum memory.

### Cassandra as DB 
Don't work. jaegar pod get crashlooped.

```yaml
helm upgrade --install jaeger jaegertracing/jaeger \
--set provisionDataStore.cassandra=true \
--set provisionDataStore.elasticsearch=false \
--set allInOne.enabled=true \
--set storage.type=cassandra \
--set storage.cassandra.hosts=jaeger-cassandra \
--set storage.cassandra.keyspace=jaeger_v1_test \
--set storage.cassandra.datacenter=datacenter1 \
--set storage.cassandra.port=9042 \
--set cassandra.config.cluster_size=1 \
--set cassandra.persistence.size=5Gi \
--set cassandra.resources.requests.memory=512Mi \
--set cassandra.resources.requests.cpu=200m \
--set collector.enabled=false \
--set query.enabled=false \
--set agent.enabled=false


```
--------


## Deploy OpenTelemetry Collector
```yaml
  helm upgrade --install otel-collector open-telemetry/opentelemetry-collector \
     --values otel-collector-values.yaml \
     --set image.repository="otel/opentelemetry-collector-k8s"
```
--------------------

## Resource
[https://medium.com/@blackhorseya/deploying-opentelemetry-and-jaeger-with-helm-on-kubernetes-d86cc8ba0332 [this works like boom]]()

----------------

## Current Architechture

[https://chatgpt.com/share/689c556f-8fac-800c-97f0-118885fd20d7]


## Todo
```
💡 Rule of thumb:

Same cluster, same namespace → jaeger-collector:4317

Same cluster, different namespace → jaeger-collector.<namespace>.svc.cluster.local:4317

Different cluster or bare metal → Use IP or external DNS name


```