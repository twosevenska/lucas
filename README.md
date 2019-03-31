# Lucas

![lucas-banner.jpg](/lucas-banner.jpg)

Lucas is a friendly spider/web crawler made in go. It's still quite young so it can only deal with simple HTML webpages.

## How to run

This program runs from the terminal and expects a url as an argument.

### Example

```shell
go run lucas.go https://golangweekly.com/

/rss/25ccg42o 1
/favicon.png 1
/issues 1
/rss/1kgddaf9 1
/issues/255 1
https://cooperpress.com/ 1
/css/app.css 1
/ 1
/latest 1
```

### Advanced

One can also provide an optional `lines` (default:10) argument to control how many lines each chunk will have to be processed by a worker.

## How to build

One can run `make bin` to generate usable binaries for different systems.
