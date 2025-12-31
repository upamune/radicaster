---
title: "eslint/valid-typeof | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/valid-typeof"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/valid-typeof.md for this page in Markdown format

# eslint/valid-typeof Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/valid-typeof.html#eslint-valid-typeof)

✅ This rule is turned on by default.

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/valid-typeof.html#what-it-does)

Enforce comparing `typeof` expressions against valid strings.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/valid-typeof.html#why-is-this-bad)

For a vast majority of use cases, the result of the `typeof` operator is one of the following string literals: `"undefined"`, `"object"`, `"boolean"`, `"number"`, `"string"`, `"function"`, `"symbol"`, and `"bigint"`. It is usually a typing mistake to compare the result of a `typeof` operator to other string literals.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/valid-typeof.html#examples)

Examples of **incorrect** code for this rule:

js

```
typeof foo === "strnig";
typeof foo == "undefimed";
typeof bar != "nunber"; // spellchecker:disable-line
typeof bar !== "fucntion"; // spellchecker:disable-line
```

Examples of **correct** code for this rule:

js

```
typeof foo === "string";
typeof bar == "undefined";
typeof foo === baz;
typeof bar === typeof qux;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/valid-typeof.html#configuration)

This rule accepts a configuration object with the following properties:

### requireStringLiterals [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/valid-typeof.html#requirestringliterals)

type: `boolean`

default: `false`

The `requireStringLiterals` option when set to `true`, allows the comparison of `typeof` expressions with only string literals or other `typeof` expressions, and disallows comparisons to any other value. Default is `false`.

With `requireStringLiterals` set to `true`, the following are examples of **incorrect** code:

js

```
typeof foo === undefined;
typeof bar == Object;
typeof baz === "strnig";
typeof qux === "some invalid type";
typeof baz === anotherVariable;
typeof foo == 5;
```

With `requireStringLiterals` set to `true`, the following are examples of **correct** code:

js

```
typeof foo === "undefined";
typeof bar == "object";
typeof baz === "string";
typeof bar === typeof qux;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/valid-typeof.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "valid-typeof": "error"
  }
}
```

bash

```
oxlint --deny valid-typeof
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/valid-typeof.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/valid_typeof.rs)
