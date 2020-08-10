# Hello World (tutorial)
Hello World is a tutorial for building plugins for `vroomy`. Please follow along step-by-step using the project branches.

## What is a plugin?
A plugin is a module used by vroomy to serve HTTP endpoints. A plugin can be thought of as a self contained service. A plugin can be completely independent, or it can leverage other plugins running within the vroomy instance as well. Typically, a plugin will leverage one or more `controllers` in order to serve it's required endpoints.

Vroomy plugins are `controller-agnostic`. There isn't a single enforced `controller` type or framework to be utilized. In the case of this tutorial, we will be utilizing the `gdbu/hello-world` example for our `controller`.

## Why would I want to utilize a plugin?
Plugins provide a huge performance advantage over `proxy` and `FastCGI` services. The reason for this difference in performance is due to `vroomy` plugins running directly on the `vroomy` server. With `proxy` and `FastCGI`, all data must incur an additional network cost when the request is forwarded.
