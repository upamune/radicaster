---
title: "Minifier | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/minifier"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/minifier.md for this page in Markdown format

# Minifier [​](https://oxc.rs/docs/guide/usage/minifier.html#minifier)

TIP

`oxc-minify` is currently in alpha and may still have bugs. We recommend thoroughly testing its output before deploying to production environments.

## Features [​](https://oxc.rs/docs/guide/usage/minifier.html#features)

* [Eliminate dead code.](https://oxc.rs/docs/guide/usage/minifier/dead-code-elimination.html)
* [Transforms syntaxes to make the output shorter and repetitive.](https://oxc.rs/docs/guide/usage/minifier/syntax-normalization.html)
* [Mangle variable names.](https://oxc.rs/docs/guide/usage/minifier/mangling.html)
* [Remove whitespace and comments.](https://oxc.rs/docs/guide/usage/minifier/whitespace-stripping.html)

## Assumptions [​](https://oxc.rs/docs/guide/usage/minifier.html#assumptions)

To allow better optimizations, Oxc minifier makes some assumptions about your code. See [Assumptions document](https://github.com/oxc-project/oxc/blob/main/crates/oxc_minifier/docs/ASSUMPTIONS.md) for more information.

## FAQ [​](https://oxc.rs/docs/guide/usage/minifier.html#faq)

See [FAQ](https://oxc.rs/docs/guide/usage/minifier/faq.html) for common questions.

## Installation [​](https://oxc.rs/docs/guide/usage/minifier.html#installation)

### With Rolldown [​](https://oxc.rs/docs/guide/usage/minifier.html#with-rolldown)

If you are using [Rolldown](https://rolldown.rs), `oxc-minify` will be used for minification by default. No extra installation is required.

### Node.js [​](https://oxc.rs/docs/guide/usage/minifier.html#node-js)

* Use the node binding [oxc-minify](https://www.npmjs.com/package/oxc-minify).
* Try on [stackblitz](https://stackblitz.com/edit/oxc-minify).

### Rust [​](https://oxc.rs/docs/guide/usage/minifier.html#rust)

Use the umbrella crate [oxc](https://docs.rs/oxc) with the `minifier` feature.

Rust usage example can be found [here](https://github.com/oxc-project/oxc/blob/main/crates/oxc_minifier/examples/minifier.rs).

## Integrations [​](https://oxc.rs/docs/guide/usage/minifier.html#integrations)

* [`unplugin-oxc`](https://github.com/unplugin/unplugin-oxc)
