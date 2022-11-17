go build 2> .\g.tmp
gg < .\g.tmp
del .\g.tmp
go build
