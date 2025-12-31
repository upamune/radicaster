---
title: "vitest/require-local-test-context-for-concurrent-snapshots | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vitest/require-local-test-context-for-concurrent-snapshots"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vitest/require-local-test-context-for-concurrent-snapshots.md for this page in Markdown format

# vitest/require-local-test-context-for-concurrent-snapshots Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/require-local-test-context-for-concurrent-snapshots.html#vitest-require-local-test-context-for-concurrent-snapshots)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/require-local-test-context-for-concurrent-snapshots.html#what-it-does)

The rule is intended to ensure that concurrent snapshot tests are executed within a properly configured local test context.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/require-local-test-context-for-concurrent-snapshots.html#why-is-this-bad)

Running snapshot tests concurrently without a proper context can lead to unreliable or inconsistent snapshots. Ensuring that concurrent tests are correctly configured with the appropriate context helps maintain accurate and stable snapshots, avoiding potential conflicts or failures.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/require-local-test-context-for-concurrent-snapshots.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
test.concurrent("myLogic", () => {
  expect(true).toMatchSnapshot();
});

describe.concurrent("something", () => {
  test("myLogic", () => {
    expect(true).toMatchInlineSnapshot();
  });
});
```

Examples of **correct** code for this rule:

javascript

```
test.concurrent("myLogic", ({ expect }) => {
  expect(true).toMatchSnapshot();
});

test.concurrent("myLogic", (context) => {
  context.expect(true).toMatchSnapshot();
});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/require-local-test-context-for-concurrent-snapshots.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vitest"],
  "rules": {
    "vitest/require-local-test-context-for-concurrent-snapshots": "error"
  }
}
```

bash

```
oxlint --deny vitest/require-local-test-context-for-concurrent-snapshots --vitest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/require-local-test-context-for-concurrent-snapshots.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vitest/require_local_test_context_for_concurrent_snapshots.rs)
