# aero

## 🚀 Getting Started
```shell
goreleaser init

goreleaser release --snapshot --skip-publish --rm-dist

goreleaser build --single-target --snapshot --rm-dist

export GITHUB_TOKEN=""; 

git tag -a v1.0.0 -m "First release" && git push origin v1.0.0

 brew tap donwany/aero https://github.com/donwany/aero 

brew install aero

aero weather --city Dallas --apikey $APIKEY

```
