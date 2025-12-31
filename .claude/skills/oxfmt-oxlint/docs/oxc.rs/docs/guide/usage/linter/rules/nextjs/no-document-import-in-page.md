---
title: "nextjs/no-document-import-in-page | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-document-import-in-page"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/no-document-import-in-page.md for this page in Markdown format

# nextjs/no-document-import-in-page Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-document-import-in-page.html#nextjs-no-document-import-in-page)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-document-import-in-page.html#what-it-does)

Prevent importing `next/document` outside of `pages/_document.js`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-document-import-in-page.html#why-is-this-bad)

Importing `next/document` outside of `pages/_document.js` can cause unexpected issues in your Next.js application.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-document-import-in-page.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
// `components/MyDocument.jsx`
import Document from "next/document";

class MyDocument extends Document {
  //...
}

export default MyDocument;
```

Examples of **correct** code for this rule:

jsx

```
// `pages/_document.jsx`
import Document from "next/document";

class MyDocument extends Document {
  //...
}

export default MyDocument;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-document-import-in-page.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/no-document-import-in-page": "error"
  }
}
```

bash

```
oxlint --deny nextjs/no-document-import-in-page --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-document-import-in-page.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/no_document_import_in_page.rs)
