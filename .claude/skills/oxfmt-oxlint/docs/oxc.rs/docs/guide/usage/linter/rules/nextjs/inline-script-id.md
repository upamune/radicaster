---
title: "nextjs/inline-script-id | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/inline-script-id"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/inline-script-id.md for this page in Markdown format

# nextjs/inline-script-id Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/inline-script-id.html#nextjs-inline-script-id)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/inline-script-id.html#what-it-does)

Enforces that all `next/script` components with inline content or `dangerouslySetInnerHTML` must have an `id` prop.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/inline-script-id.html#why-is-this-bad)

Next.js requires a unique `id` prop for inline scripts to properly deduplicate them during page renders. Without an `id`, the same inline script might be executed multiple times, leading to unexpected behavior or performance issues. This is particularly important for scripts that modify global state or perform one-time initializations.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/inline-script-id.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
import Script from 'next/script';

export default function Page() {
  return (
    <Script>
      {`console.log('Hello world');`}
    </Script>
  );
}

// Also incorrect with dangerouslySetInnerHTML
export default function Page() {
  return (
    <Script
      dangerouslySetInnerHTML={{
        __html: `console.log('Hello world');`
      }}
    />
  );
}
```

Examples of **correct** code for this rule:

javascript

```
import Script from 'next/script';

export default function Page() {
  return (
    <Script id="my-script">
      {`console.log('Hello world');`}
    </Script>
  );
}

// Correct with dangerouslySetInnerHTML
export default function Page() {
  return (
    <Script
      id="my-script"
      dangerouslySetInnerHTML={{
        __html: `console.log('Hello world');`
      }}
    />
  );
}

// No id required for external scripts
export default function Page() {
  return (
    <Script src="https://example.com/script.js" />
  );
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/inline-script-id.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/inline-script-id": "error"
  }
}
```

bash

```
oxlint --deny nextjs/inline-script-id --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/inline-script-id.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/inline_script_id.rs)
