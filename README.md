# go-kostal-piko

A small go lib port from [Kostal Piko Dataexporter](https://gitlab.com/svij/kostal-dataexporter/-/tree/master) for grabbing Kostal Piko Inverter data.

## build and run the executable

```bash
Usage: ./do
	 go-fmt              format go code
	 build-cli [OS]      builds the go executable for cli
	 build-influx [OS]   builds the go executable with influx for container usage
	 build-container     builds the container image
```

the cli version works with parameters and the container version uses environment variables

## Dockerfile for K8s

if you like to use the dockerfile within your cluster, you need to set at least the following config:

```yaml
ports:
- name: http
    containerPort: 8080
    protocol: TCP
env:
- name: KOSTAL_URL
    value: <kostal-forwarding-service>.<namespace>
- name: INFLUXDB_HOST
    value: http://<influxdb-service>.<namespace>
- name: INFLUXDB_PORT
    value: "<influxdb port>"
- name: INFLUXDB_DB
    value: "<target db>"
```
