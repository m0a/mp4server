
#!/bin/sh
go-bindata-assetfs -pkg="main" assets/...
go install github.com/m0a/mp4server
rehash
