---
title: "nextjs/no-head-element | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-head-element"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/no-head-element.md for this page in Markdown format

# nextjs/no-head-element Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-head-element.html#nextjs-no-head-element)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-head-element.html#what-it-does)

Prevents the usage of the native `<head>` element inside a Next.js application.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-head-element.html#why-is-this-bad)

A `<head>` element can cause unexpected behavior in a Next.js application. Use Next.js' built-in `next/head` component instead.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-head-element.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
function Index() {
  return (
    <>
      <head>
        <title>My page title</title>
        <meta name="viewport" content="initial-scale=1.0, width=device-width" />
      </head>
    </>
  );
}

export default Index;
```

Examples of **correct** code for this rule:

jsx

```
import Head from "next/head";

function Index() {
  return (
    <>
      <Head>
        <title>My page title</title>
        <meta name="viewport" content="initial-scale=1.0, width=device-width" />
      </Head>
    </>
  );
}

export default Index;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-head-element.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/no-head-element": "error"
  }
}
```

bash

```
oxlint --deny nextjs/no-head-element --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-head-element.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/no_head_element.rs)
