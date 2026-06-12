# Day -1 Setup and Baseline

Complete this before Day 00.

## Goal

Your machine is ready, and you have a clean workspace where this curriculum can run standalone.

## Install

- Git
- Go (1.22+ recommended)
- Docker Desktop
- Python 3.11+

## Verify install

Run in PowerShell:

```powershell
git --version
go version
docker --version
python --version
```

## Prepare workspace (standalone mode)

Create a new folder anywhere (example only):

```powershell
mkdir backend-learning
cd backend-learning
```

Place `Rithesh-backend-docs` inside this folder.

Expected layout now:

```text
backend-learning/
  Rithesh-backend-docs/
```

## Run preflight

From workspace root:

```powershell
. .\Rithesh-backend-docs\scripts\00_preflight.ps1
```

Note:
- Before Day00, `services/api` will not exist yet. That is expected.

## Optional tools (recommended)

```powershell
go install github.com/swaggo/swag/cmd/swag@latest
```

If `swag` does not work, add Go bin to PATH:
- `%USERPROFILE%\go\bin`

## Acceptance checks

- all required tools are installed
- preflight script runs
- workspace has `Rithesh-backend-docs` folder

## Deliverable

- machine + workspace ready for Day00 coding.
