---
title: "jsx_a11y/alt-text | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/alt-text"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/alt-text.md for this page in Markdown format

# jsx\_a11y/alt-text Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/alt-text.html#jsx-a11y-alt-text)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/alt-text.html#what-it-does)

Enforce that all elements that require alternative text have meaningful information to relay back to the end user.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/alt-text.html#why-is-this-bad)

Alternative text is a critical component of accessibility for screen reader users, enabling them to understand the content and function of an element. Missing or inadequate alt text makes content inaccessible to users who rely on assistive technologies.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/alt-text.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<img src="flower.jpg" />
<img src="flower.jpg" alt="" />
<object />
<area />
```

Examples of **correct** code for this rule:

jsx

```
<img src="flower.jpg" alt="A close-up of a white daisy" />
<img src="decorative.jpg" alt="" role="presentation" />
<object aria-label="Interactive chart" />
<area alt="Navigation link" />
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/alt-text.html#configuration)

This rule accepts a configuration object with the following properties:

### area [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/alt-text.html#area)

type: `string[]`

default: `[]`

Custom components to check for alt text on `area` elements.

### img [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/alt-text.html#img)

type: `string[]`

default: `[]`

Custom components to check for alt text on `img` elements.

### input[type="image"] [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/alt-text.html#input-type-image)

type: `string[]`

default: `[]`

Custom components to check for alt text on `input[type="image"]` elements.

### object [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/alt-text.html#object)

type: `string[]`

default: `[]`

Custom components to check for alt text on `object` elements.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/alt-text.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/alt-text": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/alt-text --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/alt-text.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/alt_text.rs)
