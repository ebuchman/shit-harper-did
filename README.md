# Shit Harper Did

This repository is a self-contained "Shit Harper Did" website.

It is hosted at http://shd.coinculture.info

See the content in `site/shit-harper-did.md`

Serve the site by running the docker container:

```
docker build -t shd .
docker run --name shd -d -p 80:80 shd
```

You can regenerate `index.html` from the `shit-harper-did.md` by running `go run main.go` in the `sites` directory (you have to have [go installed](https://golang.org/doc/install))

Please feel free to contribute to making it look pretty, and to running mirrors.
