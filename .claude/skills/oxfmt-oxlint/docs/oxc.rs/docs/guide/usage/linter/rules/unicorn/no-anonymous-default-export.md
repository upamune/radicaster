---
title: "unicorn/no-anonymous-default-export | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-anonymous-default-export"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-anonymous-default-export.md for this page in Markdown format

# unicorn/no-anonymous-default-export Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-anonymous-default-export.html#unicorn-no-anonymous-default-export)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-anonymous-default-export.html#what-it-does)

Disallows anonymous functions and classes as default exports.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-anonymous-default-export.html#why-is-this-bad)

Naming default exports improves searchability and ensures consistent identifiers for a module’s default export in both declaration and import.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-anonymous-default-export.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
export default class {}
export default function () {}
export default () => {};
module.exports = class {};
module.exports = function () {};
module.exports = () => {};
```

Examples of **correct** code for this rule:

javascript

```
export default class Foo {}
export default function foo () {}

const foo = () => {};
export default foo;

module.exports = class Foo {};
module.exports = function foo () {};

const foo = () => {};
module.exports = foo;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-anonymous-default-export.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-anonymous-default-export": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-anonymous-default-export
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-anonymous-default-export.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_anonymous_default_export.rs)
