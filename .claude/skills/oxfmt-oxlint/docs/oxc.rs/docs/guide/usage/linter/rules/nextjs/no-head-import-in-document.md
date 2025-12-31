---
title: "nextjs/no-head-import-in-document | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-head-import-in-document"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/no-head-import-in-document.md for this page in Markdown format

# nextjs/no-head-import-in-document Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-head-import-in-document.html#nextjs-no-head-import-in-document)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-head-import-in-document.html#what-it-does)

Prevents the usage of `next/head` inside a Next.js document.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-head-import-in-document.html#why-is-this-bad)

Importing `next/head` inside `pages/_document.js` can cause unexpected issues in your Next.js application.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-head-import-in-document.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
import Document, { Html, Main, NextScript } from "next/document";
import Head from "next/head";

class MyDocument extends Document {
  static async getInitialProps(ctx) {
    //...
  }

  render() {
    return (
      <Html>
        <Head></Head>
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
  static async getInitialProps(ctx) {
    //...
  }

  render() {
    return (
      <Html>
        <Head></Head>
      </Html>
    );
  }
}

export default MyDocument;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-head-import-in-document.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/no-head-import-in-document": "error"
  }
}
```

bash

```
oxlint --deny nextjs/no-head-import-in-document --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-head-import-in-document.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/no_head_import_in_document.rs)
