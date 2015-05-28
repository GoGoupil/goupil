# Goupil
Go web load testing tool

# Plan description example

This is a JSON plan description example.

```json
{
    "Name": "test",
    "BaseURL": "http://www.google.com",
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
