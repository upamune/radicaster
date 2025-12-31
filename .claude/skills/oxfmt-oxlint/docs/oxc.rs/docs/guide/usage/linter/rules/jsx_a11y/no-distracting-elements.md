---
title: "jsx_a11y/no-distracting-elements | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-distracting-elements"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/no-distracting-elements.md for this page in Markdown format

# jsx\_a11y/no-distracting-elements Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-distracting-elements.html#jsx-a11y-no-distracting-elements)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-distracting-elements.html#what-it-does)

Enforces that no distracting elements are used.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-distracting-elements.html#why-is-this-bad)

Elements that can be visually distracting can cause accessibility issues with visually impaired users. Such elements are most likely deprecated, and should be avoided. By default, `<marquee>` and `<blink>` elements are visually distracting and can trigger vestibular disorders.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-distracting-elements.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<marquee />
<marquee {...props} />
<marquee lang={undefined} />
<blink />
<blink {...props} />
<blink foo={undefined} />
```

Examples of **correct** code for this rule:

jsx

```
<div />
<Marquee />
<Blink />
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-distracting-elements.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/no-distracting-elements": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/no-distracting-elements --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-distracting-elements.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/no_distracting_elements.rs)
