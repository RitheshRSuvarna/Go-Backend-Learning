. (Join-Path $PSScriptRoot "_state.ps1")

$p = Get-AskHubStatePath
if (Test-Path $p) {
  Remove-Item $p -Force
  Write-Host "Deleted $p"
} else {
  Write-Host "No state file found at $p"
}

Write-Host "State reset complete."
