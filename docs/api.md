# API

## List `VirtualBox` VMs _(via API)_

```sh
 $ http --verify=no https://127.0.0.1:3993/vbox/list
```

```http
HTTP/1.1 500 Internal Server Error
Content-Length: 85
Content-Type: application/json; charset=UTF-8
Date: Mon, 19 Mar 2018 22:03:43 GMT

VBoxManage not found. Make sure VirtualBox is installed and VBoxManage is in the path
```
