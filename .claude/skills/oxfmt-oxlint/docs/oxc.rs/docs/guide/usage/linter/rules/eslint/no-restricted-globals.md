---
title: "eslint/no-restricted-globals | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-restricted-globals"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-restricted-globals.md for this page in Markdown format

# eslint/no-restricted-globals Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-restricted-globals.html#eslint-no-restricted-globals)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-restricted-globals.html#what-it-does)

This rule allows you to specify global variable names that you don't want to use in your application.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-restricted-globals.html#why-is-this-bad)

Disallowing usage of specific global variables can be useful if you want to allow a set of global variables by enabling an environment, but still want to disallow some of those.

For instance, early Internet Explorer versions exposed the current DOM event as a global variable `event`, but using this variable has been considered as a bad practice for a long time. Restricting this will make sure this variable isn't used in browser code.

### Example [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-restricted-globals.html#example)

If we have options:

json

```
"no-restricted-globals": ["error", "event"]
```

The following patterns are considered problems:

javascript

```
function onClick() {
  console.log(event); // Unexpected global variable 'event'. Use local parameter instead.
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-restricted-globals.html#configuration)

This rule accepts a configuration object with the following properties:

### restrictedGlobals [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-restricted-globals.html#restrictedglobals)

type: `Record<string, string>`

default: `{}`

Objects in the format `{ "name": "event", "message": "Use local parameter instead." }`, which define what globals are restricted from use.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-restricted-globals.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-restricted-globals": "error"
  }
}
```

bash

```
oxlint --deny no-restricted-globals
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-restricted-globals.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_restricted_globals.rs)
