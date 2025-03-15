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

## Deprecating Older Versions

At some point, you may want to start deprecating and removing older versions of your API. There are many reasons as to why you might want to do this. From, removing additional endpoints that you have to constantly support, to possibly removing the feature entirely. When it comes to managing this, you need to be careful and ensure that you give your customers a massive lead time in which they can migrate their applications to use newer versions of the system. 

**Note** - Having good monitoring and logging systems in place within your APIs can drastically help improve this process. If you can accurately map which of your clients is using which version of the endpoint, you can focus your efforts on migrating these clients off and ensure nobody is still using the version before removing it entirely.

## Use Appropriate HTTP Verbs
When it comes to writing APIs it is incredibly important to use the correct HTTP verbs when it comes to creating your API endpoints and don’t include verbs in your API path.

**For example**, don’t define endpoints like this:

- /api/v1/getposts
- /api/v1/newpost
- /api/v1/updatepost/:id
- /api/v1/deletepost/:id

If another developer is going to be using your API then they won’t instinctively know what endpoints they have to use to perform basic CRUD operations. 

The better approach would be to **use appropriate HTTP verbs instead**:

- /api/v1/post - HTTP GET request - All Posts
- /api/v1/post/:id - HTTP GET request - Single post
- /api/v1/post/:id - HTTP POST request - Publish a post
- /api/v1/post/:id - HTTP PATCH request - update an existing post
- /api/v1/post/:id - HTTP DELETE request - deletes a post

**Note** - An excellent example of this approach is the `Open Service Broker API Spec`

## Use of HTTP Response Codes

When it comes to returning a response to whomever called your API then it’s best to utilize the correct HTTP status codes.

- 1xx ## Informational Status Codes
- 2xx ## Success status codes
- 3xx ## Redirection status codes
- 4xx ## Client Error status codes
- 5xx ## Server Error status codes

### Example
When a standard user calls `GET` on the first API you would typically return the `json` response as well as a `200 - OK` status code which indicates everything is `ok`. However, imagine that same user tries to update our article using the `UPDATE` endpoint. As the user isn’t an administrator, we would want to return a different status code such as `401 - Unauthorized` which would indicate that the user does not have the sufficient level or permissions. Once the user becomes authorized as an `admin` and tries to `UPDATE` again to that same endpoint, we then return a `200 - OK` status which indicates they were successfully able to update the article.

## Media Types

If you are building a general purpose API that could be used for a massively variety of reasons then you need to consider adding different responses based on the `Content-Type` **header** passed in with the request.

This gives the developer trying to interact with your application the option of requesting a response in a variety of different data formats.

```go
- /api/v1/posts - HTTP GET - Content-Type: application/json

{ "posts": [
    ... all the posts in JSON
]}

```
Imagine we needed to feed the information from this into a build pipeline but needed the posts in a `yaml` format. Ideally, our API would be able to serialize the response into `yaml` and return it if the `Content-Type` was set to **application/yaml**:

```go
- /api/v1/posts - HTTP GET - Content-Type: application/yaml

posts:
 ... all posts in yaml format
```

## Swagger

**Swagger** is something that makes quickly testings and validating **API** endpoints a treat. If you are building an **API** that is going to be consumed by a wide variety of users then providing a page which allows them to instantly interact with your API offers huge value for something that takes incredibly minimal amounts of effort to set up.
