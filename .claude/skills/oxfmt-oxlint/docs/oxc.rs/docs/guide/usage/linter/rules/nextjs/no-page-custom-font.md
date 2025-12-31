---
title: "nextjs/no-page-custom-font | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-page-custom-font"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/no-page-custom-font.md for this page in Markdown format

# nextjs/no-page-custom-font Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-page-custom-font.html#nextjs-no-page-custom-font)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-page-custom-font.html#what-it-does)

Prevent page-only custom fonts.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-page-custom-font.html#why-is-this-bad)

* The custom font you're adding was added to a page - this only adds the font to the specific page and not the entire application.
* The custom font you're adding was added to a separate component within `pages/_document.js` - this disables automatic font optimization.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-page-custom-font.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
// pages/index.jsx
import Head from "next/head";
function IndexPage() {
  return (
    <Head>
      <link
        href="https://fonts.googleapis.com/css2?family=Krona+One&display=swap"
        rel="stylesheet"
      />
    </Head>
  );
}
export default IndexPage;
```

Examples of **correct** code for this rule:

jsx

```
// pages/_document.jsx
import NextDocument, { Html, Head } from "next/document";
class Document extends NextDocument {
  render() {
    return (
      <Html>
        <Head>
          <link
            href="https://fonts.googleapis.com/css2?family=Krona+One&display=swap"
            rel="stylesheet"
          />
        </Head>
      </Html>
    );
  }
}
export default Document;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-page-custom-font.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/no-page-custom-font": "error"
  }
}
```

bash

```
oxlint --deny nextjs/no-page-custom-font --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-page-custom-font.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/no_page_custom_font.rs)
