![www](./www.png)

# www

A speed dial for your browser, in your terminal

## Install and Use

The easiest way to install `www` is with `go install`. Please have Go installed and `$GOPATH/bin` in your `$PATH`

(If you'd rather build manually please clone the repo and `go build` directly)

1. Install
```sh
go install github.com/jkulton/www@latest
```

2. Create a `.www` file in your home directory, add a JSON object to the file. Each key should be the name of your bookmark, and each value the URL.

```sh
echo '{\n  "gh": "https://github.com/"\n}' > ~/.www
```

3. Execute `www` with the name of a bookmark
```sh
www gh # opens default browser to https://github.com/
```

4. Define more bookmarks in `~/.www`!
