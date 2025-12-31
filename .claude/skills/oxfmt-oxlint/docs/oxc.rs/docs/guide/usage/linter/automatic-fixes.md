---
title: "Automatic fixes | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/automatic-fixes"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/automatic-fixes.md for this page in Markdown format

# Automatic fixes [​](https://oxc.rs/docs/guide/usage/linter/automatic-fixes.html#automatic-fixes)

Oxlint can automatically fix some lint violations. Automatic fixes are not enabled by default. You choose when to apply them.

## Safe fixes [​](https://oxc.rs/docs/guide/usage/linter/automatic-fixes.html#safe-fixes)

Safe fixes are changes that do not alter program behavior.

Apply safe fixes:

bash

```
oxlint --fix
```

## Suggestions [​](https://oxc.rs/docs/guide/usage/linter/automatic-fixes.html#suggestions)

Suggestions are changes that may alter behavior or may not match your intent.

Apply suggestions:

bash

```
oxlint --fix-suggestions
```

## Dangerous fixes [​](https://oxc.rs/docs/guide/usage/linter/automatic-fixes.html#dangerous-fixes)

Dangerous fixes are aggressive changes that may break your code.

Apply dangerous fixes:

bash

```
oxlint --fix-dangerously
```

## Combining fix modes [​](https://oxc.rs/docs/guide/usage/linter/automatic-fixes.html#combining-fix-modes)

You can combine safe fixes and suggestions:

bash

```
oxlint --fix --fix-suggestions
```

You can also include dangerous fixes:

bash

```
oxlint --fix --fix-suggestions --fix-dangerously
```

## Rule support [​](https://oxc.rs/docs/guide/usage/linter/automatic-fixes.html#rule-support)

Not all rules provide fixes. Some rules support safe fixes, some provide suggestions, and some do not provide fixes yet.

If a rule is missing a fixer, contributions are welcome.

## Type aware linting and fixes [​](https://oxc.rs/docs/guide/usage/linter/automatic-fixes.html#type-aware-linting-and-fixes)

Type aware linting requires building the project.

You can apply safe fixes with type aware linting enabled:

bash

```
oxlint --type-aware --fix
```
