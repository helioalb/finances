#!/usr/bin/env bash
set -Eeuo pipefail
IFS=$'\n\t'

SCRIPT_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd -- "$SCRIPT_DIR/.." && pwd)"
DOCKERFILE="$REPO_ROOT/deployments/Dockerfile"

if ! command -v docker >/dev/null 2>&1; then
	echo "[ERROR] docker nao encontrado no PATH." >&2
	exit 1
fi

if [[ ! -f "$DOCKERFILE" ]]; then
	echo "[ERROR] Dockerfile nao encontrado em: $DOCKERFILE" >&2
	exit 1
fi

VERSION="${VERSION:-latest}"
BUILD_DATE="${BUILD_DATE:-$(date -u +"%Y-%m-%dT%H:%M:%SZ")}"
IMAGE_NAME="${IMAGE_NAME:-helioalb/finances}"

if git -C "$REPO_ROOT" rev-parse --git-dir >/dev/null 2>&1; then
	GIT_COMMIT="${GIT_COMMIT:-$(git -C "$REPO_ROOT" rev-parse --short HEAD)}"
else
	GIT_COMMIT="${GIT_COMMIT:-unknown}"
fi

IMAGE_TAG="$IMAGE_NAME:$VERSION"

echo "[INFO] Building image: $IMAGE_TAG"
echo "[INFO] Build args: VERSION=$VERSION BUILD_DATE=$BUILD_DATE GIT_COMMIT=$GIT_COMMIT"

docker build \
	--file "$DOCKERFILE" \
	--tag "$IMAGE_TAG" \
	--build-arg "VERSION=$VERSION" \
	--build-arg "BUILD_DATE=$BUILD_DATE" \
	--build-arg "GIT_COMMIT=$GIT_COMMIT" \
	"$REPO_ROOT"
