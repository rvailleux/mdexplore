#!/bin/bash
# bump-version.sh - Increment version after PR merge
# Usage: ./scripts/bump-version.sh [major|minor|patch]

set -e

# Default bump type
BUMP_TYPE="${1:-patch}"

# Get current version from git tags
CURRENT_VERSION=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")

# Remove 'v' prefix
VERSION_NUM=${CURRENT_VERSION#v}

# Split version into components
IFS='.' read -r MAJOR MINOR PATCH <<< "$VERSION_NUM"

# Bump version
case "$BUMP_TYPE" in
    major)
        MAJOR=$((MAJOR + 1))
        MINOR=0
        PATCH=0
        ;;
    minor)
        MINOR=$((MINOR + 1))
        PATCH=0
        ;;
    patch)
        PATCH=$((PATCH + 1))
        ;;
    *)
        echo "Error: Invalid bump type '$BUMP_TYPE'. Use major, minor, or patch."
        exit 1
        ;;
esac

NEW_VERSION="v${MAJOR}.${MINOR}.${PATCH}"

echo "Bumping version: $CURRENT_VERSION -> $NEW_VERSION"

# Create git tag
git tag -a "$NEW_VERSION" -m "Release $NEW_VERSION"

echo ""
echo "Created tag: $NEW_VERSION"
echo "To push the tag and trigger release:"
echo "  git push origin $NEW_VERSION"
