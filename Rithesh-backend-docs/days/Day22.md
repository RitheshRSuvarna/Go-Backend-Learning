# Day 22 - Better replan heuristics (non-LLM)

## Goal (end of day)
- deterministic replan logic is more realistic
- API returns debug metadata explaining decisions

## Learn (20-45 minutes)
- Go time package: `https://pkg.go.dev/time`
- Rule-engine style design basics

## Tasks
### 1) Extend deterministic rules
Implement and keep deterministic:
- if drift >= 20 minutes: reduce remaining `stay_minutes` (minimum 30)
- if next stop has high busy risk soon: move earlier when valid

### 2) Return `replan_debug` in response (dev-only)
Suggested structure:
- fired rules
- moved stops
- before/after positions

### 3) Keep reason codes
Ensure new plan version includes reason codes array.

### 4) Manual verification
Run same request twice and compare outputs.

## Acceptance checks
- same input state gives same output
- replan response explains changes clearly
- no random or hidden behavior

## Deliverables
- stronger deterministic backend replan.
