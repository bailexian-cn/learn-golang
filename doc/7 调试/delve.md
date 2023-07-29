# 本地调试

```shell
dlv exec -- ./e2e -e2e.artifacts-folder="/home/cnv/cluster-api/_artifacts" -e2e.config="/home/cnv/cluster-api/test/e2e/config/docker.yaml" -e2e.skip-resource-cleanup=false -e2e.use-existing-cluster=false

(dlv) b /root/minio/cmd/server-main.go:531
(dlv) b /root/cluster-api-provider-cke/e2e/create_kubevirtmachinetemplates_test.go:72
(dlv) c
```

