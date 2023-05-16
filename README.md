# go-log
Golang logging implementation with echo framework and uber zap

Feature :
- Capture request on echo middleware and write to session
- Create log file base on configurable properties
- Capture log to console
- Standarization log format
- Available for INFO, DEBUG, FATAL, ERROR, PANIC

## Log Output Examples

```json
{
  "timestamp": "2023-05-16 15:30:26.961",
  "level": "info",
  "message": "Hello From Logger",
  "app_name": "github.com/denysetiawan28/go-log",
  "app_tag": "go-log",
  "app_version": "1.0.0",
  "app_port": 9090,
  "app_thread_id": "",
  "app_journey_id": "",
  "app_req_ip": "::1",
  "app_method": "GET",
  "app_uri": "/api/v1/hello",
  "app_req_body": "eyJhc2QiOiJhc2Rhc2QifQ==",
  "app_additional_data": null
}
```
