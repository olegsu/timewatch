# timewatch

CLI for timewatch report application


## Start
Init first
```
timewatch init --help
```

Now Report
```
timewatch report --help
```

## Run from Codefresh pipeline
```
version: '1.0'

steps:
  Report:
    image: olsynt/timewatch
    environment:
    - TW_COMPONY_ID={{COMPONY_ID}}
    - TW_USER_ID={{YOUR_USER_ID}}
    - TW_PASSWORD={{YOUR_PASSWORD}}
    commands:
    # Change timezone for IL
    - export TW_TIME=$(TZ=UTC-2 date +%H-%M)

    # Login
    - timewatch init --verbose
    
    # Report checkin
    - timewatch report checking --verbose
  
```
