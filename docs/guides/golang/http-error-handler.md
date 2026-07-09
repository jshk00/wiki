---
title: Handling HTTP Error in golang
---

Generally we need naked return to handle some kind of error in default http handler in golang. This is also true for framework like Gin and router like Chi too. So in order to handle errors in default http handler you will do as in following snippet.
```golang
func GetHandler(w http.ResponseWriter, r *http.Request) {
    err := errors.New("some error")
    if err != nil {
        http.Error(err)
        return 
    }
    w.Write([]byte("successful response"))
}
```

As a human we may eventually forgot to return just after writing error. so we are exploring ways to handle error gracefully leaveraging the interface implementation of golang. If you look at the core of http handler it's just a interface. 😄 Good thing for us that we can easily implement it. Lets dig in.

## Error handling in standard http handler by creating custom handler method
Golang have excellent http library built in but it does not offer returning error for convient use case. 
where other router like echo does offer error returing handler. now in this post we are going to look at creating http error returning handler.
how create how to execute it.

## Creating custom http handler
What is http handler in go or rather what is `http.HandlerFunc` is?
In simple term http.HandlerFunc is whatever that implements `ServeHTTP(w http.ResponseWriter, r *http.Request)`.
See subsequent snippet to follow how we are creating custom handler which returns error.

```go linenums="1"
package main

import (
    "http"
    "errors"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)	    	
		return
	}
}

func main() {
    mux := http.NewServeMux()
    mux.Handle("/", HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
        return errors.New("this is error")
    }))
    http.ListenAndServe(":3000", mux)
}
```

- On line number 10 we implmented ServeHTTP method on our handler function
- Line number 11 executes our incoming handler and reurns any error from it
- On line 19 we converted `http.HandlerFunc` to our custom `HandlerFunc`

notice that this type of error handling now we can't specifiy status code and serve http will always return 500 as status.
we will look at how to tackle it in next section.

## Creating custom error for better response
- Basic idea here is we will create struct that implments both error and Marshler interfaces. 
- We need to implement Marsheler as golang doesn't know how to json encode errors
- Notice in MarshalJSON method we are returning custom struct with all field exported to encode json.

```go linenums="1"
--8<-- "golang/error_handler/handler.go:41:67"
```

## Changing our ServeHTTP to handle custom error

```go linenums="1"
--8<-- "golang/error_handler/handler.go:26:37"
```

## Complete Code

```go linenums="1"
--8<-- "golang/error_handler/handler.go"
```
