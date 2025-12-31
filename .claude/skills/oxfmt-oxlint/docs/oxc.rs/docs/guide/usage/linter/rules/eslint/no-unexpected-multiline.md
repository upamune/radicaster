---
title: "eslint/no-unexpected-multiline | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unexpected-multiline"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-unexpected-multiline.md for this page in Markdown format

# eslint/no-unexpected-multiline Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unexpected-multiline.html#eslint-no-unexpected-multiline)

⚠️🛠️️ A dangerous auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unexpected-multiline.html#what-it-does)

In most cases, semicolons are not required in JavaScript in order for code to be parsed and executed as expected. Typically this occurs because semicolons are automatically inserted based on a fixed set of rules. This rule exists to detect those cases where a semicolon is NOT inserted automatically, and may be parsed differently than expected.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unexpected-multiline.html#why-is-this-bad)

Code that has unexpected newlines may be parsed and executed differently than what the developer intended. This can lead to bugs that are difficult to track down.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unexpected-multiline.html#examples)

Examples of **incorrect** code for this rule:

js

```
var a = b(x || y).doSomething();

var a = b[(a, b, c)].forEach(doSomething);

let x = (function () {})`hello`;

foo / bar / g.test(baz);
```

Examples of **correct** code for this rule:

js

```
var a = b;
(x || y).doSomething();

var a = b;
[a, b, c].forEach(doSomething);

let x = function () {};
`hello`;

foo;
/bar/g.test(baz);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unexpected-multiline.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-unexpected-multiline": "error"
  }
}
```

bash

```
oxlint --deny no-unexpected-multiline
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unexpected-multiline.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_unexpected_multiline.rs)
