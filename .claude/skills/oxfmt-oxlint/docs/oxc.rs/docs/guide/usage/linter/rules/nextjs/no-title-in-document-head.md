---
title: "nextjs/no-title-in-document-head | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-title-in-document-head"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/no-title-in-document-head.md for this page in Markdown format

# nextjs/no-title-in-document-head Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-title-in-document-head.html#nextjs-no-title-in-document-head)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-title-in-document-head.html#what-it-does)

Prevent usage of `<title>` with `Head` component from `next/document`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-title-in-document-head.html#why-is-this-bad)

A `<title>` element should only be used for any `<head>` code that is common for all pages. Title tags should be defined at the page-level using `next/head` instead.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-title-in-document-head.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
import { Head } from "next/document";

export function Home() {
  return (
    <div>
      <Head>
        <title>My page title</title>
      </Head>
    </div>
  );
}
```

Examples of **correct** code for this rule:

javascript

```
import Head from "next/head";

export function Home() {
  return (
    <div>
      <Head>
        <title>My page title</title>
      </Head>
    </div>
  );
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-title-in-document-head.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/no-title-in-document-head": "error"
  }
}
```

bash

```
oxlint --deny nextjs/no-title-in-document-head --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-title-in-document-head.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/no_title_in_document_head.rs)
