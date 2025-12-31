---
title: "Multi-file analysis | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/multi-file-analysis"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/multi-file-analysis.md for this page in Markdown format

# Multi-file analysis [​](https://oxc.rs/docs/guide/usage/linter/multi-file-analysis.html#multi-file-analysis)

Multi-file analysis allows rules to use project-wide information, such as the module dependency graph, instead of analyzing each file in isolation.

## Performance and architecture [​](https://oxc.rs/docs/guide/usage/linter/multi-file-analysis.html#performance-and-architecture)

ESLint evaluates rules per file and does not provide a built-in project graph. Plugins such as [`eslint-plugin-import`](https://github.com/import-js/eslint-plugin-import) must rebuild module resolution and the module graph outside of ESLint’s core, and projects report `import/no-cycle` taking tens of seconds to over a minute on large repositories.

Oxlint implements multi-file analysis in the core engine. It lints files and builds the module graph in parallel, shares parsing and resolution across rules, and typically completes comparable `import/no-cycle` checks in a few seconds.

## Enable the import plugin [​](https://oxc.rs/docs/guide/usage/linter/multi-file-analysis.html#enable-the-import-plugin)

To use multi-file analysis, you must enable the `import` plugin and configure at least one `import/*` rule.

## Detect cycles with `import/no-cycle` [​](https://oxc.rs/docs/guide/usage/linter/multi-file-analysis.html#detect-cycles-with-import-no-cycle)

Enable [`import/no-cycle`](https://oxc.rs/docs/guide/usage/linter/rules/import/no-cycle.html) to detect circular dependencies.

Import cycles:

* obscure dependency direction
* make refactors harder
* can cause imported values to be `undefined` due to evaluation order

.oxlintrc.json

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/no-cycle": ["error", { "maxDepth": 3 }]
  }
}
```

## TypeScript path aliases [​](https://oxc.rs/docs/guide/usage/linter/multi-file-analysis.html#typescript-path-aliases)

When running `import/*` rules, Oxlint automatically discovers `tsconfig.json` to resolve TypeScript path aliases such as `compilerOptions.paths`.
