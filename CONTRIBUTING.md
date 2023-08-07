# Contributing guidelines

## Create an issue

Please, before creating a pool-request, make sure that you create an issue with the correct tags, which will be a full description of the feature or bug that you want to fix.

## Versioning

At the moment of active development of this tool we support only the current major version. At the moment it is version v2. If you find a bug in the minor version (v2.1, for example), make sure that this problem exists in the latest minor version of boxie (v2.2, for example). After that, create an issue. <br>
Publish your code in the most extreme minor branch of the current major version. It will go into the main as soon as all pooled requesters are closed.

## Vision

We try to make boxie as high-level and versatile as possible. Complex cluster configurations and other settings should be optional.

## Tests and quality

Please cover your code with unit tests. And also write comments on everything you add. Codacy linter is available for all release branches.