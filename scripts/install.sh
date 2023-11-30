#!/bin/bash

extensions=(
    "Blackboxapp.blackboxauto"
    "mintlify.document"
    "ritwickdey.LiveServer"
    "fabiospampinato.vscode-todo-plus"
    "shakram02.bash-beautify"
    "rvest.vs-code-prettier-eslint"
    "MS-vsliveshare.vsliveshare"
    "dracula-theme.theme-dracula"
    "usernamehw.errorlens"
    "jinliming2.vscode-go-template"
    "eamodio.gitlens"
    "vscode-icons-team.vscode-icons"
    
)

for extension in "${extensions[@]}"
do
    code --install-extension "$extension"
done
