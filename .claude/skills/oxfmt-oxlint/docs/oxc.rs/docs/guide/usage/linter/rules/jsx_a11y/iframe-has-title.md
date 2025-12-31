---
title: "jsx_a11y/iframe-has-title | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/iframe-has-title"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/iframe-has-title.md for this page in Markdown format

# jsx\_a11y/iframe-has-title Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/iframe-has-title.html#jsx-a11y-iframe-has-title)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/iframe-has-title.html#what-it-does)

Enforce iframe elements have a title attribute.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/iframe-has-title.html#why-is-this-bad)

Screen reader users rely on a iframe title to describe the contents of the iframe. Navigating through iframe and iframe elements quickly becomes difficult and confusing for users of this technology if the markup does not contain a title attribute.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/iframe-has-title.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<iframe />
<iframe {...props} />
<iframe title="" />
<iframe title={''} />
<iframe title={``} />
<iframe title={undefined} />
<iframe title={false} />
<iframe title={true} />
<iframe title={42} />
```

Examples of **correct** code for this rule:

jsx

```
<iframe title="This is a unique title" />
<iframe title={uniqueTitle} />
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/iframe-has-title.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/iframe-has-title": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/iframe-has-title --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/iframe-has-title.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/iframe_has_title.rs)
