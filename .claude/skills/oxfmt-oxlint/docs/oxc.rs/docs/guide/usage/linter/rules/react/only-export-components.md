---
title: "react/only-export-components | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/only-export-components"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/only-export-components.md for this page in Markdown format

# react/only-export-components Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/react/only-export-components.html#react-only-export-components)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/only-export-components.html#what-it-does)

Ensures that modules only **export React components (and related HMR-safe items)** so that Fast Refresh (a.k.a. hot reloading) can safely preserve component state. Concretely, it validates the shape of your module’s exports and common entrypoints (e.g. `createRoot(...).render(<App />)`) to match what integrations like `react-refresh` expect. The rule name is `react-refresh/only-export-components`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/only-export-components.html#why-is-this-bad)

Fast Refresh can only reliably retain state if a module exports components and avoids patterns that confuse the refresh runtime. Problematic patterns (like `export *`, anonymous default functions, exporting arrays of JSX, or mixing non-component exports in unsupported ways) can cause:

* Components to remount and lose state on edit
* Missed updates (no refresh) or overly broad reloads
* Fragile HMR behavior that differs between bundlers

By enforcing predictable exports, edits stay fast and stateful during development.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/only-export-components.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
// 1) Mixing util exports with components in unsupported ways
export const foo = () => {}; // util, not a component
export const Bar = () => <></>; // component
```

jsx

```
// 2) Anonymous default export (name is required)
export default function () {}
```

jsx

```
// 3) Re-exporting everything hides what’s exported
export * from "./foo";
```

jsx

```
// 4) Exporting JSX collections makes components non-discoverable
const Tab = () => null;
export const tabs = [<Tab />, <Tab />];
```

jsx

```
// 5) Bootstrapping a root within the same module that defines components
const App = () => null;
createRoot(document.getElementById("root")).render(<App />);
```

Examples of **correct** code for this rule:

jsx

```
// Named or default component exports are fine
export default function Foo() {
  return null;
}
```

jsx

```
// Utilities may coexist if allowed by options or naming conventions
const foo = () => {};
export const Bar = () => null;
```

jsx

```
// Entrypoint files may render an imported component
import { App } from "./App";
createRoot(document.getElementById("root")).render(<App />);
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/only-export-components.html#configuration)

This rule accepts a configuration object with the following properties:

### allowConstantExport [​](https://oxc.rs/docs/guide/usage/linter/rules/react/only-export-components.html#allowconstantexport)

type: `boolean | null`

default: `null`

Allow exporting primitive constants (string/number/boolean/template literal) alongside component exports without triggering a violation. Recommended when your bundler’s Fast Refresh integration supports this (enabled by the plugin’s `vite` preset).

jsx

```
// Allowed when allowConstantExport: true
export const VERSION = "3";
export const Foo = () => null;
```

### allowExportNames [​](https://oxc.rs/docs/guide/usage/linter/rules/react/only-export-components.html#allowexportnames)

type: `string[]`

default: `null`

Treat specific named exports as HMR-safe (useful for frameworks that hot-replace certain exports). For example, in Remix: `{ "allowExportNames": ["meta", "links", "headers", "loader", "action"] }`

### checkJS [​](https://oxc.rs/docs/guide/usage/linter/rules/react/only-export-components.html#checkjs)

type: `boolean | null`

default: `null`

Check `.js` files that contain JSX (in addition to `.tsx`/`.jsx`). To reduce false positives, only files that import React are checked when this is enabled.

### customHOCs [​](https://oxc.rs/docs/guide/usage/linter/rules/react/only-export-components.html#customhocs)

type: `string[]`

default: `null`

If you export components wrapped in custom higher-order components, list their identifiers here to avoid false positives.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/only-export-components.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/only-export-components": "error"
  }
}
```

bash

```
oxlint --deny react/only-export-components --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/only-export-components.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/only_export_components.rs)
