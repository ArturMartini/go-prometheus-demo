# go-prometheus-demo
This project is example of usage prometheus with golang

## Case
We have service and expose endpoint GET /demo, for each request send metrics to prometheus to collect data

## Prometheus
Prometheus collects metrics from monitored targets by scraping metrics HTTP endpoints on these targets. 

Configuration:
```yaml
global:
  scrape_interval:     15s # By default, scrape targets every 15 seconds.

  # Attach these labels to any time series or alerts when communicating with
  # external systems (federation, remote storage, Alertmanager).
  external_labels:
    monitor: 'codelab-monitor'

# A scrape configuration containing exactly one endpoint to scrape:
# Here it's Prometheus itself.
scrape_configs:
  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.
  - job_name: 'myapp'

    # Override the global default and scrape targets from this job every 5 seconds.
    scrape_interval: 5s

    static_configs:
      - targets: ['localhost:8080']

```

## For run
```shell
cp prometheus.yml /tmp/
docker run --network host -v  /tmp/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus
docker run --network host go-prometheus
docker build . -t go-prometheus
curl http://localhost:8080/demo
curl http://localhost:9090/api/v1/query?query=myapp_processed_ops_total
```
  

If already success you need view look like json bellow:
```json
{
	"status": "success",
	"data": {
		"resultType": "vector",
		"result": [{
			"metric": {
				"__name__": "myapp_processed_ops_total",
				"instance": "localhost:8080",
				"job": "myapp"
			},
			"value": [1593530043.635, "1"]
		}]
	}
}
```

We can execute prometheus functions with query param:
```shell
curl http://localhost:9090/api/v1/query?query="rate(myapp_processed_ops_total[10s])
```