# ci-cd-reference

A small, working reference for CI/CD on a multi-module Go monorepo. The services
are deliberately trivial; the pipeline around them is the point.

```
modules/api      HTTP service + tests + Dockerfile
modules/worker   job processor + tests + Dockerfile
.github/workflows/ci.yml           the pipeline
.github/workflows/ci-watchdog.yml  alarm for "main is red or untested"
```

## Pipeline

```
preflight ─┬─ lint (matrix) ─┐
           └─ test (matrix) ─┴─ build (matrix) ─┐
secrets ────────────────────────────────────────┴─ ci (single gate)
```

| Stage | What it does | Why |
|---|---|---|
| `preflight` | Compares the installed Go toolchain with the `go.work` directive and fails if it is older. | Without it a stale toolchain fails later with a confusing build error, or CI stays red for days unnoticed. |
| `secrets` | gitleaks over full history. | A leaked token is cheaper to catch before merge than to rotate after. |
| `lint` / `test` | `golangci-lint` and `go test -race`, one matrix leg per module. | `fail-fast: false` so one broken module does not hide the state of the others. |
| `build` | Builds each image with buildx (no push) and generates an SPDX SBOM. | Proves the image builds and leaves an inventory artifact for every commit. |
| `ci` | One job that depends on everything. | Make *this* the only required check. Adding a module never means editing branch protection. |

`concurrency.cancel-in-progress` keeps a fast-moving `main` from queueing stale
runs. That choice has a cost, which the watchdog covers.

## The watchdog

A green *last* run does not prove `main` is tested. A push that races an active
run can be cancelled, leaving the newest commit with no run at all.
`ci-watchdog.yml` runs every 30 minutes and checks two things:

1. the last finished (non-cancelled) run on `main` succeeded, and
2. the head SHA of `main` has at least one run.

If either fails it opens one tracking issue, comments while the problem
persists, and closes the issue on recovery.

## Running in a restricted network

If your runners cannot reach the public internet (or reach it unreliably), keep
the pipeline unchanged and move the dependencies inside:

- Mirror the third-party actions as repositories on your own forge, with tag sync enabled.
- Seed the runner's shared toolcache with the Go toolchain and the binaries the
  pipeline shells out to (SBOM generator, scanner), instead of downloading them at job time.
- Mirror base images into your own registry and pass the reference as a build
  `ARG` declared *before* the first `FROM`, so the Dockerfile stays portable.
- Keep the `preflight` stage. When a toolchain download silently fails, it is
  what turns "red for days" into a clear failure on the first run.

## Local

```bash
make lint test build
```

Requires Go, `golangci-lint`, and a Docker daemon for the image builds.

## License

MIT
