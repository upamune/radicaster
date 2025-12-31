---
title: "jsdoc/check-tag-names | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-tag-names"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsdoc/check-tag-names.md for this page in Markdown format

# jsdoc/check-tag-names Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-tag-names.html#jsdoc-check-tag-names)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-tag-names.html#what-it-does)

Reports invalid block tag names. Additionally checks for tag names that are redundant when using a type checker such as TypeScript.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-tag-names.html#why-is-this-bad)

Using invalid tags can lead to confusion and make the documentation harder to read.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-tag-names.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
/** @Param */
/** @foo */

/**
 * This is redundant when typed.
 * @type {string}
 */
```

Examples of **correct** code for this rule:

javascript

```
/** @param */
```

### Settings [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-tag-names.html#settings)

Configuration for allowed tags is done via [`settings.jsdoc.tagNamePreference`](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-jsdoc-tagnamepreference). There is no CLI-only parameter for this rule.

You can add custom tags by adding a key-value pair where both match the name of the tag you want to add, like so:

Config (.oxlintrc.json)

json

```
{
  "plugins": ["jsdoc"],
  "rules": {
    "jsdoc/check-tag-names": "error"
  },
  "settings": {
    "jsdoc": {
      "tagNamePreference": {
        "customTagName": "customTagName"
      }
    }
  }
}
```

Examples of correct code for this rule with the above configuration, adding the `customTagName` tag:

js

```
/**
 * @customTagName
 */
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-tag-names.html#configuration)

This rule accepts a configuration object with the following properties:

### definedTags [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-tag-names.html#definedtags)

type: `string[]`

default: `[]`

Additional tag names to allow.

### jsxTags [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-tag-names.html#jsxtags)

type: `boolean`

default: `false`

Whether to allow JSX-related tags:

* `jsx`
* `jsxFrag`
* `jsxImportSource`
* `jsxRuntime`

### typed [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-tag-names.html#typed)

type: `boolean`

default: `false`

If typed is `true`, disallow tags that are unnecessary/duplicative of TypeScript functionality.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-tag-names.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsdoc"],
  "rules": {
    "jsdoc/check-tag-names": "error"
  }
}
```

bash

```
oxlint --deny jsdoc/check-tag-names --jsdoc-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-tag-names.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsdoc/check_tag_names.rs)
