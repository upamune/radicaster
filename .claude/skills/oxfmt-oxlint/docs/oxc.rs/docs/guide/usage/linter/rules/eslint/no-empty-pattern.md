---
title: "eslint/no-empty-pattern | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-pattern"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-empty-pattern.md for this page in Markdown format

# eslint/no-empty-pattern Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-pattern.html#eslint-no-empty-pattern)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-pattern.html#what-it-does)

Disallow empty destructuring patterns

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-pattern.html#why-is-this-bad)

When using destructuring, it’s possible to create a pattern that has no effect. This happens when empty curly braces are used to the right of an embedded object destructuring pattern, such as:

JavaScript

```
// doesn't create any variables
var {a: {}} = foo;
```

In this code, no new variables are created because a is just a location helper while the `{}` is expected to contain the variables to create, such as:

JavaScript

```
// creates variable b
var {a: { b }} = foo;
```

In many cases, the empty object pattern is a mistake where the author intended to use a default value instead, such as:

JavaScript

```
// creates variable a
var {a = {}} = foo;
```

The difference between these two patterns is subtle, especially because the problematic empty pattern looks just like an object literal.

### Examples of incorrect code for this rule: [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-pattern.html#examples-of-incorrect-code-for-this-rule)

JavaScript

```
var {} = foo;
var [] = foo;
var {a: {}} = foo;
var {a: []} = foo;
function foo({}) {}
function foo([]) {}
function foo({a: {}}) {}
function foo({a: []}) {}
```

### Examples of correct code for this rule: [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-pattern.html#examples-of-correct-code-for-this-rule)

JavaScript

```
var {a = {}} = foo;
var {a = []} = foo;
function foo({a = {}}) {}
function foo({a = []}) {}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-pattern.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-empty-pattern": "error"
  }
}
```

bash

```
oxlint --deny no-empty-pattern
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-pattern.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_empty_pattern.rs)
