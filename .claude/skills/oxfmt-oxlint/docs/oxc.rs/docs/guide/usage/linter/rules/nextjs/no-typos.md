---
title: "nextjs/no-typos | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-typos"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/no-typos.md for this page in Markdown format

# nextjs/no-typos Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-typos.html#nextjs-no-typos)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-typos.html#what-it-does)

Detects common typos in Next.js data fetching function names.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-typos.html#why-is-this-bad)

Next.js will not call incorrectly named data fetching functions, causing pages to render without expected data.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-typos.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
export default function Page() {
  return <div></div>;
}
export async function getServurSideProps() {}
```

Examples of **correct** code for this rule:

javascript

```
export default function Page() {
  return <div></div>;
}
export async function getServerSideProps() {}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-typos.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/no-typos": "error"
  }
}
```

bash

```
oxlint --deny nextjs/no-typos --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-typos.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/no_typos.rs)
