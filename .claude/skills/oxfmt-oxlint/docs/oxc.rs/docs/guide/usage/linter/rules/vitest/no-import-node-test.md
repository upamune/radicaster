---
title: "vitest/no-import-node-test | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vitest/no-import-node-test"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vitest/no-import-node-test.md for this page in Markdown format

# vitest/no-import-node-test Style [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/no-import-node-test.html#vitest-no-import-node-test)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/no-import-node-test.html#what-it-does)

This rule warns when `node:test` is imported (usually accidentally). With `--fix`, it will replace the import with `vitest`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/no-import-node-test.html#why-is-this-bad)

Using `node:test` instead of `vitest` can lead to inconsistent test results and missing features. `vitest` should be used for all testing to ensure compatibility and access to its full functionality.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/no-import-node-test.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
import { test } from "node:test";
import { expect } from "vitest";

test("foo", () => {
  expect(1).toBe(1);
});
```

Examples of **correct** code for this rule:

javascript

```
import { test, expect } from "vitest";

test("foo", () => {
  expect(1).toBe(1);
});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/no-import-node-test.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vitest"],
  "rules": {
    "vitest/no-import-node-test": "error"
  }
}
```

bash

```
oxlint --deny vitest/no-import-node-test --vitest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/no-import-node-test.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vitest/no_import_node_test.rs)
