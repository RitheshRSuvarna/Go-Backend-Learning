# Start Here (Read This First)

This curriculum is for Rithesh R Suvarna, who is new to backend work.

You will build a real backend project in small daily steps.

Important: this package is standalone. You only need this folder (`Rithesh-backend-docs`) to start.

You will create and grow this structure over time:
- Go API (`services/api`)
- Postgres
- Python LLM microservice (`services/llm`)

Suggested workspace layout:

```text
backend-learning/
  Rithesh-backend-docs/
  services/          (created during the plan)
  infra/             (created during the plan)
```

## What success looks like

By Day 30 you should be able to:
- run DB + API + LLM locally
- call the full backend flow end-to-end
- explain design decisions and tradeoffs
- hand over docs so another dev can run everything

## Required workflow every day

1. Read the day file fully before coding.
2. Do the Learn links (at least 20 minutes).
3. Implement the tasks in code.
4. Run the acceptance checks for that day.
5. Write a short log:
   - what you changed
   - what passed
   - what is still unclear

## Before Day 00

Complete:
- [DAY_MINUS_1_SETUP.md](./DAY_MINUS_1_SETUP.md)

## Weekly control points

Use:
- [WEEKLY_CHECKPOINTS.md](./WEEKLY_CHECKPOINTS.md)

If a checkpoint fails, do not continue to next week. Fix first.

## If you get blocked

Use:
- [DEBUG_PLAYBOOK.md](./DEBUG_PLAYBOOK.md)

Rule:
- If blocked for more than 45 minutes, write:
  - exact command you ran
  - exact error
  - what you already tried

This is required for fast help from mentor/team.
