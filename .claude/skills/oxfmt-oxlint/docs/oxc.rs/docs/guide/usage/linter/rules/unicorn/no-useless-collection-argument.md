---
title: "unicorn/no-useless-collection-argument | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-collection-argument"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-useless-collection-argument.md for this page in Markdown format

# unicorn/no-useless-collection-argument Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-collection-argument.html#unicorn-no-useless-collection-argument)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-collection-argument.html#what-it-does)

Disallow useless values or fallbacks in Set, Map, WeakSet, or WeakMap

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-collection-argument.html#why-is-this-bad)

It's unnecessary to pass an empty array or string when constructing a Set, Map, WeakSet, or WeakMap, since they accept nullish values. It's also unnecessary to provide a fallback for possible nullish values.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-collection-argument.html#examples)

Examples of **incorrect** code for this rule:

js

```
const set = new Set([]);
const set = new Set("");
```

Examples of **correct** code for this rule:

js

```
const set = new Set();
```

Examples of **incorrect** code for this rule:

js

```
const set = new Set(foo ?? []);
const set = new Set(foo ?? "");
```

Examples of **correct** code for this rule:

js

```
const set = new Set(foo);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-collection-argument.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-useless-collection-argument": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-useless-collection-argument
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-collection-argument.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_useless_collection_argument.rs)
