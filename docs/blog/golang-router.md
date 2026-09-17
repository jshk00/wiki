---
title: Chi like routing methods in go 1.22 above
tags: ["go", "http", "rest-api"]
---

!!! warning
    This is applicable in new golang router which include ability to find path parameters.
    You can you use go1.22+ or go1.21 using experimental features enabled

## The missing apis of golang net/http
From version go 1.22 we have new `http.ServeMux` though you can also call it default router from standard library. It does exactly what third party routers used to do. Routing the paths using more efficient algorigthm in this case `Radix tree`. Still `net/http` have some methods missing for grouping and middlware, Good news is that we can just implement it by our own. Although not all feature from custom router like chi supported eg. regex matching. If you're willing to sacrifice that well let's dig in. 

## Getting started with router structure

```go
type Router struct {
	mux         *http.ServeMux
	middlewares []func(http.Handler) http.Handler
	prefix      string
}
```

We will keep all fields to private for now more on this later. we have three fields `mux` which is default http.Mux, `middlewares` holds the list of middleare to apply to perticular mux,`prefix` this is empty for root mux because there's no group in it; Though it will require when we start to create groups of mux and attaching those to parent.

## Designing generic http methods
```go
func (r *Router) Get(pattern string, handler http.HandlerFunc) {
	r.handle(http.MethodGet, pattern, handler)
}

func (r *Router) Head(pattern string, handler http.HandlerFunc) {
	r.handle(http.MethodHead, pattern, handler)
}

func (r *Router) Post(pattern string, handler http.HandlerFunc) {
	r.handle(http.MethodPost, pattern, handler)
}

func (r *Router) Put(pattern string, handler http.HandlerFunc) {
	r.handle(http.MethodPut, pattern, handler)
}

func (r *Router) Patch(pattern string, handler http.HandlerFunc) {
	r.handle(http.MethodPatch, pattern, handler)
}

func (r *Router) Delete(pattern string, handler http.HandlerFunc) {
	r.handle(http.MethodDelete, pattern, handler)
}

func (r *Router) Connect(pattern string, handler http.HandlerFunc) {
	r.handle(http.MethodConnect, pattern, handler)
}

func (r *Router) Options(pattern string, handler http.HandlerFunc) {
	r.handle(http.MethodOptions, pattern, handler)
}

func (r *Router) Trace(pattern string, handler http.HandlerFunc) {
	r.handle(http.MethodTrace, pattern, handler)
}

func (r *Router) handle(methodType, pattern string, handler http.Handler) {
	if len(r.prefix) > 0 && r.prefix[0] != '/' {
		panic("invalid grouping pattern")
	}
	pattern = fmt.Sprintf("%s %s%s", methodType, r.prefix, pattern)
	r.mux.Handle(pattern, chain(r.middlewares, handler))
}

func (r *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(res, req)
}
```
All generic http methods will call handle function which in term composed of logic for checking prefix is correct for group and make the path. go net/http provded a way to attach POST, PUT, GET in path string which may be errorprone that's why we are calling them statically to resolve in handler function itself. `ServeHTTP` implements the standard `http.Handler` interface which helps to directly register our router with `http.Server`.

## Crafting the middlware chaining
```go
// chain builds a http.Handler composed of an inline middleware stack and endpoint
// handler in the order they are passed.
func chain(middlewares []func(http.Handler) http.Handler, h http.Handler) http.Handler {
	// Return ahead of time if there aren't any middlewares for the chain
	if len(middlewares) == 0 {
		return h
	}

	// Wrap the end handler with the middleware chain
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
```
Code is pretty much self explaintory though what is does is it wraps all the middlware from last to first around our http handlers. now because we are wrapping it around the order of execution will be in order you provided the middlwares.

## Using the middlware and grouping
```go
func (r *Router) Group(opts ...func(*Options)) *Router {
	o := &Options{}
	for _, fn := range opts {
		fn(o)
	}
	im := r.with(o.mws...)
	im.prefix = r.prefix + o.prefix
	if o.handlers != nil {
		o.handlers(im)
	}
	return im
}

// Appened handlers for router or router group
func (r *Router) Use(mws ...func(http.Handler) http.Handler) {
	r.middlewares = append(r.middlewares, mws...)
}

func (r *Router) with(middlewares ...func(http.Handler) http.Handler) *Router {
	mws := make(Middlewares, len(r.middlewares))
	copy(mws, r.middlewares)
	mws = append(mws, middlewares...)
	return &Router{mux: r.mux, middlewares: mws}
}
```
`Group` function will create another group if you register it with prefix `/test` and your original url is `/example` the group url will be `/example/test`. Keep in mind that the middlware provide to group will only execute for perticualr group if want to apply globally try using `*(router).(Use())` method. `with` is just a helper function to facilitate copying of original router middlware and also append unique middlware related to group.

This `Group` function is powerd by the `functional option` pattern and finally you can understand it below.

## Options for grouping
```go
type Options struct {
	prefix   string
	mws      []func(http.Handler) http.Handler
	handlers func(r *Router)
}

func WithHandlers(h func(r *Router)) func(*Options) {
	return func(o *Options) {
		o.handlers = h
	}
}

func WithMiddlwares(mws ...func(http.Handler) http.Handler) func(*Options) {
	return func(o *Options) {
		o.mws = mws
	}
}

func WithPattern(s string) func(*Options) {
	return func(o *Options) {
		o.prefix = s
	}
}
```
You can follow along with any functional option pattern for how that works
