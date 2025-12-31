---
title: "eslint/vars-on-top | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/vars-on-top"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/vars-on-top.md for this page in Markdown format

# eslint/vars-on-top Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/vars-on-top.html#eslint-vars-on-top)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/vars-on-top.html#what-it-does)

Enforces that all `var` declarations are placed at the top of their containing scope.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/vars-on-top.html#why-is-this-bad)

In JavaScript, `var` declarations are hoisted to the top of their containing scope. Placing `var` declarations at the top explicitly improves code readability and maintainability by making the scope of variables clear.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/vars-on-top.html#examples)

Examples of **incorrect** code for this rule:

js

```
function doSomething() {
  if (true) {
    var first = true;
  }
  var second;
}

function doSomethingElse() {
  for (var i = 0; i < 10; i++) {}
}

f();
var a;

class C {
  static {
    if (something) {
      var a = true;
    }
  }
  static {
    f();
    var a;
  }
}
```

Examples of **correct** code for this rule:

js

```
function doSomething() {
  var first;
  var second;
  if (true) {
    first = true;
  }
}

function doSomethingElse() {
  var i;
  for (i = 0; i < 10; i++) {}
}

var a;
f();

class C {
  static {
    var a;
    if (something) {
      a = true;
    }
  }
  static {
    var a;
    f();
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/vars-on-top.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "vars-on-top": "error"
  }
}
```

bash

```
oxlint --deny vars-on-top
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/vars-on-top.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/vars_on_top.rs)
