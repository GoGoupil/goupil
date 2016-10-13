**DEPRECATED: WILL BE REMOVED SOON, AND REPLACED BY A NEW TOOL USING HTTPTRACE FROM GO 1.7**

# Goupil
Goupil is a Go HTTP load testing tool allowing you to test your web servers loading capacity. You can check some values like the average response time of the server or the error rate. For more details, read the documentation on http://goupil.devatoria.ovh.

# Plan description example
This is a JSON plan description example.

```json
{
	"Host": "devatoria.info",
	"Port": 80,
	"Https": false,
	"Threads": [
		{
			"Duration": 5000,
			"Gap": 100,
			"Count": 10,
			"Route": "/",
			"Method": "GET"
		}
	]
}
```

# Increase "Open files" limit
Goupil must open a lot of sockets (and files) to run load tests. There are chances that you will hit max open files limit allowed by linux. To avoid this problem, you should increase this limit.

Edit `/etc/security/limits.conf` file and add the following lines at the end, then logout and login again.

```
*         hard    nofile      999999
*         soft    nofile      999999
root      hard    nofile      999999
root      soft    nofile      999999
```

## Remember
Goupil should be used on testing machines only and not on production servers or other critical machines. So, you must be sure that editing this configuration will not affect your machine security!

## Also
Thank you to https://rtcamp.com/tutorials/linux/increase-open-files-limit/ for this trick ;)
