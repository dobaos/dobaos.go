## dobaos.go

This is not an official dobaos client library for Golang, just a basic example of how dobaos can be connected with programs in other programming languages.

## Usage

```go
dobaosClient := dobaos.New()

value, err := dobaosClient.GetValue(1)
if err != nil {
    // handle error
}

resp, err := dobaosClient.SetValue("value to be set")
if err != nil {
    // handle error
}
```