[CmdletBinding()]
param(
    [Parameter()][switch] $Cover,
    [Parameter()][switch] $Fuzz,
    [Parameter()][string] $FuzzTime = "1s"
)

# https://github.com/wgross/fswatcher-engine-event

$Command = "go test ./..."
if ($VerbosePreference) {
    $Command = $Command + " -v"
}
if ($Fuzz) {
    $Command = $Command + " -fuzz=Fuzz -fuzztime=$FuzzTime"
}
if ($Cover) {
    $PackageList = go list ./...
    $Command = $Command + " -covermode=count -coverprofile coverage.out `"$PackageList`""
} else {
    $Command = $Command + " -covermode=count"
}
Write-Host $Command -ForegroundColor Green
Invoke-Expression $Command

if ($Cover) {
    $Command = "go tool cover '-html=coverage.out'"
    Write-Host $Command -ForegroundColor Green
    Invoke-Expression $Command
}
