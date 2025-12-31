---
title: "nextjs/no-styled-jsx-in-document | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-styled-jsx-in-document"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/no-styled-jsx-in-document.md for this page in Markdown format

# nextjs/no-styled-jsx-in-document Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-styled-jsx-in-document.html#nextjs-no-styled-jsx-in-document)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-styled-jsx-in-document.html#what-it-does)

Prevent usage of styled-jsx in pages/\_document.js.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-styled-jsx-in-document.html#why-is-this-bad)

Custom CSS like styled-jsx is not allowed in a [Custom Document](https://nextjs.org/docs/pages/building-your-application/routing/custom-document).

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-styled-jsx-in-document.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
// pages/_document.js
import Document, { Html, Head, Main, NextScript } from "next/document";

class MyDocument extends Document {
  render() {
    return (
      <Html>
        <Head />
        <body>
          <Main />
          <NextScript />
          <style jsx>{`
            body {
              background: hotpink;
            }
          `}</style>
        </body>
      </Html>
    );
  }
}
```

Examples of **correct** code for this rule:

javascript

```
// pages/_document.js
import Document, { Html, Head, Main, NextScript } from "next/document";

class MyDocument extends Document {
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
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-styled-jsx-in-document.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/no-styled-jsx-in-document": "error"
  }
}
```

bash

```
oxlint --deny nextjs/no-styled-jsx-in-document --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-styled-jsx-in-document.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/no_styled_jsx_in_document.rs)
