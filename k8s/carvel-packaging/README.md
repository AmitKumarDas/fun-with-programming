## Carvel Packaging the Hard Way

### Motivation
Convert following Carvel documentation into reproducible code:
- https://carvel.dev/kapp-controller/docs/v0.40.0/packaging-tutorial/

## Design
- Shell script `run.sh` is the single source of truth
- Rest of the files & folders are generated from `run.sh`

## Run
- `sh run.sh`

## References
- https://github.com/vmware-tanzu/carvel-kapp-controller/tree/develop/examples/local-kind-environment
