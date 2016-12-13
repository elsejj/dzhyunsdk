SED=sed

UNAME := $(shell uname)
ifeq (${UNAME}, Darwin)
SED=gsed
endif

go:
	protoc --go_out=. *.proto
	${SED} -i 's/int64/YFloat/g' *.pb.go
	${SED} -i 's/YFloat/int64/g' dzhpbtable.pb.go


