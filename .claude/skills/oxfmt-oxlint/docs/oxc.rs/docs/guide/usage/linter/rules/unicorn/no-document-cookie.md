---
title: "unicorn/no-document-cookie | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-document-cookie"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-document-cookie.md for this page in Markdown format

# unicorn/no-document-cookie Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-document-cookie.html#unicorn-no-document-cookie)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-document-cookie.html#what-it-does)

Disallow direct use of [`document.cookie`](https://developer.mozilla.org/en-US/docs/Web/API/Document/cookie).

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-document-cookie.html#why-is-this-bad)

It's not recommended to use [`document.cookie`](https://developer.mozilla.org/en-US/docs/Web/API/Document/cookie) directly as it's easy to get the string wrong. Instead, you should use the [Cookie Store API](https://developer.mozilla.org/en-US/docs/Web/API/Cookie_Store_API) or a [cookie library](https://www.npmjs.com/search?q=cookie).

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-document-cookie.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
document.cookie =
  "foo=bar" +
  "; Path=/" +
  "; Domain=example.com" +
  "; expires=Fri, 31 Dec 9999 23:59:59 GMT" +
  "; Secure";
```

Examples of **correct** code for this rule:

javascript

```
async function storeCookies() {
  await cookieStore.set({
    name: "foo",
    value: "bar",
    expires: Date.now() + 24 * 60 * 60 * 1000,
    domain: "example.com",
  });
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-document-cookie.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-document-cookie": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-document-cookie
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-document-cookie.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_document_cookie.rs)
