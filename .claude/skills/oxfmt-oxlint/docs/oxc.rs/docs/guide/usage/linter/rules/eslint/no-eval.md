---
title: "eslint/no-eval | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-eval"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-eval.md for this page in Markdown format

# eslint/no-eval Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-eval.html#eslint-no-eval)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-eval.html#what-it-does)

Disallows referencing the `eval` function. This rule is aimed at preventing potentially dangerous, unnecessary, and slow code by disallowing the use of the `eval()` function.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-eval.html#why-is-this-bad)

JavaScript’s `eval()` function is potentially dangerous and is often misused. Using `eval()` on untrusted code can open a program up to several different injection attacks. The use of `eval()` in most contexts can be substituted for a better, safer alternative approach to solving the problem, such as using `JSON.parse()` or `Function` constructors in safer ways.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-eval.html#examples)

Examples of **incorrect** code for this rule:

js

```
const obj = { x: "foo" },
  key = "x",
  value = eval("obj." + key);

(0, eval)("const a = 0");

const foo = eval;
foo("const a = 0");

this.eval("const a = 0");
```

Examples of **correct** code for this rule:

js

```
const obj = { x: "foo" },
  key = "x",
  value = obj[key];

class A {
  foo() {
    this.eval("const a = 0");
  }

  eval() {}

  static {
    this.eval("const a = 0");
  }

  static eval() {}
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-eval.html#configuration)

This rule accepts a configuration object with the following properties:

### allowIndirect [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-eval.html#allowindirect)

type: `boolean`

default: `true`

This `allowIndirect` option allows indirect `eval()` calls.

Indirect calls to `eval`(e.g., `window['eval']`) are less dangerous than direct calls because they cannot dynamically change the scope. Indirect `eval()` calls also typically have less impact on performance compared to direct calls, as they do not invoke JavaScript's scope chain.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-eval.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-eval": "error"
  }
}
```

bash

```
oxlint --deny no-eval
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-eval.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_eval.rs)
