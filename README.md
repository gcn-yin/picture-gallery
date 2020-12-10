# Picture Gallery

Usage: generate index.html that contains all pictures url in the same directory.

# Compile

```
go build -o pg main.go
```

# Example

If you have a directory like this structure.

```
$ ls
1.png
2.jpg
3.jpeg
```

You can enter

```
$ pg
```

It will generate a `index.html` file, and set up a static server.

You can open `http://localhost:8080` and view pictures.
