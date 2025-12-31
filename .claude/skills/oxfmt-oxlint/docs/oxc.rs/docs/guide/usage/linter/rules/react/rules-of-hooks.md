---
title: "react/rules-of-hooks | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/rules-of-hooks"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/rules-of-hooks.md for this page in Markdown format

# react/rules-of-hooks Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/react/rules-of-hooks.html#react-rules-of-hooks)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/rules-of-hooks.html#what-it-does)

Enforces the Rules of Hooks, ensuring that React Hooks are only called in valid contexts and in the correct order.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/rules-of-hooks.html#why-is-this-bad)

React Hooks must follow specific rules to ensure they work correctly:

1. Only call Hooks at the top level (never inside loops, conditions, or nested functions)
2. Only call Hooks from React function components or custom Hooks
3. Hooks must be called in the same order every time a component renders

Breaking these rules can lead to bugs where state gets corrupted or component behavior becomes unpredictable.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/rules-of-hooks.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
// Don't call Hooks inside loops, conditions, or nested functions
function BadComponent() {
  if (condition) {
    const [state, setState] = useState(); // ❌ Hook in condition
  }

  for (let i = 0; i < 10; i++) {
    useEffect(() => {}); // ❌ Hook in loop
  }
}

// Don't call Hooks from regular JavaScript functions
function regularFunction() {
  const [state, setState] = useState(); // ❌ Hook in regular function
}
```

Examples of **correct** code for this rule:

javascript

```
// ✅ Call Hooks at the top level of a React component
function GoodComponent() {
  const [state, setState] = useState();

  useEffect(() => {
    // Effect logic here
  });

  return <div>{state}</div>;
}

// ✅ Call Hooks from custom Hooks
function useCustomHook() {
  const [state, setState] = useState();
  return state;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/rules-of-hooks.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/rules-of-hooks": "error"
  }
}
```

bash

```
oxlint --deny react/rules-of-hooks --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/rules-of-hooks.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/rules_of_hooks.rs)
