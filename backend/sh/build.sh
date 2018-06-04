#!/bin/sh
go build -ldflags "-s -w" -x -o ../spiders/usa/UIP ../spiders/usa/ippool.go
go build -ldflags "-s -w" -x -o ../spiders/uk/UIP ../spiders/uk/ippool.go
go build -ldflags "-s -w" -x -o ../spiders/jp/UIP ../spiders/jp/ippool.go
go build -ldflags "-s -w" -x -o ../spiders/de/UIP ../spiders/de/ippool.go

go build -ldflags "-s -w" -x -o ../spiders/usa/UURL ../spiders/usa/urlpool.go
go build -ldflags "-s -w" -x -o ../spiders/uk/UURL ../spiders/uk/urlpool.go
go build -ldflags "-s -w" -x -o ../spiders/jp/UURL ../spiders/jp/urlpool.go
go build -ldflags "-s -w" -x -o ../spiders/de/UURL ../spiders/de/urlpool.go

go build -ldflags "-s -w" -x -o ../spiders/usa/ULIST ../spiders/usa/listmain.go
go build -ldflags "-s -w" -x -o ../spiders/uk/ULIST ../spiders/uk/listmain.go
go build -ldflags "-s -w" -x -o ../spiders/jp/ULIST ../spiders/jp/listmain.go
go build -ldflags "-s -w" -x -o ../spiders/de/ULIST ../spiders/de/listmain.go

go build -ldflags "-s -w" -x -o ../spiders/usa/UASIN ../spiders/usa/asinmain.go
go build -ldflags "-s -w" -x -o ../spiders/uk/UASIN ../spiders/uk/asinmain.go
go build -ldflags "-s -w" -x -o ../spiders/jp/UASIN ../spiders/jp/asinmain.go
go build -ldflags "-s -w" -x -o ../spiders/de/UASIN ../spiders/de/asinmain.go


go build -ldflags "-s -w" -x -o ../spiders/usa/USQL ../spiders/usa/initsql.go
go build -ldflags "-s -w" -x -o ../spiders/uk/USQL ../spiders/uk/initsql.go
go build -ldflags "-s -w" -x -o ../spiders/jp/USQL ../spiders/jp/initsql.go
go build -ldflags "-s -w" -x -o ../spiders/de/USQL ../spiders/de/initsql.go