# import
$utils = (Resolve-Path -Path ".\utils.ps1").ProviderPath
. $utils


<#
    Writing FileA and FileB is asynchronous, their relative order doesn’t matter
    for correctness.
    They can be executed serially, or at the same time, in any order and all
    outcomes are still correct.
    -- Asynchrony does not require concurrency.
#>

# serial
Write-File -Path "./FileA"
Write-File -Path "./FileB"

# parallel
@(
    "./FileA",
    "./FileB",
    "./FileA",
    "./FileB"
) | Foreach-Object -Parallel {
    . $using:utils
    Write-File -Path $PSItem
}