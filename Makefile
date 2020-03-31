CWD=$(shell pwd)

go-bindata:
	mkdir -p cmd/go-bindata
	mkdir -p cmd/go-bindata-assetfs
	curl -s -o cmd/go-bindata/main.go https://raw.githubusercontent.com/whosonfirst/go-bindata/master/cmd/go-bindata/main.go
	curl -s -o cmd/go-bindata-assetfs/main.go https://raw.githubusercontent.com/whosonfirst/go-bindata-assetfs/master/cmd/go-bindata-assetfs/main.go

bake:
	@make bake-static
	@make bake-templates

bake-static:
	go build -mod vendor -o bin/go-bindata cmd/go-bindata/main.go
	go build -mod vendor -o bin/go-bindata-assetfs cmd/go-bindata-assetfs/main.go
	rm -f static/*~ static/css/*~ static/javascript/*~
	@PATH=$(PATH):$(CWD)/bin bin/go-bindata-assetfs -pkg www -o http/www/staticfs.go static static/javascript/ static/css/

bake-templates:
	rm -rf templates/html/*~
	bin/go-bindata -pkg templates -o assets/templates/html.go templates/html

debug:
	@make bake
	go run -mod vendor cmd/spatial-server/main.go -enable-www -enable-extras -extras-database 'reader:///?reader=fs:///usr/local/data/sfomuseum-data-maps/data&cache=gocache://' -nextzen-apikey $(APIKEY) -mode repo:// /usr/local/data/sfomuseum-data-maps/

