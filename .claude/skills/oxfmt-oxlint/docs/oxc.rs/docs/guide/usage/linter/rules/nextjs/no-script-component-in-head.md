---
title: "nextjs/no-script-component-in-head | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-script-component-in-head"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/no-script-component-in-head.md for this page in Markdown format

# nextjs/no-script-component-in-head Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-script-component-in-head.html#nextjs-no-script-component-in-head)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-script-component-in-head.html#what-it-does)

Prevent usage of `next/script` in `next/head` component.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-script-component-in-head.html#why-is-this-bad)

The `next/script` component should not be used in a `next/head` component. Instead move the `<Script />` component outside of `<Head>` instead.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-script-component-in-head.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
import Script from "next/script";
import Head from "next/head";

export default function Index() {
  return (
    <Head>
      <title>Next.js</title>
      <Script src="/my-script.js" />
    </Head>
  );
}
```

Examples of **correct** code for this rule:

jsx

```
import Script from "next/script";
import Head from "next/head";

export default function Index() {
  return (
    <>
      <Head>
        <title>Next.js</title>
      </Head>
      <Script src="/my-script.js" />
    </>
  );
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-script-component-in-head.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/no-script-component-in-head": "error"
  }
}
```

bash

```
oxlint --deny nextjs/no-script-component-in-head --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-script-component-in-head.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/no_script_component_in_head.rs)
