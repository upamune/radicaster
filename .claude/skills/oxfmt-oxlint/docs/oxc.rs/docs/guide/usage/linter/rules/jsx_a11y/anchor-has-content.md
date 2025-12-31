---
title: "jsx_a11y/anchor-has-content | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-has-content"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/anchor-has-content.md for this page in Markdown format

# jsx\_a11y/anchor-has-content Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-has-content.html#jsx-a11y-anchor-has-content)

💡 A suggestion is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-has-content.html#what-it-does)

Enforce that anchors have content and that the content is accessible to screen readers. Accessible means that it is not hidden using the `aria-hidden` prop.

Alternatively, you may use the `title` prop or the `aria-label` prop.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-has-content.html#why-is-this-bad)

### Example [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-has-content.html#example)

#### good [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-has-content.html#good)

```
<a>Anchor Content!</a>
 <a><TextWrapper /></a>
 <a dangerouslySetInnerHTML={{ __html: 'foo' }} />
 <a title='foo' />
 <a aria-label='foo' />
```

#### bad [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-has-content.html#bad)

```
<a />
<a><TextWrapper aria-hidden /></a>
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-has-content.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/anchor-has-content": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/anchor-has-content --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-has-content.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/anchor_has_content.rs)
