---
title: "nextjs/no-unwanted-polyfillio | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-unwanted-polyfillio"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/no-unwanted-polyfillio.md for this page in Markdown format

# nextjs/no-unwanted-polyfillio Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-unwanted-polyfillio.html#nextjs-no-unwanted-polyfillio)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-unwanted-polyfillio.html#what-it-does)

Prevent use of unsafe polyfill.io domains and duplicate polyfills.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-unwanted-polyfillio.html#why-is-this-bad)

**Security Risk:** The domains `cdn.polyfill.io` and `polyfill.io` were compromised in a supply chain attack in 2024, where the domain was acquired by a malicious actor and began injecting harmful code into websites. Over 380,000+ websites were affected. These domains should not be used under any circumstances.

**Performance Issue:** For safe alternatives like `cdnjs.cloudflare.com/polyfill/`, including polyfills already shipped with Next.js unnecessarily increases page weight which can affect loading performance.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-unwanted-polyfillio.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
// Security risk - compromised domain
<script src='https://cdn.polyfill.io/v2/polyfill.min.js'></script>
<script src='https://polyfill.io/v3/polyfill.min.js'></script>

// Duplicate polyfills
<script src='https://cdnjs.cloudflare.com/polyfill/v3/polyfill.min.js?features=Array.prototype.copyWithin'></script>
<script src='https://cdnjs.cloudflare.com/polyfill/v3/polyfill.min.js?features=WeakSet%2CPromise'></script>
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-unwanted-polyfillio.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/no-unwanted-polyfillio": "error"
  }
}
```

bash

```
oxlint --deny nextjs/no-unwanted-polyfillio --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-unwanted-polyfillio.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/no_unwanted_polyfillio.rs)
