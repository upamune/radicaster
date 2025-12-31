---
title: "jest/no-test-prefixes | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-test-prefixes"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-test-prefixes.md for this page in Markdown format

# jest/no-test-prefixes Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-test-prefixes.html#jest-no-test-prefixes)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-test-prefixes.html#what-it-does)

Require using `.only` and `.skip` over `f` and `x`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-test-prefixes.html#why-is-this-bad)

Jest allows you to choose how you want to define focused and skipped tests, with multiple permutations for each:

* only & skip: it.only, test.only, describe.only, it.skip, test.skip, describe.skip.
* 'f' & 'x': fit, fdescribe, xit, xtest, xdescribe.

This rule enforces usages from the only & skip list.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-test-prefixes.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
fit("foo"); // invalid
fdescribe("foo"); // invalid
xit("foo"); // invalid
xtest("foo"); // invalid
xdescribe("foo"); // invalid
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/v1.1.9/docs/rules/no-test-prefixes.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/no-test-prefixes": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-test-prefixes.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-test-prefixes": "error"
  }
}
```

bash

```
oxlint --deny jest/no-test-prefixes --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-test-prefixes.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_test_prefixes.rs)
