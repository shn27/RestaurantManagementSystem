```yaml
helm repo add elastic https://helm.elastic.co
helm repo update
```

```yaml
helm install elasticsearch elastic/elasticsearch \
--set replicas=1 \
--set minimumMasterNodes=1 \
--set resources.requests.cpu="100m" \
--set resources.requests.memory="512Mi" \
--set volumeClaimTemplate.resources.requests.storage="1Gi" \
--set esJavaOpts="-Xms512m -Xmx512m"
```

`kubectl port-forward svc/elasticsearch-master 9200:9200
`

`curl http://localhost:9200
`