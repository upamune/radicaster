---
title: "unicorn/prefer-node-protocol | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-node-protocol"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-node-protocol.md for this page in Markdown format

# unicorn/prefer-node-protocol Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-node-protocol.html#unicorn-prefer-node-protocol)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-node-protocol.html#what-it-does)

Prefer using the `node:protocol` when importing Node.js builtin modules

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-node-protocol.html#why-is-this-bad)

Node.js builtin modules should be imported using the `node:` protocol to avoid ambiguity with local modules.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-node-protocol.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
import fs from "fs";
```

Examples of **correct** code for this rule:

javascript

```
import fs from "node:fs";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-node-protocol.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-node-protocol": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-node-protocol
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-node-protocol.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_node_protocol.rs)
