# vecty-sample

dead-simple chat app sample

Demo: https://chatapp.irieda.com

## get sample go source code

```sh
go get -d github.com/nobonobo/vecty-chatapp
cd $GOPATH/src/github.com/nobonobo/vecty-chatapp
```

## development server start

```sh
make gopherjs &
make devrun
```

open: http://localhost:8888/

## local run native binary

```sh
make build
make run
```

open: http://localhost:8888/

## local run by docker

```sh
docker run -it --rm -p 8888:8888 nobonobo/chatapp
```

open: http://localhost:8888/
