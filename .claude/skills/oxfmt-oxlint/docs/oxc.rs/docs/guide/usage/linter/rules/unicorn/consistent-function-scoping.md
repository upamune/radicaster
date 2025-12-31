---
title: "unicorn/consistent-function-scoping | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-function-scoping"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/consistent-function-scoping.md for this page in Markdown format

# unicorn/consistent-function-scoping Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-function-scoping.html#unicorn-consistent-function-scoping)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-function-scoping.html#what-it-does)

Disallow functions that are declared in a scope which does not capture any variables from the outer scope.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-function-scoping.html#why-is-this-bad)

Moving function declarations to the highest possible scope improves readability, directly [improves performance](https://stackoverflow.com/questions/80802/does-use-of-anonymous-functions-affect-performance/81329#81329) and allows JavaScript engines to better [optimize your performance](https://ponyfoo.com/articles/javascript-performance-pitfalls-v8#optimization-limit).

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-function-scoping.html#examples)

Examples of **incorrect** code for this rule:

js

```
export function doFoo(foo) {
  // Does not capture anything from the scope, can be moved to the outer scope
  function doBar(bar) {
    return bar === "bar";
  }
  return doBar;
}

function doFoo(foo) {
  const doBar = (bar) => {
    return bar === "bar";
  };
}
```

Examples of **correct** code for this rule:

js

```
function doBar(bar) {
  return bar === "bar";
}

export function doFoo(foo) {
  return doBar;
}

export function doFoo(foo) {
  function doBar(bar) {
    return bar === "bar" && foo.doBar(bar);
  }
  return doBar;
}
```

### Limitations [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-function-scoping.html#limitations)

This rule does not detect or remove extraneous code blocks inside of functions:

js

```
function doFoo(foo) {
  {
    function doBar(bar) {
      return bar;
    }
  }

  return foo;
}
```

It also ignores functions that contain `JSXElement` references:

jsx

```
function doFoo(FooComponent) {
  function Bar() {
    return <FooComponent />;
  }

  return Bar;
}
```

[Immediately invoked function expressions (IIFE)](https://en.wikipedia.org/wiki/Immediately_invoked_function_expression) are ignored:

js

```
(function () {
  function doFoo(bar) {
    return bar;
  }
})();
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-function-scoping.html#configuration)

This rule accepts a configuration object with the following properties:

### checkArrowFunctions [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-function-scoping.html#checkarrowfunctions)

type: `boolean`

default: `true`

Whether to check scoping with arrow functions.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-function-scoping.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/consistent-function-scoping": "error"
  }
}
```

bash

```
oxlint --deny unicorn/consistent-function-scoping
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-function-scoping.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/consistent_function_scoping.rs)
