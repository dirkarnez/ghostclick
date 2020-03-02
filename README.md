ghostmouse
===============
Prevents the computer from sleeping when idle. Rewrite [https://code.joejag.com/2013/move-your-mouse-pointer-with-java.html](https://code.joejag.com/2013/move-your-mouse-pointer-with-java.html) in Golang for getting rid of JVM

### Build
```
go build -ldflags "-linkmode external -extldflags -static" -a
```
