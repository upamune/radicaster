---
title: "unicorn/prefer-object-from-entries | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-object-from-entries"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-object-from-entries.md for this page in Markdown format

# unicorn/prefer-object-from-entries Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-object-from-entries.html#unicorn-prefer-object-from-entries)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-object-from-entries.html#what-it-does)

Encourages using `Object.fromEntries` when converting an array of key-value pairs into an object.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-object-from-entries.html#why-is-this-bad)

Manually constructing objects from key-value pairs using `reduce` or `forEach` is more verbose, error-prone, and harder to understand. The `Object.fromEntries` method is clearer, more declarative, and built for exactly this purpose.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-object-from-entries.html#examples)

Examples of **incorrect** code for this rule:

js

```
const result = pairs.reduce((obj, [key, value]) => {
  obj[key] = value;
  return obj;
}, {});

const result = {};
pairs.forEach(([key, value]) => {
  result[key] = value;
});
```

Examples of **correct** code for this rule:

js

```
const result = Object.fromEntries(pairs);
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-object-from-entries.html#configuration)

This rule accepts a configuration object with the following properties:

### functions [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-object-from-entries.html#functions)

type: `string[]`

default: `["_.fromPairs", "lodash.fromPairs"]`

Additional functions to treat as equivalents to `Object.fromEntries`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-object-from-entries.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-object-from-entries": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-object-from-entries
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-object-from-entries.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_object_from_entries.rs)
