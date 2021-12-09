# go-turbo-frame

This repo has been created mainly as an example to debug the following issue.

https://github.com/hotwired/turbo/issues/493

To run, execute the following.

```
go build
./turbo-frame
```

And check the service at http://localhost:8080

## Answer

In the end, the response returning the frame did not set the `Content-Type` to `text/html`, so Turbo Drive did not swap.

This has been fixed in the following commit and in release 0.0.2.

https://github.com/lobre/go-turbo-frame/commit/6b42ae553d12338dd7a7978f1fce66cb5ebc2b97

And to give more details, when the content type is not set explicitly, `net/http` tries to detect it.
This is done in the [DetectContentType](https://pkg.go.dev/net/http#DetectContentType) function.

But as the only tag in the response is a custom element (`<turbo-frame>`), it isn’t recognized as `text/html`.
