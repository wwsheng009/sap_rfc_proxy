export SAPNWRFC_HOME=/root/projects/nwrfcsdk
export LD_LIBRARY_PATH=LD_LIBRARY_PATH:$(SAPNWRFC_HOME)/lib

fmt:
	@echo LD_LIBRARY_PATH=$$LD_LIBRARY_PATH
	@echo SAPNWRFC_HOME=$$SAPNWRFC_HOME
	go fmt ./...

# https://pkg.go.dev/cmd/cgo
# When building, the CGO_CFLAGS, CGO_CPPFLAGS, CGO_CXXFLAGS and CGO_LDFLAGS environment variables are added to the flags derived from these directives. 
# Package-specific flags should be set using the directives, not the environment variables, so that builds work in unmodified environments.

# export SDL_PATH=/home/mark/where/I/installed/sdl
# CGO_CFLAGS="-I$SDL_PATH/include" CGO_LDFLAGS="-L$SDL_PATH/lib" go build hello.go
# LD_LIBRARY_PATH="$SDL_PATH/lib" ./hello

build:fmt
	CGO_LDFLAGS="-L$(SAPNWRFC_HOME)/lib" CGO_CFLAGS="-I$(SAPNWRFC_HOME)/include" go build -o sap_rfc_proxy

run:
	@echo LD_LIBRARY_PATH=$$LD_LIBRARY_PATH
	@echo SAPNWRFC_HOME=$$SAPNWRFC_HOME
	./sap_rfc_proxy