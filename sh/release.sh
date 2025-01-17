#!/bin/bash

export GOOS=linux

BUILD_PATH="build/release"
VERSION_DIR="metadata/version"
VERSION_FILE="$VERSION_DIR/.env"

[ -d "$VERSION_DIR" ] || mkdir -p "$VERSION_DIR"
[ -f "$VERSION_FILE" ] || touch "$VERSION_FILE"

source "$VERSION_FILE"

CURRENT_VERSION=$(echo $LAST_VERSION)

if [ -z "$CURRENT_VERSION" ]; then
  NEW_VERSION="0.0.1"
else
  echo "Select breaking level change:"
  options=("Low" "Medium" "High")
  selected=0

  while true; do
    for i in "${!options[@]}"; do
      if [ $i -eq $selected ]; then
        echo -e "\e[1;32m> ${options[$i]}\e[0m"
      else
        echo "  ${options[$i]}"
      fi
    done

    read -rsn1 input
    case $input in
      $'\x1b')
        read -rsn2 -t 0.1 input
        if [ "$input" == "[A" ]; then
          ((selected--))
          if [ $selected -lt 0 ]; then
            selected=$((${#options[@]} - 1))
          fi
        elif [ "$input" == "[B" ]; then
          ((selected++))
          if [ $selected -ge ${#options[@]} ]; then
            selected=0
          fi
        fi
        ;;
      "")
        character=${options[$selected]}
        break
        ;;
    esac

    for ((i=0; i<${#options[@]}; i++)); do
      echo -en "\033[1A\033[2K"
    done
  done

  echo "Selected character: $character"

  IFS='.' read -r -a version_parts <<< "$CURRENT_VERSION"
  major=${version_parts[0]}
  minor=${version_parts[1]}
  patch=${version_parts[2]}
  case $character in
    High)
      major=$((major + 1))
      minor=0
      patch=0
      ;;
    Medium)
      minor=$((minor + 1))
      patch=0
      ;;
    Low)
      patch=$((patch + 1))
      ;;
  esac
  NEW_VERSION="$major.$minor.$patch"
fi

echo "LAST_VERSION=$NEW_VERSION" > "$VERSION_FILE"

rm -rf "$BUILD_PATH"
mkdir "$BUILD_PATH"

go build -o "$BUILD_PATH/main" cmd/*.go

#? Git tags and Docker tags

git add .
git commit -m "Releasing version v$NEW_VERSION"
git tag -a "v$NEW_VERSION" -m "Releasing version v$NEW_VERSION"
git push origin "v$NEW_VERSION"


