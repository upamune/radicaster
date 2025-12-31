---
title: "nextjs/no-css-tags | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-css-tags"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/no-css-tags.md for this page in Markdown format

# nextjs/no-css-tags Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-css-tags.html#nextjs-no-css-tags)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-css-tags.html#what-it-does)

Prevents manual inclusion of stylesheets using `<link>` tags in Next.js applications. This rule checks for `<link>` tags with `rel="stylesheet"` that reference local CSS files.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-css-tags.html#why-is-this-bad)

Next.js handles CSS imports automatically through its built-in CSS support. Manual stylesheet inclusion bypasses Next.js's built-in CSS optimization, prevents proper code splitting and optimization of styles, and may cause Flash of Unstyled Content (FOUC). This also breaks automatic CSS hot reloading during development.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-css-tags.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
// Manually including local CSS file
<link href="/_next/static/css/styles.css" rel="stylesheet" />

// In pages/_document.js
<Head>
  <link href="css/my-styles.css" rel="stylesheet" />
</Head>
```

Examples of **correct** code for this rule:

jsx

```
// Importing CSS file directly
import '../styles/global.css'

// Using CSS Modules
import styles from './Button.module.css'

// Using external stylesheets (allowed)
<link
  href="https://fonts.googleapis.com/css?family=Open+Sans"
  rel="stylesheet"
/>

// Using styled-jsx
<style jsx>{`
  .button { color: blue; }
`}</style>
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-css-tags.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/no-css-tags": "error"
  }
}
```

bash

```
oxlint --deny nextjs/no-css-tags --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-css-tags.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/no_css_tags.rs)
