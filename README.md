zbubble
========

Fun game in Golang

[Try the game!](http://kostspielig.github.io/zbubble/)

Development
-----------

### Depedencies

For a full blown development environment, you need
[Golang](https://golang.org/), [Ebiten](https://hajimehoshi.github.io/ebiten/),
[GopherJS](https://github.com/gopherjs/gopherjs) and other
stuff.  On a Debian based distribution, you may run:

```
go get github.com/hajimehoshi/ebiten/...

go get -u github.com/gopherjs/gopherjs
go get -u github.com/gopherjs/webgl
```

These dependencies might be missing:

```
sudo apt-get install freeglut3-dev libalut-dev
```

### Compile

```
make build
```

### Compile and run

```
make
```

### Serving on the web

**TL;DR** run this and go to:
[localhost:8080/github.com/kostspielig/zbubble/](http://localhost:8080/github.com/kostspielig/zbubble/)

```
gopherjs serve
```

Before serving you must compile the JS version by running:

```
make buildweb
```


License
-------

![license](http://www.gnu.org/graphics/agplv3-155x51.png)

> Copyright (c) Maria Carrasco
>
> This file is part of zbubble.
>
> zbubble is free software: you can redistribute it and/or modify
> it under the terms of the GNU Affero General Public License as
> published by the Free Software Foundation, either version 3 of the
> License, or (at your option) any later version.
>
> zbubble is distributed in the hope that it will be useful, but
> WITHOUT ANY WARRANTY; without even the implied warranty of
> MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
> Affero General Public License for more details.
>
> You should have received a copy of the GNU Affero General Public
> License along with Mittagessen.  If not, see
> <http://www.gnu.org/licenses/>.
