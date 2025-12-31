---
title: "unicorn/no-instanceof-builtins | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-builtins"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-instanceof-builtins.md for this page in Markdown format

# unicorn/no-instanceof-builtins Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-builtins.html#unicorn-no-instanceof-builtins)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-builtins.html#what-it-does)

Disallows the use of `instanceof` with ECMAScript built-in constructors because:

* it breaks across execution contexts (`iframe`, Web Worker, Node VM, etc.);
* it is often misleading (e.g. `instanceof Array` fails for a subclass);
* there is always a clearer and safer alternative (`Array.isArray`, `typeof`, `Buffer.isBuffer`, …).

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-builtins.html#why-is-this-bad)

`instanceof` breaks across execution contexts (`iframe`, Web Worker, Node `vm`), and may give misleading results for subclasses or exotic objects.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-builtins.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if (arr instanceof Array) { … }
if (el instanceof HTMLElement) { … }
```

Examples of **correct** code for this rule:

javascript

```
if (Array.isArray(arr)) { … }
if (el?.nodeType === 1) { … }
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-builtins.html#configuration)

This rule accepts a configuration object with the following properties:

### exclude [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-builtins.html#exclude)

type: `string[]`

default: `[]`

Constructor names to exclude from checking.

### include [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-builtins.html#include)

type: `string[]`

default: `[]`

Additional constructor names to check beyond the default set. Use this to extend the rule with additional constructors.

### strategy [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-builtins.html#strategy)

type: `"strict" | "loose"`

default: `"loose"`

Controls which built-in constructors are checked.

* `"loose"` (default): Only checks Array, Function, Error (if `useErrorIsError` is true), and primitive wrappers
* `"strict"`: Additionally checks Error types, collections, typed arrays, and other built-in constructors

### useErrorIsError [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-builtins.html#useerroriserror)

type: `boolean`

default: `false`

When `true`, checks `instanceof Error` and suggests using `Error.isError()` instead. Requires [the `Error.isError()` function](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Error/isError) to be available.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-builtins.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-instanceof-builtins": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-instanceof-builtins
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-builtins.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_instanceof_builtins.rs)
