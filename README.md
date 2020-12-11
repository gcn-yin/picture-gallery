# Picture Gallery

Usage: generate index.html that contains all pictures url in the same directory.

# Compile

```
go build -o pg main.go
```

# Run

Enter a directory containing pictures.

```
$ pg .
```

It will generate a `index.html` file, and set up a static server.

You can open `http://localhost:8080` and view pictures.

# Cross compile

Install `gox`

```
go get github.com/mitchellh/gox
```

Run `gox`

```
gox
```
