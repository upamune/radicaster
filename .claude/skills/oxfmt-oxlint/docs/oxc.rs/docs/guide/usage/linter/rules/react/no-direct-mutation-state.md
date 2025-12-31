---
title: "react/no-direct-mutation-state | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/no-direct-mutation-state"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/no-direct-mutation-state.md for this page in Markdown format

# react/no-direct-mutation-state Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-direct-mutation-state.html#react-no-direct-mutation-state)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-direct-mutation-state.html#what-it-does)

The restriction coder cannot directly change the value of this.state

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-direct-mutation-state.html#why-is-this-bad)

calling setState() afterwards may replace the mutation you made

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-direct-mutation-state.html#examples)

jsx

```
 // error
 var Hello = createReactClass({
   componentDidMount: function() {
     this.state.name = this.props.name.toUpperCase();
   },
   render: function() {
     return <div>Hello {this.state.name}</div>;
   }
 });

 class Hello extends React.Component {
   constructor(props) {
     super(props)

     doSomethingAsync(() => {
       this.state = 'bad';
     });
   }
 }

 // success
 var Hello = createReactClass({
   componentDidMount: function() {
     this.setState({
       name: this.props.name.toUpperCase();
     });
   },
   render: function() {
     return <div>Hello {this.state.name}</div>;
   }
 });

 class Hello extends React.Component {
   constructor(props) {
     super(props)

     this.state = {
       foo: 'bar',
     }
   }
 }
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-direct-mutation-state.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/no-direct-mutation-state": "error"
  }
}
```

bash

```
oxlint --deny react/no-direct-mutation-state --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-direct-mutation-state.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/no_direct_mutation_state.rs)
