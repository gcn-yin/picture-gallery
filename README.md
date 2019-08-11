# Picture Gallery

Usage: generate index.html that contains all pictures url in the same directory.

# Compile

```
$ go build -o pg main.go
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

It will generate a `index.html` file, content is

```html
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<meta http-equiv="X-UA-Compatible" content="ie=edge">
<style type="text/css">
body { display: flex; flex-direction: column; align-items: center; }
</style>
<title>Picture Gallery</title>
</head>
<body>
<img src="1.png" alt="1.png"/>
<img src="2.jpg" alt="2.jpg"/>
<img src="3.jpeg" alt="3.jpeg"/>
</body>
</html>
```

And set up a static server.

```
$ pg
1.png
2.jpg
3.jpeg
generate index.html done
517 bytes written successfully
2019/07/01 09:37:13 Listening... http://localhost:8080
```

You can open `http://localhost:8080` and view pictures.

# Docker

If you want to use docker to set up environment to view pictures, you can use `hezhaohuang/pictures-gallery:1` image.

```
$ docker run -d -v /path/to/pictures:/workdir/pictures -p 8080:8080 --name pg-local hezhaohuang/pictures-gallery:1 
```

And you can open `http://localhost:8080` to view.