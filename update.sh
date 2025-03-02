#!/bin/bash

git clone https://github.com/lucide-icons/lucide.git ./lucide-git
cd ./lucide-git
git checkout $(git describe --tags $(git rev-list --tags --max-count=1))
git describe --tags $(git rev-list --tags --max-count=1) > ../luicde-version.txt

mv ./icons/*.svg ../icons/

#convert from thing-test.svg to ThingTest.svg removing the - and capitalizing the next letter
cd ../icons
for f in *.svg; do
    new_name=$(echo "$f" | sed -E 's/-(.)/\U\1/g' | sed -E 's/^(.)/\U\1/')
    mv -- "$f" "$new_name"
done

cd ..
rm -rf ./lucide-git

echo "updating icons completed"