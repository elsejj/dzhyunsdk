
PACKAGE=github.com/elsejj/dzhyunsdk
OUTPUTDIR=output
JAVA_PACKAGE=com.dzhyun
OBJC_PREFIX=Dzhyun
EXE_NAME=dzhyunsdk


android:
	mkdir -p $(OUTPUTDIR)
	cd ${OUTPUTDIR} ; gomobile bind -target=android -javapkg=${JAVA_PACKAGE} ${PACKAGE}

ios:
	mkdir -p $(OUTPUTDIR)
	cd ${OUTPUTDIR} ; gomobile bind -target=ios -prefix=${OBJC_PREFIX} ${PACKAGE}

exe:
	mkdir -p $(OUTPUTDIR)
	GOOS=linux GOARCH=amd64 go build -o ${OUTPUTDIR}/${EXE_NAME}.linux main/main.go 
	GOOS=windows GOARCH=amd64 go build -o ${OUTPUTDIR}/${EXE_NAME} main/main.go 
	GOOS=darwin GOARCH=amd64 go build -o ${OUTPUTDIR}/${EXE_NAME}.mac main/main.go 



proto:
	make -C dzhyun
