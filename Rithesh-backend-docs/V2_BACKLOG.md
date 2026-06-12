# V2 Backlog (Day 30)

## Priority 1 (must)

- [ ] **Production auth hardening**
  - Secure cookie settings (`Secure`, `SameSite`, domain/path), CSRF mitigation, brute-force protection
  - Add basic rate limiting for auth endpoints
- [ ] **Real email provider for password reset**
  - Integrate an email provider (SES/Mailgun/etc)
  - Store and audit reset requests, add throttling
- [ ] **LLM quality evaluation loop**
  - Define a golden dataset of day-sessions
  - Automated scoring for plan constraints + regression checks

## Priority 2 (should)

- [ ] **Real POI provider integration**
  - Provider adapter(s) + caching strategy
  - Normalized POI model (id, geo, hours, category, etc)
- [ ] **Similar POI search + caching**
  - Given a stop, suggest replacements from provider
  - Cache by geo/category/time window
- [ ] **Better deterministic replan objective**
  - More explicit constraints (hard vs soft)
  - Penalize busy-risk, minimize drift, minimize travel time

## Priority 3 (nice to have)

- [ ] **Metrics + tracing**
  - request latency metrics, error counters
  - optional tracing correlation with request IDs
- [ ] **CI pipeline**
  - run `task build`
  - run unit tests
  - verify sqlc + swagger generation is up-to-date
- [ ] **API quotas**
  - rate limits/quotas for LLM endpoints and demo reset

## Notes per item

For each backlog item document:
- user value
- implementation notes
- risks/unknowns
- acceptance criteria

# V2 Backlog (Fill This)

Use `V2_BACKLOG_TEMPLATE.md` as base and replace this file with your final Day30 backlog.
