#!/bin/bash

# Obtenir le nom du répertoire actuel (nom du projet Git)
project_name=$(basename "$(git rev-parse --show-toplevel)")

git config --global user.email "jeangabrielgoudiaby@gmail.com"
git config --global user.name "jgoudiab"

while true; do
    current_time=$(date "+%Y-%m-%d %H:%M:%S")
    echo "Votre dernier commit était à : $current_time dans le projet : $project_name"

    # Vérifier s'il y a des modifications
    if [ -n "$(git status --porcelain)" ]; then
        git add .
        git commit -m "$current_time $1"
        git push

    # Afficher les derniers commits
    echo "Derniers commits:"
    git log --pretty=format:"%h - %an, %ar : %s" -n 5
    else
        echo "Aucune modification détectée. Pas de push."
    fi

    sleep 1500
done
# to run :
# ./push.sh > mon_fichier.log 2> erreurs.log

