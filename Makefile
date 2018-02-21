
gopherjs:
	gopherjs serve github.com/nobonobo/vecty-sample/app

devrun:
	go run main.go -dev

build:
	mkdir -p ./dist/app/assets
	gopherjs build -m -o ./dist/app/app.js github.com/nobonobo/vecty-sample/app
	cp -f ./app/index.html ./dist/app/
	cp -Rf ./app/assets/* ./dist/app/assets/
	go build -o ./dist/server main.go

run:
	cd dist && ./server
