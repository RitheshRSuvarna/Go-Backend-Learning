$ErrorActionPreference = "Stop"

Write-Host "== AskHub Backend Preflight =="

function Test-Cmd([string]$name) {
  $cmd = Get-Command $name -ErrorAction SilentlyContinue
  if ($null -eq $cmd) {
    Write-Host "MISSING: $name" -ForegroundColor Red
    return $false
  }
  Write-Host "OK: $name -> $($cmd.Source)"
  return $true
}

$ok = $true
$ok = (Test-Cmd "git") -and $ok
$ok = (Test-Cmd "go") -and $ok
$ok = (Test-Cmd "docker") -and $ok
$ok = (Test-Cmd "python") -and $ok

if (-not $ok) {
  throw "Install missing tools, then rerun this script."
}

Write-Host ""
Write-Host "Versions:"
git --version
go version
docker --version
python --version

Write-Host ""
Write-Host "Workspace checks:"
if (-not (Test-Path "Rithesh-backend-docs")) {
  throw "Missing Rithesh-backend-docs folder. Run this from workspace root."
}
Write-Host "OK: Rithesh-backend-docs is present."

if (Test-Path "services/api") {
  Write-Host "OK: services/api exists."
} else {
  Write-Host "INFO: services/api not found yet (expected before Day00/Day01)."
}

if ((Test-Path "docker-compose.yml") -or (Test-Path "infra/docker-compose.yml")) {
  Write-Host "OK: docker compose file found."
} else {
  Write-Host "INFO: compose file not found yet (expected before Day01)."
}

Write-Host ""
Write-Host "Next steps:"
Write-Host "1) Complete Day00 to create API skeleton"
Write-Host "2) Complete Day01 to add Postgres + compose"
Write-Host "3) Then run API and /health"
