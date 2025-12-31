---
title: "react/no-unknown-property | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/no-unknown-property"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/no-unknown-property.md for this page in Markdown format

# react/no-unknown-property Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unknown-property.html#react-no-unknown-property)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unknown-property.html#what-it-does)

Disallow usage of unknown DOM properties.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unknown-property.html#why-is-this-bad)

You can use unknown property name that has no effect.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unknown-property.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
// Unknown properties
const Hello = <div class="hello">Hello World</div>;
const Alphabet = <div abc="something">Alphabet</div>;

// Invalid aria-* attribute
const IconButton = <div aria-foo="bar" />;
```

Examples of **correct** code for this rule:

jsx

```
// Unknown properties
const Hello = <div className="hello">Hello World</div>;
const Alphabet = <div>Alphabet</div>;

// Invalid aria-* attribute
const IconButton = <div aria-label="bar" />;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unknown-property.html#configuration)

This rule accepts a configuration object with the following properties:

### ignore [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unknown-property.html#ignore)

type: `string[]`

default: `[]`

List of properties to ignore.

### requireDataLowercase [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unknown-property.html#requiredatalowercase)

type: `boolean`

default: `false`

Require `data-*` attributes to be lowercase, e.g. `data-foobar` instead of `data-fooBar`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unknown-property.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/no-unknown-property": "error"
  }
}
```

bash

```
oxlint --deny react/no-unknown-property --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unknown-property.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/no_unknown_property.rs)
