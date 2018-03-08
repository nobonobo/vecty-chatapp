
gopherjs:
	gopherjs serve github.com/nobonobo/vecty-chatapp/app

devrun:
	go run main.go -dev

assets:
	mkdir -p ./dist/app/assets
	gopherjs build -m -o ./dist/app/app.js github.com/nobonobo/vecty-chatapp/app
	cp -f ./app/index.html ./dist/app/
	cp -Rf ./app/assets/* ./dist/app/assets/

build: assets
	go build -o ./dist/server main.go

run:
	cd dist && ./server

docker-build:
	docker build --rm -t nobonobo/chatapp .

docker-run:
	docker run -it --rm -p 8888:8888 nobonobo/chatapp .

docker-shell:
	docker run -it --rm -p 8888:8888 --entrypoint=sh nobonobo/chatapp
