![www](./www.png)

# www

A web speed dial in your terminal

## Install

Before installing please have Go installed and `$GOPATH/bin` in your `$PATH`

```sh
go install https://github.com/jkulton/www
```

## Set up

1. Create a JSON file titled '.www' in your home directory
2. Add a single object to the file, where each key is the name of your bookmark, and each value is the URL
3. Execute `www` with the name of a bookmark

Example .www file:
```
{
  "cf": "https://dash.cloudflare.com",
  "gh": "https://github.com/",
  "gm": "https://gmail.com"
}
```

## Usage

```sh
$ www <bookmark>

$ www gh # (opens https://github.com/ in the default browser)
```

---

### Build binary manually

If you'd rather install manually, clone the repository. Learn more about building binaries with `go build` [here](https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies). Once you build the binary, ensure it is accessible from your `$PATH`.# www
