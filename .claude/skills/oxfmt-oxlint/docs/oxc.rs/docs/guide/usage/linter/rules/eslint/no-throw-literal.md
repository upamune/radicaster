---
title: "eslint/no-throw-literal | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-throw-literal"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-throw-literal.md for this page in Markdown format

# eslint/no-throw-literal Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-throw-literal.html#eslint-no-throw-literal)

💡 A suggestion is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-throw-literal.html#what-it-does)

Disallows throwing literals or non-Error objects as exceptions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-throw-literal.html#why-is-this-bad)

It is considered good practice to only throw the Error object itself or an object using the Error object as base objects for user-defined exceptions. The fundamental benefit of Error objects is that they automatically keep track of where they were built and originated.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-throw-literal.html#examples)

Examples of **incorrect** code for this rule:

js

```
throw "error";

throw 0;

throw undefined;

throw null;

var err = new Error();
throw "an " + err;
// err is recast to a string literal

var err = new Error();
throw `${err}`;
```

Examples of **correct** code for this rule:

js

```
throw new Error();

throw new Error("error");

var e = new Error("error");
throw e;

try {
  throw new Error("error");
} catch (e) {
  throw e;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-throw-literal.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-throw-literal": "error"
  }
}
```

bash

```
oxlint --deny no-throw-literal
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-throw-literal.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_throw_literal.rs)
