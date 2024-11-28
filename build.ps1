# Set-Item -Path Env:CGO_LDFLAGS -Value '-L E:/sap/nwrfc750P_15-70002755/nwrfcsdk/lib -lsapnwrfc -llibsapucum -O2 -g -pthread -pie -fPIE -OPT:REF -LTCG'
Set-Item -Path Env:CGO_LDFLAGS -Value '-L E:/sap/nwrfc750P_15-70002755/nwrfcsdk/lib'
Set-Item -Path Env:CGO_CPPFLAGS -Value '-I E:/sap/nwrfc750P_15-70002755/nwrfcsdk/include'

echo "build yao"

go build -v -o proxy.exe