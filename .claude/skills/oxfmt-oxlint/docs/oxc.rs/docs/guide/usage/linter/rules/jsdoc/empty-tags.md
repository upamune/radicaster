---
title: "jsdoc/empty-tags | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/empty-tags"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsdoc/empty-tags.md for this page in Markdown format

# jsdoc/empty-tags Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/empty-tags.html#jsdoc-empty-tags)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/empty-tags.html#what-it-does)

Expects the following tags to be empty of any content:

* `@abstract`
* `@async`
* `@generator`
* `@global`
* `@hideconstructor`
* `@ignore`
* `@inner`
* `@instance`
* `@override`
* `@readonly`
* `@inheritDoc`
* `@internal`
* `@overload`
* `@package`
* `@private`
* `@protected`
* `@public`
* `@static`

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/empty-tags.html#why-is-this-bad)

The void tags should be empty.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/empty-tags.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
/** @async foo */

/** @private bar */
```

Examples of **correct** code for this rule:

javascript

```
/** @async */

/** @private */
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/empty-tags.html#configuration)

This rule accepts a configuration object with the following properties:

### tags [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/empty-tags.html#tags)

type: `string[]`

default: `[]`

Additional tags to check for their descriptions.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/empty-tags.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsdoc"],
  "rules": {
    "jsdoc/empty-tags": "error"
  }
}
```

bash

```
oxlint --deny jsdoc/empty-tags --jsdoc-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/empty-tags.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsdoc/empty_tags.rs)
