# Embedded Device

## client
### Commands


    client show
    client push
    client pull
    client authorize [--revoke]

### Devspec

    self <- for the client
    all  <- all the devices
    oth  <- other specific device

### Command Register

#### All devices
    client register
##### API
    /d/reg/:deviceid/:password/:uuid
#### Publisher nodes
    client register --publisher     <- publisher node>
##### API
    /d/reg/:deviceid/:password/:uuid/pub



### API Show
    /d/show/:deviceid/:password/:devspec     <- show auth status of devspec
    /d/activity/:deviceid/:password/:devspec <- file transfer activity


### API Authorize
    /d/auth/:deviceid/:password/:devspec   <- authorize devspec
    /d/revoke/:deviceid/:password/:devspec <- revoke authorization


```
## server
### Commands
```
    server install
    server stop
    server start
    server restart
    server log
    server purge
```