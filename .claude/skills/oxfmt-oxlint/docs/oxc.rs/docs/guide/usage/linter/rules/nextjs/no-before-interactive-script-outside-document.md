---
title: "nextjs/no-before-interactive-script-outside-document | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-before-interactive-script-outside-document"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/no-before-interactive-script-outside-document.md for this page in Markdown format

# nextjs/no-before-interactive-script-outside-document Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-before-interactive-script-outside-document.html#nextjs-no-before-interactive-script-outside-document)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-before-interactive-script-outside-document.html#what-it-does)

Prevents the usage of `next/script`'s `beforeInteractive` strategy outside of `pages/_document.js`. This rule ensures that scripts with the `beforeInteractive` loading strategy are only used in the document component where they are most effective.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-before-interactive-script-outside-document.html#why-is-this-bad)

The `beforeInteractive` strategy is specifically designed to load scripts before any page hydration occurs, which is only guaranteed to work correctly when placed in `pages/_document.js`. Using it elsewhere:

* May not achieve the intended early loading behavior
* Can lead to inconsistent script loading timing
* Might cause hydration mismatches or other runtime issues
* Could impact the application's performance optimization

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-before-interactive-script-outside-document.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
// pages/index.js
import Script from "next/script";

export default function HomePage() {
  return (
    <div>
      <Script
        src="https://example.com/script.js"
        strategy="beforeInteractive" // ❌ Wrong placement
      />
    </div>
  );
}
```

Examples of **correct** code for this rule:

jsx

```
// pages/_document.js
import Document, { Html, Head, Main, NextScript } from "next/document";
import Script from "next/script";

class MyDocument extends Document {
  render() {
    return (
      <Html>
        <Head />
        <body>
          <Script
            src="https://example.com/script.js"
            strategy="beforeInteractive" // ✅ Correct placement
          />
          <Main />
          <NextScript />
        </body>
      </Html>
    );
  }
}

export default MyDocument;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-before-interactive-script-outside-document.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/no-before-interactive-script-outside-document": "error"
  }
}
```

bash

```
oxlint --deny nextjs/no-before-interactive-script-outside-document --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-before-interactive-script-outside-document.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/no_before_interactive_script_outside_document.rs)
