# Day 27 - Similar POIs + replace stop (backend-only v1)

## Goal (end of day)
- replace-stop endpoint exists
- replacement creates new plan version and records `manual_edit` event

## Learn (20-40 minutes)
- REST endpoint design for domain actions

## Tasks
### 1) Add endpoint
- `POST /day-sessions/{id}/replace-stop`

Body (v1):
- `old_stop_id`
- `new_title`
- `new_category_label`

### 2) Validate input
- old stop exists in active plan
- new title/category are not empty

### 3) Apply replacement
- create new plan version with replacement
- preserve order except replaced stop fields
- reason code includes `user_manual_replace`

### 4) Update active plan
Set created version as active.

### 5) Record event
Insert event with type `manual_edit` containing useful metadata.

## Acceptance checks
- response returns new plan version id
- active plan now points to new version
- events endpoint shows `manual_edit`

## Deliverables
- user-driven plan edit flow implemented.
