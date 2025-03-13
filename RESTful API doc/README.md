# What is REST API?

## REST Basics

* **REST** - Representational State Transfer
* **API** - Application Programming Interface

Most, if not all, large popular websites will rely upon some form of **REST API** in order to deliver some content or functionality to their users. Some sites like *Facebook* and *Twitter* actually expose some of these APIs to outside **developers** to build their own tools and systems.

We can communicate with **REST APIs** using **HTTP requests**, much like you’d do to navigate to a website or load an image. We can do HTTP requests to certain **API urls** and these urls would then return the information we required, or we could push data to an API url in order to change some data in a database.

Typically we send HTTP requests to an URL that we have defined in our REST API and it would either **perform a given task** for us or **return a certain bit of data**. Most APIs these days would return a response to us in the form of `JSON`.

![REST API](https://images.tutorialedge.net/uploads/rest-api.png)

## Example

Imagine you wrote a bit of code that gives you the current weather conditions at your house. It reads the `temperature`, `humidity` and `rainfall` and stores them locally. How would we then expose this information in such a way that websites or other applications could view it?

One answer to this question is by wrapping it in a **RESTful API**.

We could expose our code and *wrap it in an API* so that whenever we navigated to say `http://localhost:8000/api/weatherStats` it would give us a JSON response that contained all the current weather stats.

## Why Do We Do This?

* **Improved Code Reuse** - By exposing our code through **REST APIs** we essentially give ourselves a greater degree of flexibility. We can develop our software once and should we wish to use the same code again in a different project it would be easy, we could simply send *HTTP requests* to our **API** and we’ve reduced the need to duplicate our work efforts.

* **Always Available** - **REST APIs** are typically things that are running and available all the time. We make them very stable and as a result we can interact with them wherever we are in the world as long as we have internet connectivity.
