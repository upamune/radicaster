---
title: "typescript/use-unknown-in-catch-callback-variable | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/use-unknown-in-catch-callback-variable"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/use-unknown-in-catch-callback-variable.md for this page in Markdown format

# typescript/use-unknown-in-catch-callback-variable Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/use-unknown-in-catch-callback-variable.html#typescript-use-unknown-in-catch-callback-variable)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/use-unknown-in-catch-callback-variable.html#what-it-does)

This rule enforces using `unknown` for catch clause variables instead of `any`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/use-unknown-in-catch-callback-variable.html#why-is-this-bad)

In TypeScript 4.0+, catch clause variables can be typed as `unknown` instead of `any`. Using `unknown` is safer because it forces you to perform type checking before using the error, preventing potential runtime errors.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/use-unknown-in-catch-callback-variable.html#examples)

Examples of **incorrect** code for this rule:

ts

```
try {
  somethingRisky();
} catch (error: any) {
  // Should use 'unknown'
  console.log(error.message); // Unsafe access
  error.someMethod(); // Unsafe call
}

// Default catch variable is 'any' in older TypeScript
try {
  somethingRisky();
} catch (error) {
  // Implicitly 'any'
  console.log(error.message); // Unsafe access
}
```

Examples of **correct** code for this rule:

ts

```
try {
  somethingRisky();
} catch (error: unknown) {
  // Type guard for Error objects
  if (error instanceof Error) {
    console.log(error.message); // Safe access
    console.log(error.stack);
  } else {
    console.log("Unknown error:", error);
  }
}

// More comprehensive error handling
try {
  somethingRisky();
} catch (error: unknown) {
  if (error instanceof Error) {
    // Handle Error objects
    console.error("Error:", error.message);
  } else if (typeof error === "string") {
    // Handle string errors
    console.error("String error:", error);
  } else {
    // Handle unknown error types
    console.error("Unknown error type:", error);
  }
}

// Helper function for error handling
function isError(error: unknown): error is Error {
  return error instanceof Error;
}

try {
  somethingRisky();
} catch (error: unknown) {
  if (isError(error)) {
    console.log(error.message);
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/use-unknown-in-catch-callback-variable.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/use-unknown-in-catch-callback-variable": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/use-unknown-in-catch-callback-variable
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/use-unknown-in-catch-callback-variable.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/use_unknown_in_catch_callback_variable.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/use_unknown_in_catch_callback_variable/use_unknown_in_catch_callback_variable.go)
