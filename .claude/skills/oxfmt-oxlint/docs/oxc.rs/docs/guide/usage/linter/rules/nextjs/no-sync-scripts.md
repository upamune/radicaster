---
title: "nextjs/no-sync-scripts | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-sync-scripts"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/no-sync-scripts.md for this page in Markdown format

# nextjs/no-sync-scripts Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-sync-scripts.html#nextjs-no-sync-scripts)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-sync-scripts.html#what-it-does)

This rule prevents the use of synchronous `<script>` tags in Next.js applications. It requires that any `<script>` tag with a `src` attribute must also have either the `async` or `defer` attribute.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-sync-scripts.html#why-is-this-bad)

Synchronous scripts can block the page rendering and negatively impact performance. In Next.js applications, it's recommended to use `async` or `defer` attributes to load scripts asynchronously, which improves page load times and user experience.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-sync-scripts.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
// Synchronous script without async/defer
<script src="https://example.com/script.js"></script>

// Dynamic src without async/defer
<script src={dynamicSrc}></script>
```

Examples of **correct** code for this rule:

javascript

```
// Script with async attribute
<script src="https://example.com/script.js" async></script>

// Script with defer attribute
<script src="https://example.com/script.js" defer></script>

// Script with spread props (allowed as it might include async/defer)
<script {...props}></script>
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-sync-scripts.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/no-sync-scripts": "error"
  }
}
```

bash

```
oxlint --deny nextjs/no-sync-scripts --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-sync-scripts.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/no_sync_scripts.rs)
