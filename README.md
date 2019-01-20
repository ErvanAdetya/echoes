# echoes

`echoes` is a simple web application that returns anything it receives.

### Usage

`echoes` is listening on port 5000.

Create a POST HTTP request to `/` with JSON body payload:
```json
{
    "message": "act_1"
}
```

`echoes` will return a response similar to this:
```json
{
  "Message": "act_1",
  "Timestamp": "2019-01-20T12:28:43Z"
}
```
