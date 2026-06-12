$ErrorActionPreference = "Stop"

function Get-AskHubStatePath {
  return (Join-Path $PSScriptRoot ".state.json")
}

function Read-AskHubState {
  $p = Get-AskHubStatePath
  if (-not (Test-Path $p)) {
    return @{
      apiBase = "http://localhost:8080"
    }
  }
  return (Get-Content $p -Raw | ConvertFrom-Json -AsHashtable)
}

function Write-AskHubState([hashtable]$state) {
  $p = Get-AskHubStatePath
  ($state | ConvertTo-Json -Depth 10) | Set-Content -Path $p -Encoding UTF8
}

function Require-StateValue([hashtable]$state, [string]$key) {
  if (-not $state.ContainsKey($key) -or [string]::IsNullOrWhiteSpace([string]$state[$key])) {
    throw "Missing state '$key'. Run the earlier script that sets it (see Rithesh-backend-docs/scripts/README.md)."
  }
}

function New-AskHubJsonBody([hashtable]$obj) {
  return ($obj | ConvertTo-Json -Depth 10 -Compress)
}

