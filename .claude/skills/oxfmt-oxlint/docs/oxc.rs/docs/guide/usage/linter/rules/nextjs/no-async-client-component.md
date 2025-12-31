---
title: "nextjs/no-async-client-component | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-async-client-component"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/nextjs/no-async-client-component.md for this page in Markdown format

# nextjs/no-async-client-component Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-async-client-component.html#nextjs-no-async-client-component)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-async-client-component.html#what-it-does)

Prevents the use of async functions for client components in Next.js applications. This rule checks for any async function that:

* Is marked with "use client" directive
* Has a name starting with an uppercase letter (indicating it's a component)
* Is either exported as default or assigned to a variable

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-async-client-component.html#why-is-this-bad)

Using async functions for client components can cause hydration mismatches between server and client, can break component rendering lifecycle, and can lead to unexpected behavior with React's concurrent features.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-async-client-component.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
"use client"

// Async component with default export
export default async function MyComponent() {
  return <></>
}

// Async component with named export
async function MyComponent() {
  return <></>
}
export default MyComponent

// Async arrow function component
const MyComponent = async () => {
  return <></>
}
export default MyComponent
```

Examples of **correct** code for this rule:

javascript

```
"use client"

// Regular synchronous component
export default function MyComponent() {
  return <></>
}

// Handling async operations in effects
export default function MyComponent() {
  useEffect(() => {
    async function fetchData() {
      // async operations here
    }
    fetchData();
  }, []);
  return <></>
}

// Async operations in event handlers
export default function MyComponent() {
  const handleClick = async () => {
    // async operations here
  }
  return <button onClick={handleClick}>Click me</button>
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-async-client-component.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["nextjs"],
  "rules": {
    "nextjs/no-async-client-component": "error"
  }
}
```

bash

```
oxlint --deny nextjs/no-async-client-component --nextjs-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-async-client-component.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/nextjs/no_async_client_component.rs)
