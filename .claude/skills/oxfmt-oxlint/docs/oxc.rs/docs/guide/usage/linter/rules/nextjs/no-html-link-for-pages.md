---
title: "nextjs/no-html-link-for-pages | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-html-link-for-pages"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/no-html-link-for-pages.md for this page in Markdown format

# nextjs/no-html-link-for-pages Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-html-link-for-pages.html#nextjs-no-html-link-for-pages)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-html-link-for-pages.html#what-it-does)

Prevents the usage of `<a>` elements to navigate between Next.js pages.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-html-link-for-pages.html#why-is-this-bad)

Using `<a>` elements for internal navigation in Next.js applications can cause:

* Full page reloads instead of client-side navigation
* Loss of application state
* Slower navigation performance
* Broken prefetching capabilities

Next.js provides the `<Link />` component from `next/link` for client-side navigation between pages, which provides better performance and user experience.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-html-link-for-pages.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
function HomePage() {
  return (
    <div>
      <a href="/about">About Us</a>
      <a href="/contact">Contact</a>
    </div>
  );
}
```

Examples of **correct** code for this rule:

jsx

```
import Link from "next/link";

function HomePage() {
  return (
    <div>
      <Link href="/about">About Us</Link>
      <Link href="/contact">Contact</Link>
    </div>
  );
}
```

External links are allowed:

jsx

```
function HomePage() {
  return (
    <div>
      <a href="https://example.com">External Link</a>
      <a href="mailto:contact@example.com">Email</a>
      <a href="tel:+1234567890">Phone</a>
    </div>
  );
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-html-link-for-pages.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/no-html-link-for-pages": "error"
  }
}
```

bash

```
oxlint --deny nextjs/no-html-link-for-pages --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-html-link-for-pages.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/no_html_link_for_pages.rs)
