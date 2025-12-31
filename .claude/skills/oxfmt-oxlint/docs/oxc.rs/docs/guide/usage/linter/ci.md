---
title: "CI setup | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/ci"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/ci.md for this page in Markdown format

# Setup CI [​](https://oxc.rs/docs/guide/usage/linter/ci.html#setup-ci)

## GitHub Actions [​](https://oxc.rs/docs/guide/usage/linter/ci.html#github-actions)

Create `.github/workflows/oxlint.yml`:

yaml

```
name: Lint

on:
  pull_request:
  push:
    branches: [main]

jobs:
  oxlint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-node@v4
        with:
          node-version: lts/*
          cache: pnpm

      - uses: pnpm/action-setup@v4

      - run: pnpm install --frozen-lockfile
      - run: pnpm run lint --deny-warnings # given package.json scripts "lint": "oxlint"
```

### pre-commit [​](https://oxc.rs/docs/guide/usage/linter/ci.html#pre-commit)

.pre-commit-hooks.yaml

yaml

```
repos:
  - repo: https://github.com/oxc-project/mirrors-oxlint
    rev: v0.0.0 # change to the latest version
    hooks:
      - id: oxlint
        verbose: true
```

### Unplugin [​](https://oxc.rs/docs/guide/usage/linter/ci.html#unplugin)

<https://www.npmjs.com/package/unplugin-oxlint>

### Vite plugin [​](https://oxc.rs/docs/guide/usage/linter/ci.html#vite-plugin)

<https://www.npmjs.com/package/vite-plugin-oxlint>
