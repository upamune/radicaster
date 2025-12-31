---
title: "eslint/no-unused-labels | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-labels"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-unused-labels.md for this page in Markdown format

# eslint/no-unused-labels Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-labels.html#eslint-no-unused-labels)

✅ This rule is turned on by default.

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-labels.html#what-it-does)

Disallow unused labels

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-labels.html#why-is-this-bad)

Labels that are declared and not used anywhere in the code are most likely an error due to incomplete refactoring.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-labels.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
OUTER_LOOP: for (const student of students) {
  if (checkScores(student.scores)) {
    continue;
  }
  doSomething(student);
}
```

Examples of **correct** code for this rule:

javascript

```
for (const student of students) {
  if (checkScores(student.scores)) {
    continue;
  }
  doSomething(student);
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-labels.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-unused-labels": "error"
  }
}
```

bash

```
oxlint --deny no-unused-labels
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-labels.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_unused_labels.rs)
