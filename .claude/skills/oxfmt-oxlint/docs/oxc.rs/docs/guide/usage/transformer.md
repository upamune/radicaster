---
title: "Transformer | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/transformer"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/transformer.md for this page in Markdown format

# Transformer [​](https://oxc.rs/docs/guide/usage/transformer.html#transformer)

## Features [​](https://oxc.rs/docs/guide/usage/transformer.html#features)

* [Lowering ESNext to ES2015.](https://oxc.rs/docs/guide/usage/transformer/lowering.html)
* [Transforming TypeScript to JavaScript.](https://oxc.rs/docs/guide/usage/transformer/typescript.html)
* [Transforming JSX to JavaScript, with built-in React Refresh.](https://oxc.rs/docs/guide/usage/transformer/jsx.html)
* [Built-in support for popular plugins like styled-components.](https://oxc.rs/docs/guide/usage/transformer/plugins.html)
* [Replacing global variables.](https://oxc.rs/docs/guide/usage/transformer/global-variable-replacement.html)
* [TypeScript Isolated Declarations Emit without using the TypeScript compiler.](https://oxc.rs/docs/guide/usage/transformer/isolated-declarations.html)

## Installation [​](https://oxc.rs/docs/guide/usage/transformer.html#installation)

### Node.js [​](https://oxc.rs/docs/guide/usage/transformer.html#node-js)

* Use the node binding [oxc-transform](https://www.npmjs.com/package/oxc-transform).
* Try on [stackblitz](https://stackblitz.com/edit/oxc-transform).

### Rust [​](https://oxc.rs/docs/guide/usage/transformer.html#rust)

Use the umbrella crate [oxc](https://docs.rs/oxc) with the `transformer` feature.

Rust usage example can be found [here](https://github.com/oxc-project/oxc/blob/main/crates/oxc_transformer/examples/transformer.rs).

## Integrations [​](https://oxc.rs/docs/guide/usage/transformer.html#integrations)

* [`unplugin-oxc`](https://github.com/unplugin/unplugin-oxc)
* [`unplugin-isolated-decl`](https://github.com/unplugin/unplugin-isolated-decl)
