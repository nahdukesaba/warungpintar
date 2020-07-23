# Simple Webserver

This is a simple webserver that built using go and gorilla mux library. This webserver consist of several API that can be easily access using these url

## API

### Welcome
This API uses GET method and recieve 1 input and can be access with
```
localhost:8080/welcome?name=Ephraim
```

### History
This API saves all API that has been hit before and display it as response
```
localhost:8080/history
```

### Score
This API uses websocket to make it real time API. This API collect and display the total number Kill button been clicked.
```
localhost:8080/score/
```

### SwaggerUI
For more detailed documentation of all the API, you can visit the swaggerui page available on this webserver
```
localhost:8080/swaggerui/
```

## How to run
After cloning this project, you just have to run these line of code to run this project
```
go run main.go
```
