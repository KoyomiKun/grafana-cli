# Change Tag

modify grafana dashboard alert tags 

## Config

```json
{
  "api_key": "key",
  "alerts": [
    {
      "alert_name": "alert1",
      "tags": {
		  // if key exist, modify value
		  // or, create tag with KV
		  "key": "value"
      }
    },
    {
      "alert_name": "alert2",
      "tags": {
		  // if key exist and value is empty, delete this tag
		  "key": ""
      }
    }
  ]
}
```

## Usage

`grafana-cli alert tags` 
