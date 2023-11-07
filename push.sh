#!/bin/bash
git config --global credential.helper store
git config --global user.email "jeangabrielgoudiaby@gmail.com"
git config --global user.name "jgoudiab"

git add .
git commit -m "$1"
git push
