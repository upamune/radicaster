---
title: "typescript/no-this-alias | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-this-alias"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-this-alias.md for this page in Markdown format

# typescript/no-this-alias Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-this-alias.html#typescript-no-this-alias)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-this-alias.html#what-it-does)

Disallow aliasing `this`

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-this-alias.html#why-is-this-bad)

Assigning a variable to `this` instead of properly using arrow lambdas may be a symptom of pre-ES2015 practices or not managing scope well.

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-this-alias.html#configuration)

This rule accepts a configuration object with the following properties:

### allowDestructuring [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-this-alias.html#allowdestructuring)

type: `boolean`

default: `true`

Whether to allow destructuring of `this` to local variables.

### allowedNames [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-this-alias.html#allowednames)

type: `string[]`

default: `[]`

An array of variable names that are allowed to alias `this`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-this-alias.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-this-alias": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-this-alias
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-this-alias.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_this_alias.rs)
