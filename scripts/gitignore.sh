#!/bin/bash

# Nom du fichier .gitignore
GITIGNORE_FILE=".gitignore"

# Contenu du fichier .gitignore pour les projets Go
GITIGNORE_CONTENT=$(cat <<EOL
# Dossiers
/bin/
/pkg/
/vendor/

# Fichiers exécutables
*.exe
*.out

# Fichiers de build
*.o
*.a
*.so

# Fichiers de logs
*.log

# Fichiers de configuration
*.env

# Fichiers système
.DS_Store

# Fichiers de sauvegarde
*.bak
*.swp

# Dossiers et fichiers spécifiques à Go Modules
/go.sum
/go.mod
/Gopkg.lock

EOL
)

# Vérifie si le fichier .gitignore existe déjà
if [ -e "$GITIGNORE_FILE" ]; then
  echo "$GITIGNORE_FILE existe déjà."
else
  # Crée le fichier .gitignore et ajoute le contenu
  echo "$GITIGNORE_CONTENT" > "$GITIGNORE_FILE"
  echo "Le fichier $GITIGNORE_FILE a été créé avec succès."
fi
