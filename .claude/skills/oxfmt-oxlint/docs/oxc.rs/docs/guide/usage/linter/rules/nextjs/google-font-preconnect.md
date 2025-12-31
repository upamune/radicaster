---
title: "nextjs/google-font-preconnect | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/google-font-preconnect"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/google-font-preconnect.md for this page in Markdown format

# nextjs/google-font-preconnect Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/google-font-preconnect.html#nextjs-google-font-preconnect)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/google-font-preconnect.html#what-it-does)

Enforces the presence of `rel="preconnect"` when using Google Fonts via `<link>` tags.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/google-font-preconnect.html#why-is-this-bad)

When using Google Fonts, it's recommended to include a preconnect resource hint to establish early connections to the required origin. Without preconnect, the browser needs to perform DNS lookups, TCP handshakes, and TLS negotiations before it can download the font files, which can delay font loading and impact performance.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/google-font-preconnect.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
<link href="https://fonts.gstatic.com" />
<link rel="preload" href="https://fonts.gstatic.com" />
```

Examples of **correct** code for this rule:

javascript

```
<link rel="preconnect" href="https://fonts.gstatic.com" />
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/google-font-preconnect.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/google-font-preconnect": "error"
  }
}
```

bash

```
oxlint --deny nextjs/google-font-preconnect --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/google-font-preconnect.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/google_font_preconnect.rs)
