---
title: "Migrating from Prettier to Oxfmt | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/formatter/migrate-from-prettier.md for this page in Markdown format

# Migrating from Prettier to Oxfmt [​](https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier.html#migrating-from-prettier-to-oxfmt)

If you currently use Prettier as your code formatter, you can follow this guide to migrate to Oxfmt.

Note that Oxfmt is in alpha, and may not be suitable for production use in complex setups.

## Caveats for migrating to Oxfmt [​](https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier.html#caveats-for-migrating-to-oxfmt)

Before migrating, ensure that the current release of the Oxfmt alpha meets your project's needs.

It is almost entirely compatible with Prettier v3.7 already for basic configurations, but less-common config options and other features are not yet implemented.

Other important considerations when migrating from Prettier to Oxfmt:

* Oxfmt's formatting output is closest to Prettier v3.7. You will see more differences migrating from an older version of Prettier.
* Oxfmt uses a `printWidth` of 100 characters by default, whereas Prettier's default is 80. Make sure to set `"printWidth": 80` in `.oxfmtrc.jsonc` to minimize differences if you use the Prettier default.
* Prettier plugins are not yet supported.
* Some Prettier options are not supported. See the [oxfmt CLI documentation](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html) for the full list of currently-supported options.
* Oxfmt supports an `--lsp` flag to spin up a Language Server Protocol server, but editor/IDE integration is still being developed and has not been tested/documented yet for most editors.

Many of these limitations will be addressed in the future, with the Beta or Stable releases of Oxfmt.

See also [the Oxfmt FAQ](https://oxc.rs/docs/guide/usage/formatter.html#faqs) for any other potential caveats or limitations you may need to consider.

## Step 1: Upgrade Prettier to v3.7 (Optional) [​](https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier.html#step-1-upgrade-prettier-to-v3-7-optional)

This step is optional, but will make it easier to determine which differences between Oxfmt and Prettier are "real".

To minimize the number of changes when migrating to Oxfmt, you should upgrade Prettier to version 3.7 first and reformat all JS/TS files with it, as it is the latest release of Prettier (from Nov 2025) and will be most similar to the output of Oxfmt.

## Step 2: Install Oxfmt [​](https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier.html#step-2-install-oxfmt)

Install Oxfmt as a development dependency with your package manager of choice:

npmpnpmyarnbundeno

bash

```
$ npm add -D oxfmt@latest
```

bash

```
$ pnpm add -D oxfmt@latest
```

bash

```
$ yarn add -D oxfmt@latest
```

bash

```
$ bun add -D oxfmt@latest
```

bash

```
$ deno add -D npm:oxfmt@latest
```

## Step 3: Migrate Prettier configuration file [​](https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier.html#step-3-migrate-prettier-configuration-file)

`.oxfmtrc.jsonc` is the configuration file for Oxfmt. Only JSON files are supported.

A basic `.oxfmtrc.jsonc` file looks like this:

jsonc

```
{
  "$schema": "./node_modules/oxfmt/configuration_schema.json",
  "printWidth": 80,
}
```

If you have a basic `.prettierrc` file, you can simply rename the file with `mv .prettierrc .oxfmtrc.jsonc`.

If you are using something other than JSON to configure Prettier, you will need to convert the configuration to JSON.

In either case, you can use the migrate command in your project root directory.

sh

```
oxfmt --migrate prettier
```

This command automatically finds your configuration file and converts it to `.oxfmtrc.json` if possible.

### `prettierrc.js` [​](https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier.html#prettierrc-js)

Here's an example of migrating a `prettierrc.js` file.

Before:

js

```
module.exports = {
  singleQuote: true,
  jsxSingleQuote: true,
};
```

After (`.oxfmtrc.jsonc`):

jsonc

```
{
  "$schema": "./node_modules/oxfmt/configuration_schema.json",
  "singleQuote": true,
  "jsxSingleQuote": true,
  "printWidth": 80,
}
```

### `prettierrc.yaml` [​](https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier.html#prettierrc-yaml)

Here's an example of migrating a `prettierrc.yaml` file.

Before:

yaml

```
trailingComma: "es5"
tabWidth: 4
semi: false
singleQuote: true
```

After (`.oxfmtrc.jsonc`):

jsonc

```
{
  "$schema": "./node_modules/oxfmt/configuration_schema.json",
  "trailingComma": "es5",
  "tabWidth": 4,
  "semi": false,
  "singleQuote": true,
  "printWidth": 80,
}
```

## Step 4: Update Formatting Scripts [​](https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier.html#step-4-update-formatting-scripts)

Update any formatting scripts you currently have, for example in `package.json`, shell scripts, or pre-commit scripts.

### `package.json` scripts [​](https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier.html#package-json-scripts)

diff

```
{
  "scripts": {
-   "format": "prettier --write .",
+   "format": "oxfmt",
-   "format:check": "prettier --check ."
+   "format:check": "oxfmt --check"
  }
}
```

### CI Workflows [​](https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier.html#ci-workflows)

Update any CI workflows that run Prettier, particularly `prettier --check`.

diff

```
  - name: Check formatting
-   run: yarn prettier --check .
+   run: yarn oxfmt --check
```

### Git Hooks (e.g. husky, lint-staged) [​](https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier.html#git-hooks-e-g-husky-lint-staged)

diff

```
"lint-staged": {
- "*": "prettier --write --no-error-on-unmatched-pattern"
+ "*": "oxfmt --no-error-on-unmatched-pattern"
}
```

## Step 5: Run formatter [​](https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier.html#step-5-run-formatter)

Run Oxfmt on your codebase to check for any changes and ensure that the configuration was migrated correctly:

sh

```
# Your script specified in Step 4
npm run format
```

If you no longer need Prettier, you can uninstall for now.

## Done! [​](https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier.html#done)

You have now migrated to Oxfmt 😃

Please see the section below for any additional, optional steps you may need to take.

These are only applicable for some setups, so skip them if they don't apply to you.

### Update editor integrations [​](https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier.html#update-editor-integrations)

See [the Formatter FAQ](https://oxc.rs/docs/guide/usage/formatter.html#how-does-editor-integration-work) for details on editor/IDE integration with Oxfmt.

### Update `CONTRIBUTING.md` and `AGENTS.md`/`CLAUDE.md` [​](https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier.html#update-contributing-md-and-agents-md-claude-md)

If you have a `CONTRIBUTING.md` file that references Prettier, update those references to use Oxfmt.

If you use an `AGENTS.md` or `CLAUDE.md` file to help LLM tools understand your codebase, you should also check for references to Prettier in those files.

### Update lint rules [​](https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier.html#update-lint-rules)

If you have any lint rules that explicitly check for Prettier formatting (e.g. `eslint-plugin-prettier`), you should remove them.

While you're at it, you could also consider migrating to [oxlint](https://oxc.rs/docs/guide/usage/linter.html) 😉

### Create/update `.git-blame-ignore-revs` [​](https://oxc.rs/docs/guide/usage/formatter/migrate-from-prettier.html#create-update-git-blame-ignore-revs)

If you want to avoid extra noise in your `git blame` history, you can add the commit SHA where you reformatted files using Oxfmt to your `.git-blame-ignore-revs` file. This will make `git blame` ignore that commit when showing blame information. This file is supported natively by git, and by both GitHub and GitLab.
