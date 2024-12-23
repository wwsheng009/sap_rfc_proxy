# sap rfc proxy

this project the main purpose is build as rest server proxy for calling the sap rfc functions.

use the project https://github.com/SAP/gorfc.git as the sap rfc client.

use gogin as the web framework.

export service endpoint /rfc/{service},service name is the sap rfc function name.

all the functions are exposed as http post requests.

the http request forwards to the sap rfc function.

the input json payload mapping to the sap abap rfc function parameters. 

type mappings:

- time mapping to ABAP DATE 'YYYYMMDD'
- boolean mapping to ABAP BOOL 'X' or ' '
- json struct mapping to sap abap rfc struct
- json array of objects mapping to sap abap rfc table

wrapper the sap rfc function return as json object to service return.

write the log to the log file file.


example code for connecting the sap server,read the connection infomation from the environment file .env
```go
import (
    "fmt"
    "github.com/sap/gorfc/gorfc"
    "github.com/stretchr/testify/assert"
    "reflect"
    "testing"
    "time"
)

func abapSystem() gorfc.ConnectionParameter {
    return gorfc.ConnectionParameter{
        Dest:      "I64",
        Client:    "800",
        User:      "demo",
        Passwd:    "welcome",
        Lang:      "EN",
        Ashost:    "11.111.11.111",
        Sysnr:     "00",
        Saprouter: "/H/222.22.222.22/S/2222/W/xxxxx/H/222.22.222.222/H/",
    }
}
```

gorfc test case code:

https://raw.githubusercontent.com/SAP/gorfc/refs/heads/master/gorfc/gorfc_test.go


## sap rfc install

https://sap.github.io/PyRFC/install.html#install-c-connector

Set the SAPNWRFC_HOME env variable: SAPNWRFC_HOME=c:\nwrfcsdk
Include the lib directory to the library search path: PATH=PATH;%SAPNWRFC_HOME%\lib
```sh
# for run
export SAPNWRFC_HOME=$HOME/nwrfcsdk
export LD_LIBRARY_PATH="/usr/local/sap/nwrfcsdk/lib"
```
befure build the go project,set the sap rfc sdk path.

need to overwrite the environment variable.

- CGO_LDFLAGS, for the link object in c.
- CGO_CPPFLAGS, for the include source file from c.

```sh
export SAPNWRFC_HOME="E:/sap/nwrfcsdk/lib"

CGO_LDFLAGS="-L E:/sap/nwrfcsdk/lib" CGO_CFLAGS="-I E:/sap/nwrfcsdk/include" go build
```


## sap connection parameter

https://help.sap.com/doc/saphelp_nw73ehp1/7.31.19/en-us/48/ce50e418d3424be10000000a421937/frameset.htm


## add env variable


```sh
vi ~/.bashrc

## append lines to the end of the file.

export SAPNWRFC_HOME=/root/projects/nwrfcsdk
export LD_LIBRARY_PATH=LD_LIBRARY_PATH:$SAPNWRFC_HOME/lib
```

## service

```sh
mkdir -p /opt/sap_rfc_proxy
cp ./sap_rfc_proxy /opt/sap_rfc_proxy/

cp ./.env.sample /opt/sap_rfc_proxy/.env

sudo chmod +x /opt/sap_rfc_proxy/sap_rfc_proxy

cp ./sap_rfc_proxy.service /etc/systemd/system/sap_rfc_proxy.service

sudo systemctl daemon-reload

sudo systemctl start sap_rfc_proxy.service
sudo systemctl stop sap_rfc_proxy.service
sudo systemctl restart sap_rfc_proxy.service
sudo systemctl status sap_rfc_proxy.service
sudo systemctl enable sap_rfc_proxy.service
sudo systemctl disable sap_rfc_proxy.service
```