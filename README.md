"LinkPreview" backend
===

This is backend repository for the "LinkPreview" tool I created that takes a URL as input from user and displays a "preview" for that link by doing some runtime web-scraping.

[It's deployed here.](http://194.195.117.9/)

## Stack used
- Vue with typescript (+ tailwindcss) for frontend ([its repo](https://github.com/SHD-420/linkpreview-frontend))
- Golang server as backend (this repo)

## About the backend
This backend serves two main purposes:
- First is to serve the static content from __client__ directory.
- Second is a route that generates preview for a link input.

The first task is accomplished simply by `http.FileServer` method. For second, I had to integrate web-scraping capibilities. I've used the [colly](https://pkg.go.dev/github.com/gocolly/colly/v2) package to do so.


## Deployment strategy

The project is deployed on a Linode machine running linux. To deploy the app, I've to first pull from this repo and run this command:
```console
bash update.sh
```
It executes the shell script "update.sh" which clones the frontend repo's deploy branch into __client__ directory at project root so that the frontend bundle is available for serving.

Then, I build and run the application binary in background. Finally, I've to configure nginx to __proxy_pass__ requests to `localhost:8000`.
