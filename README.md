# venus
Venus is a k8s operator development framework

# Usage

## 1. clone repo

```shell
git clone https://github.com/huweihuang/venus.git
```

## 2. update apis type file

Modify the directory name, usually using your project name.

```shell
pkg/apis
└── venus  # update the dir
    └── v1
        ├── doc.go
        ├── memcache_types.go
        ├── redis_types.go
        ├── register.go
        └── zz_generated.deepcopy.go  # The file is automatically generated and do not edit
```

Rename the name of the CRD and add parameters for spec and status.

You can add multiple types of CRDs.

```go
// RedisSpec is the spec for a Redis resource
type RedisSpec struct {
	DeploymentName string `json:"deploymentName"`
	Replicas       *int32 `json:"replicas"`
}

// RedisStatus is the status for a Redis resource
type RedisStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}
```

## 3. run codegen script

edit update-codegen.sh and modify parameter `PKG`, `GROUP_VERSION`.

```shell
PKG="github.com/huweihuang/venus"
GROUP_VERSION='venus:v1'

...
```

run make codegen

```shell
go mod tidy
make codegen
```

It will automatically generate the following code.

- clientset
- informers
- listers

```shell
pkg/generated
├── clientset
│ └── versioned
│     ├── clientset.go
│     ├── doc.go
│     ├── fake
│     │ ├── clientset_generated.go
│     │ ├── doc.go
│     │ └── register.go
│     ├── scheme
│     │ ├── doc.go
│     │ └── register.go
│     └── typed
│         └── venus
│             └── v1
│                 ├── doc.go
│                 ├── fake
│                 │ ├── doc.go
│                 │ ├── fake_memcache.go
│                 │ ├── fake_redis.go
│                 │ └── fake_venus_client.go
│                 ├── generated_expansion.go
│                 ├── memcache.go
│                 ├── redis.go
│                 └── venus_client.go
├── informers
│ └── externalversions
│     ├── factory.go
│     ├── generic.go
│     ├── internalinterfaces
│     │ └── factory_interfaces.go
│     └── venus
│         ├── interface.go
│         └── v1
│             ├── interface.go
│             ├── memcache.go
│             └── redis.go
└── listers
    └── venus
        └── v1
            ├── expansion_generated.go
            ├── memcache.go
            └── redis.go
```

