# weather-app-go
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o weather-app-go

Compress-Archive -Path * -DestinationPath functionapp.zip

az functionapp deployment source config-zip --resource-group weather-project --name weather-go --src functionapp.zip