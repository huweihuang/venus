#!/bin/bash
set -ex

PKG="github.com/huweihuang/venus"
GROUP_VERSION='venus:v1'

KUBE_VERSION="v0.23.0"
HOME_DIR=$(pwd)
TOOL_DIR="${HOME_DIR}/tool"
CODEGEN_PKG="${TOOL_DIR}/code-generator"

function download_code_generator() {
    mkdir -p ${TOOL_DIR}
    cd ${TOOL_DIR}
    if [ ! -d "code-generator" ]; then
        git clone https://github.com/kubernetes/code-generator.git --branch ${KUBE_VERSION} 2>/dev/null
    fi
}

function generate_groups(){
    "${CODEGEN_PKG}/generate-groups.sh" "deepcopy,client,informer,lister" \
    ${PKG}/pkg/generated \
    ${PKG}/pkg/apis \
    ${GROUP_VERSION} \
    --go-header-file ${HOME_DIR}/hack/boilerplate.go.txt
}

function main(){
    download_code_generator
    generate_groups
}

main
