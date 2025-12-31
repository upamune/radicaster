---
title: "unicorn/prefer-blob-reading-methods | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-blob-reading-methods"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-blob-reading-methods.md for this page in Markdown format

# unicorn/prefer-blob-reading-methods Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-blob-reading-methods.html#unicorn-prefer-blob-reading-methods)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-blob-reading-methods.html#what-it-does)

Recommends using `Blob#text()` and `Blob#arrayBuffer()` over `FileReader#readAsText()` and `FileReader#readAsArrayBuffer()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-blob-reading-methods.html#why-is-this-bad)

`FileReader` predates promises, and the newer [`Blob#arrayBuffer()`](https://developer.mozilla.org/en-US/docs/Web/API/Blob/arrayBuffer) and [`Blob#text()`](https://developer.mozilla.org/en-US/docs/Web/API/Blob/text) methods are much cleaner and easier to use.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-blob-reading-methods.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
async function bad() {
  const arrayBuffer = await new Promise((resolve, reject) => {
    const fileReader = new FileReader();
    fileReader.addEventListener("load", () => {
      resolve(fileReader.result);
    });
    fileReader.addEventListener("error", () => {
      reject(fileReader.error);
    });
    fileReader.readAsArrayBuffer(blob);
  });
}
```

Examples of **correct** code for this rule:

javascript

```
async function good() {
  const arrayBuffer = await blob.arrayBuffer();
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-blob-reading-methods.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-blob-reading-methods": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-blob-reading-methods
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-blob-reading-methods.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_blob_reading_methods.rs)
