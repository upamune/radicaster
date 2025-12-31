---
title: "unicorn/prefer-top-level-await | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-top-level-await"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-top-level-await.md for this page in Markdown format

# unicorn/prefer-top-level-await Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-top-level-await.html#unicorn-prefer-top-level-await)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-top-level-await.html#what-it-does)

Prefer top-level await over top-level promises and async function calls.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-top-level-await.html#why-is-this-bad)

Top-level await is more readable and can prevent unhandled rejections.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-top-level-await.html#examples)

Examples of **incorrect** code for this rule:

js

```
(async () => {
  await run();
})();

run().catch((error) => {
  console.error(error);
});
```

Examples of **correct** code for this rule:

js

```
await run();

try {
  await run();
} catch (error) {
  console.error(error);
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-top-level-await.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-top-level-await": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-top-level-await
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-top-level-await.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_top_level_await.rs)
