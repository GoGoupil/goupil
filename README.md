# Goupil
Goupil is a Go HTTP load testing tool allowing you to test your web servers loading capacity. You can check some values like the average response time of the server or the error rate. For more details, read the documentation on http://goupil.devatoria.ovh.

# Plan description example
This is a JSON plan description example.

```json
{
    "Host": "devatoria.info",
    "Port": 80,
    "Threads": [
        {
            "Count": 10,
            "Route": "/"
        },
		{
		    "Count": 10,
		    "Route": "/test"
		}
    ]
}
```
