---
title: "eslint/guard-for-in | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/guard-for-in"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/guard-for-in.md for this page in Markdown format

# eslint/guard-for-in Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/guard-for-in.html#eslint-guard-for-in)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/guard-for-in.html#what-it-does)

Require for-in loops to include an if statement.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/guard-for-in.html#why-is-this-bad)

Looping over objects with a `for in` loop will include properties that are inherited through the prototype chain. Using a `for in` loop without filtering the results in the loop can lead to unexpected items in your for loop which can then lead to unexpected behaviour.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/guard-for-in.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
for (key in foo) {
  doSomething(key);
}
```

Examples of **correct** code for this rule:

javascript

```
for (key in foo) {
  if (Object.hasOwn(foo, key)) {
    doSomething(key);
  }
}
```

javascript

```
for (key in foo) {
  if (Object.prototype.hasOwnProperty.call(foo, key)) {
    doSomething(key);
  }
}
```

javascript

```
for (key in foo) {
  if ({}.hasOwnProperty.call(foo, key)) {
    doSomething(key);
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/guard-for-in.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "guard-for-in": "error"
  }
}
```

bash

```
oxlint --deny guard-for-in
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/guard-for-in.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/guard_for_in.rs)
