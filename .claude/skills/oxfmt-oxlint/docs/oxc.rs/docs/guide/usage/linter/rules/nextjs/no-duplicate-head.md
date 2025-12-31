---
title: "nextjs/no-duplicate-head | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-duplicate-head"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/no-duplicate-head.md for this page in Markdown format

# nextjs/no-duplicate-head Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-duplicate-head.html#nextjs-no-duplicate-head)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-duplicate-head.html#what-it-does)

Prevent duplicate usage of `<Head>` in `pages/\_document.js``.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-duplicate-head.html#why-is-this-bad)

This can cause unexpected behavior in your application.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-duplicate-head.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
import Document, { Html, Head, Main, NextScript } from "next/document";
class MyDocument extends Document {
  static async getInitialProps(ctx) {}
  render() {
    return (
      <Html>
        <Head />
        <Head />
        <body>
          <Main />
          <NextScript />
        </body>
      </Html>
    );
  }
}
export default MyDocument;
```

Examples of **correct** code for this rule:

jsx

```
import Document, { Html, Head, Main, NextScript } from "next/document";
class MyDocument extends Document {
  static async getInitialProps(ctx) {}
  render() {
    return (
      <Html>
        <Head />
        <body>
          <Main />
          <NextScript />
        </body>
      </Html>
    );
  }
}
export default MyDocument;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-duplicate-head.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/no-duplicate-head": "error"
  }
}
```

bash

```
oxlint --deny nextjs/no-duplicate-head --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-duplicate-head.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/no_duplicate_head.rs)
