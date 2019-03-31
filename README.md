# Lucas

![lucas-banner.jpg](/lucas-banner.jpg)

Lucas is a friendly spider/web crawler made in go. It's still quite young so it can only deal with simple HTML webpages.

## How to run

This program runs from the terminal and expects a url as an argument. The output will result in a newline separated list of urls and a counter of how many times that entry is repeated.

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

## How to build

One can run `make bin` to generate usable binaries for different systems.
