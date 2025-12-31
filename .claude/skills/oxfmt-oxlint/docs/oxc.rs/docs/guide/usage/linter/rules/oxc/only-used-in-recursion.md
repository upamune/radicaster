---
title: "oxc/only-used-in-recursion | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/only-used-in-recursion"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/only-used-in-recursion.md for this page in Markdown format

# oxc/only-used-in-recursion Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/only-used-in-recursion.html#oxc-only-used-in-recursion)

✅ This rule is turned on by default.

⚠️🛠️️ A dangerous auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/only-used-in-recursion.html#what-it-does)

Checks for arguments that are only used in recursion with no side-effects.

Inspired by <https://rust-lang.github.io/rust-clippy/master/#/only_used_in_recursion>

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/only-used-in-recursion.html#why-is-this-bad)

Supplying an argument that is only used in recursive calls is likely a mistake.

It increase cognitive complexity and may impact performance.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/only-used-in-recursion.html#examples)

Examples of **incorrect** code for this rule:

ts

```
function test(only_used_in_recursion) {
  return test(only_used_in_recursion);
}
```

Examples of **correct** code for this rule:

ts

```
function f(a: number): number {
  if (a == 0) {
    return 1;
  } else {
    return f(a - 1);
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/only-used-in-recursion.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/only-used-in-recursion": "error"
  }
}
```

bash

```
oxlint --deny oxc/only-used-in-recursion
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/only-used-in-recursion.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/only_used_in_recursion.rs)
