# Partiel CTF Golang

Ce projet permet de résoudre le Capture The Flag (CTF) en utilisant Go. Le code présenté ici est conçu pour interagir avec l'API du serveur pour le CTF en effectuant des requêtes HTTP pour résoudre les différents défis proposés.

## Fonctionnalités

Ce projet présente les fonctionnalités suivantes :

- Teste la disponibilité des ports sur le serveur CTF.
- Effectue des requêtes HTTP POST et GET pour résoudre les défis CTF.
- Affiche les réponses et les résultats des requêtes pour chaque port testé.

## Exigences

- Go (Golang) installé sur votre système. Vous pouvez le télécharger à partir de [golang.org](https://golang.org/dl/).

## Utilisation

1. Clonez ce dépôt vers votre système local :

   ```bash
   git clone https://github.com/hmzaakun/goCTF.git
   cd goCTF```
   
2. Modifiez le fichier main.go avec les informations appropriées pour votre CTF, notamment l'adresse IP du serveur CTF, les ports à tester et les données utilisateur.

3. Exécutez le programme Go :

```bash
go run main.go
```
Le programme testera la disponibilité des ports spécifiés sur le serveur CTF et effectuera des requêtes HTTP pour résoudre les défis.

## Configuration

- serverIP: Remplacez la valeur par l'adresse IP du serveur de l'API.
- minPort et maxPort: Spécifiez la plage de ports à tester.
- Les données utilisateur et les demandes de défis doivent être adaptées à votre CTF.

By Hamza :shipit:
