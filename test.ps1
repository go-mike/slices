[CmdletBinding()]
param(
    [Parameter()][switch] $Cover
)

# https://github.com/wgross/fswatcher-engine-event

$Command = "go test -fuzz=Fuzz -fuzztime=1s ./..."
if ($VerbosePreference) {
    $Command = $Command + " -v"
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
