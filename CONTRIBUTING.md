# Contributing

Thanks for your interest in improving Backend-Path.

## Project Setup Expectations
- Use the Go version compatible with `go.mod`.
- Sync modules before coding:
  - `go mod download`
- Run checks before creating a PR:
  - `go vet ./...`
  - `go build ./...`
  - `go test ./... -v`

## Branch Naming Suggestion
Use short, descriptive branch names:
- `feat/<scope>`
- `fix/<scope>`
- `refactor/<scope>`
- `chore/<scope>`

## Commit Message Guidance
Prefer clear, scoped commit messages (Conventional Commits style is encouraged):
- `feat: ...`
- `fix: ...`
- `refactor: ...`
- `test: ...`
- `chore: ...`

## Pull Request Expectations
- Keep PRs small and focused.
- Explain what changed and why.
- Mention any risk, migration, or compatibility impact.
- Link related issue(s) when available.

## Test Expectations Before Opening a PR
Run and pass:
- `go vet ./...`
- `go build ./...`
- `go test ./... -v`

If your change targets integration/e2e behavior, include relevant evidence in the PR description.

## Keep PRs Focused
Avoid bundling unrelated refactors with functional fixes. Smaller PRs are easier to review and safer to merge.

## Update Docs When Behavior Changes
If API behavior, architecture assumptions, or developer setup changes, update:
- `README.md`
- `docs/architecture.md`
- any impacted developer docs

## Respect Layering and Structure
Please keep concerns in their intended layers:
- transport in `internal/api`
- use-case orchestration in `internal/application`
- domain models/contracts in `internal/domain`
- concrete integrations in `internal/infrastructure`
