---
title: "nextjs/no-img-element | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-img-element"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/no-img-element.md for this page in Markdown format

# nextjs/no-img-element Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-img-element.html#nextjs-no-img-element)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-img-element.html#what-it-does)

Prevent the usage of `<img>` element due to slower [LCP](https://nextjs.org/learn/seo/lcp) and higher bandwidth.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-img-element.html#why-is-this-bad)

`<img>` elements are not optimized for performance and can result in slower LCP and higher bandwidth. Using [`<Image />`](https://nextjs.org/docs/pages/api-reference/components/image) from `next/image` will automatically optimize images and serve them as static assets.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-img-element.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
export function MyComponent() {
  return (
    <div>
      <img src="/test.png" alt="Test picture" />
    </div>
  );
}
```

Examples of **correct** code for this rule:

javascript

```
import Image from "next/image";
import testImage from "./test.png";
export function MyComponent() {
  return (
    <div>
      <Image src={testImage} alt="Test picture" />
    </div>
  );
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-img-element.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/no-img-element": "error"
  }
}
```

bash

```
oxlint --deny nextjs/no-img-element --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-img-element.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/no_img_element.rs)
