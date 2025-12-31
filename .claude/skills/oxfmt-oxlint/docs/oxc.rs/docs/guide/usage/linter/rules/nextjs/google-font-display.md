---
title: "nextjs/google-font-display | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/google-font-display"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/google-font-display.md for this page in Markdown format

# nextjs/google-font-display Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/google-font-display.html#nextjs-google-font-display)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/google-font-display.html#what-it-does)

Enforce font-display behavior with Google Fonts.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/google-font-display.html#why-is-this-bad)

Specifying display=optional minimizes the risk of invisible text or layout shift. If swapping to the custom font after it has loaded is important to you, then use `display=swap`` instead.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/google-font-display.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
import Head from "next/head";

export default Test = () => {
  return (
    <Head>
      <link href="https://fonts.googleapis.com/css2?family=Krona+One" rel="stylesheet" />
    </Head>
  );
};
```

Examples of **correct** code for this rule:

jsx

```
import Head from "next/head";

export default Test = () => {
  return (
    <Head>
      <link
        href="https://fonts.googleapis.com/css2?family=Krona+One&display=optional"
        rel="stylesheet"
      />
    </Head>
  );
};
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/google-font-display.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/google-font-display": "error"
  }
}
```

bash

```
oxlint --deny nextjs/google-font-display --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/google-font-display.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/google_font_display.rs)
