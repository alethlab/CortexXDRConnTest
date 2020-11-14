# CortexXDRConnTest

CortexXDRConnTest is a simple cli tool test your endpoint's connection to Cortex XDR communication servers and storage buckets.

### Get CortexXDRConnTest

| Supported Platform | Download                                                | Version |
| :----------------- | ------------------------------------------------------- | ------- |
| MacOS              | [zip](./binaries/v03/macos/CortexXDRConnTest.zip)       | v0.3    |
| Linux amd64        | [zip](./binaries/v03/linux_amd64/CortexXDRConnTest.zip) | v0.3    |
| Windows            | [zip](./binaries/v03/windows/CortexXDRConnTest.zip)     | v0.3    |



# Usage Documentation

### Getting Help

`CortexXDRConnTest -h , --help`

![CLI Help](./images/readme/help.png)

### Test Connection 

`CortexXDRConnTest -region us -tenant testtenant`

![test connection](./images/readme/test_connection.png)

### Test Connection with Proxy

`CortexXDRConnTest -region us -tenant testtenant -proxy 192.168.1.223:8888`

![Test with Proxy](./images/readme/test_connection_proxy.png)