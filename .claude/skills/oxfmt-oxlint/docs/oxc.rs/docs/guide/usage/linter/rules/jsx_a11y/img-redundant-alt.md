---
title: "jsx_a11y/img-redundant-alt | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/img-redundant-alt"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/img-redundant-alt.md for this page in Markdown format

# jsx\_a11y/img-redundant-alt Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/img-redundant-alt.html#jsx-a11y-img-redundant-alt)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/img-redundant-alt.html#what-it-does)

Enforce that `img` alt attributes do not contain redundant words like "image", "picture", or "photo".

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/img-redundant-alt.html#why-is-this-bad)

Screen readers already announce `img` elements as an image, so there is no need to use words such as "image", "photo", or "picture" in the alt text. This creates redundant information for users of assistive technologies and makes the alt text less concise and useful.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/img-redundant-alt.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<img src="foo" alt="Photo of foo being weird." />
<img src="bar" alt="Image of me at a bar!" />
<img src="baz" alt="Picture of baz fixing a bug." />
```

Examples of **correct** code for this rule:

jsx

```
<img src="foo" alt="Foo eating a sandwich." />
<img src="bar" aria-hidden alt="Picture of me taking a photo of an image" /> // Will pass because it is hidden.
<img src="baz" alt={`Baz taking a ${photo}`} /> // This is valid since photo is a variable name.
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/img-redundant-alt.html#configuration)

This rule accepts a configuration object with the following properties:

### components [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/img-redundant-alt.html#components)

type: `string[]`

default: `["img"]`

JSX element types to validate (component names) where the rule applies. For example, `["img", "Image"]`.

### words [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/img-redundant-alt.html#words)

type: `string[]`

default: `["image", "photo", "picture"]`

Words considered redundant in alt text that should trigger a warning.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/img-redundant-alt.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/img-redundant-alt": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/img-redundant-alt --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/img-redundant-alt.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/img_redundant_alt.rs)
