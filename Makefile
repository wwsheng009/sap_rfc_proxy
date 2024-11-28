fmt:
	go fmt ./...

build:fmt
	export LD_LIBRARY_PATH="/mnt/e/sap/nwrfc750P_15-linux/nwrfcsdk/lib"
	export SAPNWRFC_HOME="/mnt/e/sap/nwrfc750P_15-linux/nwrfcsdk"
	export CGO_LDFLAGS="-L /mnt/e/sap/nwrfc750P_15-linux/nwrfcsdk/lib"
	export CGO_CPPFLAGS="-I /mnt/e/sap/nwrfc750P_15-linux/nwrfcsdk/include"
	go build -o sap_rfc_proxy

run:build
	export LD_LIBRARY_PATH=/mnt/e/sap/nwrfc750P_15-linux/nwrfcsdk/lib && \
	export SAPNWRFC_HOME=/mnt/e/sap/nwrfc750P_15-linux/nwrfcsdk && \
	./sap_rfc_proxy