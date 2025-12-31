---
title: "eslint/no-extra-label | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-label"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-extra-label.md for this page in Markdown format

# eslint/no-extra-label Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-label.html#eslint-no-extra-label)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-label.html#what-it-does)

Disallow unnecessary labels.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-label.html#why-is-this-bad)

If a loop contains no nested loops or switches, labeling the loop is unnecessary.

js

```
A: while (a) {
  break A;
}
```

You can achieve the same result by removing the label and using `break` or `continue` without a label. Probably those labels would confuse developers because they expect labels to jump to further.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-label.html#examples)

Examples of **incorrect** code for this rule:

js

```
A: while (a) {
  break A;
}

B: for (let i = 0; i < 10; ++i) {
  break B;
}

C: switch (a) {
  case 0:
    break C;
}
```

Examples of **correct** code for this rule:

js

```
while (a) {
  break;
}

for (let i = 0; i < 10; ++i) {
  break;
}

switch (a) {
  case 0:
    break;
}

A: {
  break A;
}

B: while (a) {
  while (b) {
    break B;
  }
}

C: switch (a) {
  case 0:
    while (b) {
      break C;
    }
    break;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-label.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-extra-label": "error"
  }
}
```

bash

```
oxlint --deny no-extra-label
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-label.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_extra_label.rs)
