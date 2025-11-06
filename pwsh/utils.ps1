Function Write-File {
    param (
        [string]$Path,
        [string]$Value = (Get-Date -Format 'yyyy-MM-dd HH:mm:ss.fff')
    )
    $Value | Out-File -FilePath $path -Append
}