# Set the environment variable for SAP NW RFC SDK
Set-Item -Path Env:SAP_NWRFC_SDK_PATH -Value "E:/sap/nwrfc750P_15-70002755/nwrfcsdk"

# Print the build message
echo "build yao"

# works on windows cgo build

# Build the Go program with the updated environment variables
$env:CGO_LDFLAGS = "-L $env:SAP_NWRFC_SDK_PATH/lib"
$env:CGO_CFLAGS = "-I $env:SAP_NWRFC_SDK_PATH/include"
$env:CGO_CPPFLAGS = "-I $env:SAP_NWRFC_SDK_PATH/include"

# Run the Go build command
go build -v -o proxy.exe