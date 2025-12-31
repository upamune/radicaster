---
title: "Oxlint Configuration File | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/config-file-reference"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/config-file-reference.md for this page in Markdown format

# Oxlint Configuration File [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#oxlint-configuration-file)

This configuration is aligned with ESLint v8's configuration schema (`eslintrc.json`).

Usage: `oxlint -c oxlintrc.json --import-plugin`

NOTE

Only the `.json` format is supported. You can use comments in configuration files.

Example

`.oxlintrc.json`

json

```
{
  "$schema": "./node_modules/oxlint/configuration_schema.json",
  "plugins": ["import", "typescript", "unicorn"],
  "env": {
    "browser": true
  },
  "globals": {
    "foo": "readonly"
  },
  "settings": {},
  "rules": {
    "eqeqeq": "warn",
    "import/no-cycle": "error",
    "react/self-closing-comp": [
      "error",
      {
        "html": false
      }
    ]
  },
  "overrides": [
    {
      "files": ["*.test.ts", "*.spec.ts"],
      "rules": {
        "@typescript-eslint/no-explicit-any": "off"
      }
    }
  ]
}
```

## $schema [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#schema)

type: `string | null`

Schema URI for editor tooling.

## categories [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#categories)

type: `object`

Configure an entire category of rules all at once.

Rules enabled or disabled this way will be overwritten by individual rules in the `rules` field.

Example

json

```
{
  "$schema": "./node_modules/oxlint/configuration_schema.json",
  "categories": {
    "correctness": "warn"
  },
  "rules": {
    "eslint/no-unused-vars": "error"
  }
}
```

### categories.correctness [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#categories-correctness)

### categories.nursery [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#categories-nursery)

### categories.pedantic [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#categories-pedantic)

### categories.perf [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#categories-perf)

### categories.restriction [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#categories-restriction)

### categories.style [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#categories-style)

### categories.suspicious [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#categories-suspicious)

## env [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#env)

type: `Record<string, boolean>`

Predefine global variables.

Environments specify what global variables are predefined. See [ESLint's list of environments](https://eslint.org/docs/v8.x/use/configure/language-options#specifying-environments) for what environments are available and what each one provides.

## extends [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#extends)

type: `string[]`

Paths of configuration files that this configuration file extends (inherits from). The files are resolved relative to the location of the configuration file that contains the `extends` property. The configuration files are merged from the first to the last, with the last file overriding the previous ones.

## globals [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#globals)

type: `Record<string, string>`

Add or remove global variables.

For each global variable, set the corresponding value equal to `"writable"` to allow the variable to be overwritten or `"readonly"` to disallow overwriting.

Globals can be disabled by setting their value to `"off"`. For example, in an environment where most Es2015 globals are available but `Promise` is unavailable, you might use this config:

json

```
{
  "$schema": "./node_modules/oxlint/configuration_schema.json",
  "env": {
    "es6": true
  },
  "globals": {
    "Promise": "off"
  }
}
```

You may also use `"readable"` or `false` to represent `"readonly"`, and `"writeable"` or `true` to represent `"writable"`.

## ignorePatterns [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#ignorepatterns)

type: `string[]`

default: `[]`

Globs to ignore during linting. These are resolved from the configuration file path.

## jsPlugins [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#jsplugins)

type: `array | null`

JS plugins.

Note: JS plugins are experimental and not subject to semver. They are not supported in language server at present.

### jsPlugins[n] [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#jsplugins-n)

type: `object | string`

#### jsPlugins[n].name [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#jsplugins-n-name)

type: `string`

Custom name/alias for the plugin.

Note: The following plugin names are reserved because they are implemented natively in Rust within oxlint and cannot be used for JS plugins:

* react (includes react-hooks)
* unicorn
* typescript
* oxc
* import (includes import-x)
* jsdoc
* jest
* vitest
* jsx-a11y
* nextjs
* react-perf
* promise
* node
* regex
* vue
* eslint

If you need to use the JavaScript version of any of these plugins, provide a custom alias to avoid conflicts.

#### jsPlugins[n].specifier [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#jsplugins-n-specifier)

type: `string`

Path or package name of the plugin

## overrides [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#overrides)

type: `array`

### overrides[n] [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#overrides-n)

type: `object`

#### overrides[n].env [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#overrides-n-env)

type: `object | null`

Environments enable and disable collections of global variables.

#### overrides[n].files [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#overrides-n-files)

type: `string[]`

A set of glob patterns.

#### overrides[n].globals [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#overrides-n-globals)

type: `object | null`

Enabled or disabled specific global variables.

#### overrides[n].jsPlugins [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#overrides-n-jsplugins)

type: `array | null`

JS plugins for this override.

Note: JS plugins are experimental and not subject to semver. They are not supported in language server at present.

##### overrides[n].jsPlugins[n] [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#overrides-n-jsplugins-n)

type: `object | string`

###### overrides[n].jsPlugins[n].name [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#overrides-n-jsplugins-n-name)

type: `string`

Custom name/alias for the plugin.

Note: The following plugin names are reserved because they are implemented natively in Rust within oxlint and cannot be used for JS plugins:

* react (includes react-hooks)
* unicorn
* typescript
* oxc
* import (includes import-x)
* jsdoc
* jest
* vitest
* jsx-a11y
* nextjs
* react-perf
* promise
* node
* regex
* vue
* eslint

If you need to use the JavaScript version of any of these plugins, provide a custom alias to avoid conflicts.

###### overrides[n].jsPlugins[n].specifier [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#overrides-n-jsplugins-n-specifier)

type: `string`

Path or package name of the plugin

#### overrides[n].plugins [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#overrides-n-plugins)

type: `array | null`

default: `null`

Optionally change what plugins are enabled for this override. When omitted, the base config's plugins are used.

##### overrides[n].plugins[n] [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#overrides-n-plugins-n)

type: `string`

#### overrides[n].rules [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#overrides-n-rules)

type: `object`

See [Oxlint Rules](https://oxc.rs/docs/guide/usage/linter/rules.html)

## plugins [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#plugins)

type: `array | null`

default: `null`

Enabled built-in plugins for Oxlint. You can view the list of available plugins on [the website](https://oxc.rs/docs/guide/usage/linter/plugins.html#supported-plugins).

NOTE: Setting the `plugins` field will overwrite the base set of plugins. The `plugins` array should reflect all of the plugins you want to use.

### plugins[n] [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#plugins-n)

type: `string`

## rules [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#rules)

type: `object`

See [Oxlint Rules](https://oxc.rs/docs/guide/usage/linter/rules.html)

## settings [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings)

type: `object`

Configure the behavior of linter plugins.

Here's an example if you're using Next.js in a monorepo:

json

```
{
  "settings": {
    "next": {
      "rootDir": "apps/dashboard/"
    },
    "react": {
      "linkComponents": [
        {
          "name": "Link",
          "linkAttribute": "to"
        }
      ]
    },
    "jsx-a11y": {
      "components": {
        "Link": "a",
        "Button": "button"
      }
    }
  }
}
```

### settings.jsdoc [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-jsdoc)

type: `object`

#### settings.jsdoc.augmentsExtendsReplacesDocs [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-jsdoc-augmentsextendsreplacesdocs)

type: `boolean`

default: `false`

Only for `require-(yields|returns|description|example|param|throws)` rule

#### settings.jsdoc.exemptDestructuredRootsFromChecks [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-jsdoc-exemptdestructuredrootsfromchecks)

type: `boolean`

default: `false`

Only for `require-param-type` and `require-param-description` rule

#### settings.jsdoc.ignoreInternal [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-jsdoc-ignoreinternal)

type: `boolean`

default: `false`

For all rules but NOT apply to `empty-tags` rule

#### settings.jsdoc.ignorePrivate [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-jsdoc-ignoreprivate)

type: `boolean`

default: `false`

For all rules but NOT apply to `check-access` and `empty-tags` rule

#### settings.jsdoc.ignoreReplacesDocs [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-jsdoc-ignorereplacesdocs)

type: `boolean`

default: `true`

Only for `require-(yields|returns|description|example|param|throws)` rule

#### settings.jsdoc.implementsReplacesDocs [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-jsdoc-implementsreplacesdocs)

type: `boolean`

default: `false`

Only for `require-(yields|returns|description|example|param|throws)` rule

#### settings.jsdoc.overrideReplacesDocs [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-jsdoc-overridereplacesdocs)

type: `boolean`

default: `true`

Only for `require-(yields|returns|description|example|param|throws)` rule

#### settings.jsdoc.tagNamePreference [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-jsdoc-tagnamepreference)

type: `object`

default: `{}`

### settings.jsx-a11y [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-jsx-a11y)

type: `object`

Configure JSX A11y plugin rules.

See [eslint-plugin-jsx-a11y](https://github.com/jsx-eslint/eslint-plugin-jsx-a11y#configurations)'s configuration for a full reference.

#### settings.jsx-a11y.attributes [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-jsx-a11y-attributes)

type: `Record<string, array>`

default: `{}`

Map of attribute names to their DOM equivalents. This is useful for non-React frameworks that use different attribute names.

Example:

json

```
{
  "settings": {
    "jsx-a11y": {
      "attributes": {
        "for": ["htmlFor", "for"]
      }
    }
  }
}
```

#### settings.jsx-a11y.components [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-jsx-a11y-components)

type: `Record<string, string>`

default: `{}`

To have your custom components be checked as DOM elements, you can provide a mapping of your component names to the DOM element name.

Example:

json

```
{
  "settings": {
    "jsx-a11y": {
      "components": {
        "Link": "a",
        "IconButton": "button"
      }
    }
  }
}
```

#### settings.jsx-a11y.polymorphicPropName [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-jsx-a11y-polymorphicpropname)

type: `string | null`

An optional setting that define the prop your code uses to create polymorphic components. This setting will be used to determine the element type in rules that require semantic context.

For example, if you set the `polymorphicPropName` to `as`, then this element:

jsx

```
<Box as="h3">Hello</Box>
```

Will be treated as an `h3`. If not set, this component will be treated as a `Box`.

### settings.next [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-next)

type: `object`

Configure Next.js plugin rules.

#### settings.next.rootDir [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-next-rootdir)

type: `array | string`

##### settings.next.rootDir[n] [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-next-rootdir-n)

type: `string`

### settings.react [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-react)

type: `object`

Configure React plugin rules.

Derived from [eslint-plugin-react](https://github.com/jsx-eslint/eslint-plugin-react#configuration-legacy-eslintrc-)

#### settings.react.formComponents [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-react-formcomponents)

type: `array`

default: `[]`

Components used as alternatives to `<form>` for forms, such as `<Formik>`.

Example:

jsonc

```
{
  "settings": {
    "react": {
      "formComponents": [
        "CustomForm",
        // OtherForm is considered a form component and has an endpoint attribute
        { "name": "OtherForm", "formAttribute": "endpoint" },
        // allows specifying multiple properties if necessary
        { "name": "Form", "formAttribute": ["registerEndpoint", "loginEndpoint"] },
      ],
    },
  },
}
```

##### settings.react.formComponents[n] [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-react-formcomponents-n)

type: `object | string`

###### settings.react.formComponents[n].attribute [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-react-formcomponents-n-attribute)

type: `string`

###### settings.react.formComponents[n].name [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-react-formcomponents-n-name)

type: `string`

#### settings.react.linkComponents [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-react-linkcomponents)

type: `array`

default: `[]`

Components used as alternatives to `<a>` for linking, such as `<Link>`.

Example:

jsonc

```
{
  "settings": {
    "react": {
      "linkComponents": [
        "HyperLink",
        // Use `linkAttribute` for components that use a different prop name
        // than `href`.
        { "name": "MyLink", "linkAttribute": "to" },
        // allows specifying multiple properties if necessary
        { "name": "Link", "linkAttribute": ["to", "href"] },
      ],
    },
  },
}
```

##### settings.react.linkComponents[n] [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-react-linkcomponents-n)

type: `object | string`

###### settings.react.linkComponents[n].attribute [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-react-linkcomponents-n-attribute)

type: `string`

###### settings.react.linkComponents[n].name [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-react-linkcomponents-n-name)

type: `string`

#### settings.react.version [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-react-version)

type: `string | null`

default: `null`

React version to use for version-specific rules.

Accepts semver versions (e.g., "18.2.0", "17.0").

Example:

jsonc

```
{
  "settings": {
    "react": {
      "version": "18.2.0",
    },
  },
}
```

### settings.vitest [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-vitest)

type: `object`

Configure Vitest plugin rules.

See [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest)'s configuration for a full reference.

#### settings.vitest.typecheck [​](https://oxc.rs/docs/guide/usage/linter/config-file-reference.html#settings-vitest-typecheck)

type: `boolean`

default: `false`

Whether to enable typecheck mode for Vitest rules. When enabled, some rules will skip certain checks for describe blocks to accommodate TypeScript type checking scenarios.
