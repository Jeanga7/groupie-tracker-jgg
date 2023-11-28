#!/bin/bash

# Vérifiez si Doxygen est installé
if ! command -v doxygen &> /dev/null; then
    echo "Doxygen n'est pas installé. Veuillez l'installer pour utiliser ce script."
    exit 1
fi

# Spécifiez le répertoire source de votre code
source_dir="votre_repertoire_source"

# Spécifiez le répertoire de destination de la documentation générée
output_dir="votre_repertoire_de_documentation"

# Créez un fichier de configuration Doxygen s'il n'existe pas
if [ ! -f "Doxyfile" ]; then
    doxygen -g Doxyfile
fi

# Modifiez le fichier de configuration Doxygen selon vos besoins
# Par exemple, vous pouvez personnaliser les paramètres dans le fichier Doxyfile

# Exécutez Doxygen avec le fichier de configuration
doxygen Doxyfile

echo "La documentation a été générée avec succès dans : $output_dir"
