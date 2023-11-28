#!/bin/bash

create() {
    echo "Press any key to create folder or 's' to skip..."
    while true; do
        read -t 3 -n 1 input
        if [ $? = 0 ]; then
            if [ "$input" = "s" ]; then
                echo "Skipping current step..."
                break
            else
                clear
                break
            fi
        fi
    done
}
# Nom du projet
PROJECT_NAME="$1"
if [ -z "$PROJECT_NAME" ]; then
    echo "Le nom du projet est requis."
    exit 1
fi

# Créer la structure du répertoire
echo "Création de la structure du projet Go : $PROJECT_NAME"
create
go mod init $PROJECT_NAME
create
mkdir -p "cmd"
create
touch "cmd/main.go"
create
mkdir -p "docs"
create
mkdir -p "pkg/internal/api/handlers"
create
touch "pkg/internal/api/handlers/handlers.go"
create
mkdir -p "pkg/internal/api/models"
create
touch "pkg/internal/api/models/errors.go"
create
touch "pkg/internal/api/models/models.go"
create
mkdir -p "pkg/internal/api/server"
create
touch "pkg/internal/api/server/server.go"
create
mkdir -p "pkg/internal/utils"
create
touch "pkg/internal/utils/utils.go"
create
mkdir -p "pkg/web/static/css"
create
touch "pkg/web/static/css/style.css"
create
mkdir -p "pkg/web/assets/img"
create
mkdir -p "pkg/web/templates"
create
touch "pkg/web/templates/index.html"
create
mkdir -p "scripts"
create
mkdir -p "tests" 
create
touch ".gitignore"
echo "La structure du projet a été créée dans le répertoire courant."
create
code .