---
title: "jsdoc/check-access | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-access"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsdoc/check-access.md for this page in Markdown format

# jsdoc/check-access Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-access.html#jsdoc-check-access)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-access.html#what-it-does)

Checks that `@access` tags use one of the following values:

* "package", "private", "protected", "public"

Also reports:

* Mixing of `@access` with `@public`, `@private`, `@protected`, or `@package` on the same doc block.
* Use of multiple instances of `@access` (or the `@public`, etc) on the same doc block.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-access.html#why-is-this-bad)

It is important to have a consistent way of specifying access levels in JSDoc comments. Using invalid or multiple access level tags creates confusion about the intended visibility of documented elements and can lead to inconsistencies in API documentation generation. Mixing different access tags or using invalid values makes the documentation unclear and potentially misleading.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-access.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
/** @access private @public */

/** @access invalidlevel */
```

Examples of **correct** code for this rule:

javascript

```
/** @access private */

/** @private */
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-access.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsdoc"],
  "rules": {
    "jsdoc/check-access": "error"
  }
}
```

bash

```
oxlint --deny jsdoc/check-access --jsdoc-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-access.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsdoc/check_access.rs)
