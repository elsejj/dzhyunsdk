
PACKAGE=github.com/elsejj/dzhyunsdk
OUTPUTDIR=output


android:
	mkdir -p $(OUTPUTDIR)
	cd ${OUTPUTDIR} && gomobile bind -target=android ${PACKAGE}

ios:
	mkdir -p $(OUTPUTDIR)
	cd ${OUTPUTDIR} && gomobile bind -target=ios ${PACKAGE}

proto:
	make -C dzhyun
