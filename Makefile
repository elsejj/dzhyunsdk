
PACKAGE=github.com/elsejj/dzhyunsdk
OUTPUTDIR=output
JAVA_PACKAGE=com.dzhyun
OBJC_PREFIX=Dzhyun


android:
	mkdir -p $(OUTPUTDIR)
	cd ${OUTPUTDIR} && gomobile bind -target=android -javapkg=${JAVA_PACKAGE} ${PACKAGE}

ios:
	mkdir -p $(OUTPUTDIR)
	cd ${OUTPUTDIR} && gomobile bind -target=ios -prefix=${OBJC_PREFIX} ${PACKAGE}

proto:
	make -C dzhyun
