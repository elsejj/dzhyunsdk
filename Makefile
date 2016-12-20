
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
	GOOS=windows GOARCH=amd64 go build -o ${OUTPUTDIR}/${EXE_NAME}.exe main/main.go 
	GOOS=darwin GOARCH=amd64 go build -o ${OUTPUTDIR}/${EXE_NAME}.mac main/main.go 

zip:
	cd ${OUTPUTDIR} ; zip Dzhyunsdk.framework.zip Dzhyunsdk.framework
	cd ${OUTPUTDIR} ; zip dzhyunsdk.exe.zip dzhyunsdk.exe
	cd ${OUTPUTDIR} ; zip dzhyunsdk.linux.zip dzhyunsdk.linux
	cd ${OUTPUTDIR} ; zip dzhyunsdk.mac.zip dzhyunsdk.mac
	cd ${OUTPUTDIR} ; zip dzhyunsdk.aar.zip dzhyunsdk.aar

proto:
	make -C dzhyun
