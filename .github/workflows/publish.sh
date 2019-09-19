#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

readonly version=$(cat VERSION)
readonly git_branch=${1:11} # drop 'refs/head/' prefix
readonly git_sha=$(git rev-parse HEAD)
readonly git_timestamp=$(TZ=UTC git show --quiet --date='format-local:%Y%m%d%H%M%S' --format="%cd")
readonly slug=${version}-${git_timestamp}-${git_sha:0:16}

echo "Publishing riff System"
gsutil cp -a public-read gs://projectriff/riff-system/snapshots/riff-build-${slug}.yaml gs://projectriff/riff-system/riff-build-${version}.yaml
gsutil cp -a public-read gs://projectriff/riff-system/snapshots/riff-core-${slug}.yaml gs://projectriff/riff-system/riff-core-${version}.yaml
gsutil cp -a public-read gs://projectriff/riff-system/snapshots/riff-knative-${slug}.yaml gs://projectriff/riff-system/riff-knative-${version}.yaml
gsutil cp -a public-read gs://projectriff/riff-system/snapshots/riff-streaming-${slug}.yaml gs://projectriff/riff-system/riff-streaming-${version}.yaml

echo "Publishing version references"
gsutil -h 'Content-Type: text/plain' -h 'Cache-Control: private' cp -a public-read <(echo "${slug}") gs://projectriff/riff-system/snapshots/versions/${git_branch}
gsutil -h 'Content-Type: text/plain' -h 'Cache-Control: private' cp -a public-read <(echo "${slug}") gs://projectriff/riff-system/snapshots/versions/${version}
if [[ ${version} != *"-snapshot" ]] ; then
gsutil -h 'Content-Type: text/plain' -h 'Cache-Control: private' cp -a public-read <(echo "${version}") gs://projectriff/riff-system/versions/releases/${git_branch}
fi