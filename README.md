# tectonic_assets
- cli for flutter assets management

## installation
```bash
git clone https://github.com/NjiruClinton/tectonic_assets.git
```
```bash
go install
```
a package `tinypng` will need you to set your tinypng api key
`export TINY_PNG_API_KEY=your_api_key` for UNIX based systems and `set TINY_PNG_API_KEY=your_api_key` for windows

```bash
go build
```

## usage
> run the gfg command created on cmd inside the flutter project directory
```bash
gfg --help
```
* Available Commands:
  * `browse` *Browse assets folder*
  * `completion` *Generate the autocompletion script for the specified shell*
  * `help` *Help about any command*
  * `manifest` *Generate asset manifest*
  * `optimize` *Optimize image assets*

### example
```bash
gfg browse
```
outputs fonts, images and other assets from a flutter project

### test directories
- [x] pubspec.yaml
- [x] assets/

**LICENCE: MIT**


> project still under development