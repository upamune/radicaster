---
title: "typescript/no-implied-eval | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-implied-eval"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-implied-eval.md for this page in Markdown format

# typescript/no-implied-eval Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-implied-eval.html#typescript-no-implied-eval)

✅ This rule is turned on by default when type-aware linting is enabled.

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-implied-eval.html#what-it-does)

This rule disallows the use of eval-like methods.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-implied-eval.html#why-is-this-bad)

It's considered a good practice to avoid using eval() in JavaScript. There are security and performance implications involved with doing so, which is why many linters recommend disallowing eval(). However, there are some other ways to pass a string and have it interpreted as JavaScript code that have similar concerns.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-implied-eval.html#examples)

Examples of **incorrect** code for this rule:

ts

```
setTimeout('alert("Hi!");', 100);

setInterval('alert("Hi!");', 100);

setImmediate('alert("Hi!")');

window.setTimeout("count = 5", 10);

window.setInterval("foo = bar", 10);

const fn = new Function("a", "b", "return a + b");
```

Examples of **correct** code for this rule:

ts

```
setTimeout(() => {
  alert("Hi!");
}, 100);

setInterval(() => {
  alert("Hi!");
}, 100);

setImmediate(() => {
  alert("Hi!");
});

const fn = (a: number, b: number) => a + b;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-implied-eval.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-implied-eval": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-implied-eval
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-implied-eval.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_implied_eval.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_implied_eval/no_implied_eval.go)
