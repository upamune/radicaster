---
title: "eslint/no-caller | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-caller"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-caller.md for this page in Markdown format

# eslint/no-caller Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-caller.html#eslint-no-caller)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-caller.html#what-it-does)

Disallow the use of `arguments.caller` or `arguments.callee`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-caller.html#why-is-this-bad)

The use of `arguments.caller` and `arguments.callee` make several code optimizations impossible. They have been deprecated in future versions of JavaScript and their use is forbidden in ECMAScript 5 while in strict mode.

js

```
function foo() {
  var callee = arguments.callee;
}
```

This rule is aimed at discouraging the use of deprecated and sub-optimal code by disallowing the use of `arguments.caller` and `arguments.callee`. As such, it will warn when `arguments.caller` and `arguments.callee` are used.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-caller.html#examples)

Examples of **incorrect** code for this rule:

js

```
function foo(n) {
  if (n <= 0) {
    return;
  }

  arguments.callee(n - 1);
}

[1, 2, 3, 4, 5].map(function (n) {
  return !(n > 1) ? 1 : arguments.callee(n - 1) * n;
});
```

Examples of **correct** code for this rule:

js

```
function foo(n) {
  if (n <= 0) {
    return;
  }

  foo(n - 1);
}

[1, 2, 3, 4, 5].map(function factorial(n) {
  return !(n > 1) ? 1 : factorial(n - 1) * n;
});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-caller.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-caller": "error"
  }
}
```

bash

```
oxlint --deny no-caller
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-caller.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_caller.rs)
