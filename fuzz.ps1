[CmdletBinding()]
param(
    [Parameter()][string] $Fuzz = "Fuzz",
    [Parameter()][int] $FuzzTime = 1,
    [Parameter()][switch] $ExitOnError
)

# https://github.com/wgross/fswatcher-engine-event

$GoTestFiles = Get-ChildItem -Path . -Filter "*_test.go" -Recurse
$HasError = $false
foreach ($GoTestFile in $GoTestFiles) {
    $FileContent = Get-Content -Path $GoTestFile.FullName
    foreach ($Line in $FileContent) {
        if ($Line -match "func ($Fuzz[\w_]+)\(") {
            $FuzzFunction = $Matches[1]
            $Command = "go test . -fuzz=^$($FuzzFunction)`$ -fuzztime=$($FuzzTime)s"
            $PackageList = go list ./...
            $Command = $Command + " -covermode=count `"$PackageList`""
            if ($VerbosePreference) {
                $Command = $Command + " -v"
            }
            Write-Host "✅ $Command" -ForegroundColor Green
            Invoke-Expression "$Command -run=NopeNopeNope"
            if ($LastExitCode -ne 0) {
                Write-Host "❌ FAILED $Command" -ForegroundColor Red
                if ($ExitOnError) {
                    exit 1
                }
                $HasError = $true
            }
        }
    }
}

if ($HasError) {
    exit 1
}
