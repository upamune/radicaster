---
title: "eslint/no-multi-assign | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-multi-assign"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-multi-assign.md for this page in Markdown format

# eslint/no-multi-assign Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-multi-assign.html#eslint-no-multi-assign)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-multi-assign.html#what-it-does)

Disallow use of chained assignment expressions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-multi-assign.html#why-is-this-bad)

Chaining the assignment of variables can lead to unexpected results and be difficult to read.

js

```
(function () {
  const foo = (bar = 0); // Did you mean `foo = bar == 0`?
  bar = 1; // This will not fail since `bar` is not constant.
})();
console.log(bar); // This will output 1 since `bar` is not scoped.
```

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-multi-assign.html#examples)

Examples of **incorrect** code for this rule:

js

```
var a = (b = c = 5);

const foo = (bar = "baz");

let d = (e = f);

class Foo {
  a = (b = 10);
}

a = b = "quux";
```

Examples of **correct** code for this rule:

js

```
var a = 5;
var b = 5;
var c = 5;

const foo = "baz";
const bar = "baz";

let d = c;
let e = c;

class Foo {
  a = 10;
  b = 10;
}

a = "quux";
b = "quux";
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-multi-assign.html#configuration)

This rule accepts a configuration object with the following properties:

### ignoreNonDeclaration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-multi-assign.html#ignorenondeclaration)

type: `boolean`

default: `false`

When set to `true`, the rule allows chains that don't include initializing a variable in a declaration or initializing a class field.

Examples of **correct** code for this option set to `true`:

js

```
let a;
let b;
a = b = "baz";

const x = {};
const y = {};
x.one = y.one = 1;
```

Examples of **incorrect** code for this option set to `true`:

js

```
let a = (b = "baz");

const foo = (bar = 1);

class Foo {
  a = (b = 10);
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-multi-assign.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-multi-assign": "error"
  }
}
```

bash

```
oxlint --deny no-multi-assign
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-multi-assign.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_multi_assign.rs)
