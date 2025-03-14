# Versioning your API

The **first**, and quite frankly **most important** point that you should consider when designing a *REST API* is that the contract for all your endpoints is **immutable**.

### Example
Let’s look at an example of this in action. Say I had an *endpoint* which, when you hit that *endpoint*, it would return a `JSON` structure with a response.

```go
http.GET("https://api.tutorialedge.net/api/hello")

Output: { "response": "hello" }
```

Imagine that we had customers who started using this *endpoint* within their own services for whatever reason. Now, imagine how upset you would be if your clients were to suddenly see that their systems were breaking as you had updated the *endpoint* to return the following:

```go
http.GET("https://api.tutorialedge.net/api/hello")

Output: { "content": "hello" }
```

This may appear to have been a subtle change, but whoever was consuming this endpoint had built up their own application around the previous `JSON` structure and expects this to never change.

### Solution
```go
http.GET("https://api.tutorialedge.net/api/v1/hello")

Output: { "response": "hello" }
```

Now, if we wanted to modify the response of our `JSON` for any reason to use the new **content** key instead, we could simply add a new endpoint which uses `v2` like so:

```go
http.GET("https://api.tutorialedge.net/api/v2/hello")

Output: { "content": "hello" }
```

Now, if you are rolling out new versions all the time, then you may want to use a more granular version number in the path parameters. `i.e. /api/v1-1/hello ` so that you can roll out minor changes quicker **without** necessarily upgrading the major version entirely.

