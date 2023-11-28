#!/bin/bash

git config --global credential.helper store
git config --global user.email "jeangabrielgoudiaby@gmail.com"
git config --global user.name "jgoudiab"
# Vérifier si le nom du projet est fourni en argument
if [ -z "$1" ]; then
  echo "Utilisation : $0 <nom-du-projet>"
  exit 1
fi

# Récupérer le nom du projet à partir du premier argument
project_name="$1"

# URL du référentiel Git à cloner
repository_url=https://learn.zone01dakar.sn/git/jgoudiab/$project_name.git

# Cloner le référentiel Git dans le répertoire actuel
git clone "$repository_url"

# Accéder au répertoire du projet cloné
cd "$project_name"

# Configuration des télécommandes (origin par défaut)
git remote -v

# Mettre à jour les sous-modules Go si le projet en contient
git submodule update --init --recursive

# Vérifier l'état du référentiel
git status

# Informations de base sur le référentiel cloné
echo "Référentiel $project_name cloné avec succès dans le répertoire actuel."

code .
