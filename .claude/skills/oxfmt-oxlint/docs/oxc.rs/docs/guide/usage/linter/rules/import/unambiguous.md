---
title: "import/unambiguous | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/unambiguous"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/unambiguous.md for this page in Markdown format

# import/unambiguous Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/import/unambiguous.html#import-unambiguous)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/unambiguous.html#what-it-does)

Warn if a `module` could be mistakenly parsed as a `script` and not pure ESM module

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/unambiguous.html#why-is-this-bad)

For ESM-only environments helps to determine files that not pure ESM modules

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/unambiguous.html#examples)

Examples of **incorrect** code for this rule:

js

```
function x() {}

(function x() {
  return 42;
})();
```

Examples of **correct** code for this rule:

js

```
import "foo";
function x() {
  return 42;
}

export function x() {
  return 42;
}

(function x() {
  return 42;
})();
export {}; // simple way to mark side-effects-only file as 'module' without any imports/exports
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/unambiguous.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/unambiguous": "error"
  }
}
```

bash

```
oxlint --deny import/unambiguous --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/unambiguous.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/unambiguous.rs)
