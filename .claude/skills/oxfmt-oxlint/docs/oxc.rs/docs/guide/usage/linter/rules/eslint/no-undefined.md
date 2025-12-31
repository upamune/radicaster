---
title: "eslint/no-undefined | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-undefined"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-undefined.md for this page in Markdown format

# eslint/no-undefined Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-undefined.html#eslint-no-undefined)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-undefined.html#what-it-does)

Disallow the use of `undefined` as an identifier

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-undefined.html#why-is-this-bad)

Using undefined directly can lead to bugs, since it can be shadowed or overwritten in JavaScript. It's safer and more intentional to use null or rely on implicit undefined (e.g., missing return) to avoid accidental issues.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-undefined.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var foo = undefined;

var undefined = "foo";

if (foo === undefined) {
  // ...
}

function baz(undefined) {
  // ...
}

bar(undefined, "lorem");
```

Examples of **correct** code for this rule:

javascript

```
var foo = void 0;

var Undefined = "foo";

if (typeof foo === "undefined") {
  // ...
}

global.undefined = "foo";

bar(void 0, "lorem");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-undefined.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-undefined": "error"
  }
}
```

bash

```
oxlint --deny no-undefined
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-undefined.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_undefined.rs)
