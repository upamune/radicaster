---
title: "eslint/block-scoped-var | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/block-scoped-var"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/block-scoped-var.md for this page in Markdown format

# eslint/block-scoped-var Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/block-scoped-var.html#eslint-block-scoped-var)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/block-scoped-var.html#what-it-does)

Enforces that variables are both **declared** and **used** within the same block scope. This rule prevents accidental use of variables outside their intended block, mimicking C-style block scoping in JavaScript.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/block-scoped-var.html#why-is-this-bad)

JavaScript’s `var` declarations are hoisted to the top of their enclosing function, which can cause variables declared in a block (e.g., inside an `if` or `for`) to be accessible outside of it. This can lead to hard-to-find bugs. By enforcing block scoping, this rule helps avoid hoisting issues and aligns more closely with how other languages treat block variables.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/block-scoped-var.html#examples)

Examples of **incorrect** code for this rule:

js

```
/* block-scoped-var: "error" */

function doIf() {
  if (true) {
    var build = true;
  }
  console.log(build);
}

function doLoop() {
  for (var i = 0; i < 10; i++) {
    // do something
  }
  console.log(i); // i is accessible here
}

function doSomething() {
  if (true) {
    var foo = 1;
  }
  if (false) {
    foo = 2;
  }
}

function doTry() {
  try {
    var foo = 1;
  } catch (e) {
    console.log(foo);
  }
}
```

Examples of **correct** code for this rule:

js

```
/* block-scoped-var: "error" */

function doIf() {
  var build;
  if (true) {
    build = true;
  }
  console.log(build);
}

function doLoop() {
  var i;
  for (i = 0; i < 10; i++) {
    // do something
  }
  console.log(i);
}

function doSomething() {
  var foo;
  if (true) {
    foo = 1;
  }
  if (false) {
    foo = 2;
  }
}

function doTry() {
  var foo;
  try {
    foo = 1;
  } catch (e) {
    console.log(foo);
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/block-scoped-var.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "block-scoped-var": "error"
  }
}
```

bash

```
oxlint --deny block-scoped-var
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/block-scoped-var.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/block_scoped_var.rs)
