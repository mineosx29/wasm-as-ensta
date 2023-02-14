> [ZIANI Nadir](https://github.com/mineosx29) FIPA SE 2023 
## Introduction Générale
Dans le cadre du module Application Système, j'ai choisi le sujet : Découverte du [WebAssembly](https://webassembly.org/)
Cette application sert principalement à démontrer les possibilités de développement qu'on peut effectuer avec WASM.

## Pré-Requis
Cloner le repositorie Github: 
```
git clone https://github.com/mineosx29/wasm-as-ensta.git
```
Installer Go sur Linux (Ubuntu) : 
```
sudo apt-get install golang-go
```
Ensuite aller à la racine du projet et dans le répertoire assets : 
```
cd assets
```
Puis faire un make des Machine à Etats Finis que vous voulez simuler : 

```
make robot
make lampe
...
```

ensuite aller dans le dossier cmd/server/ et lancer la commande : 

```
go run main.go
```

Aller dans votre navigateur, et aller à l'adresse : 
http://localhost:9090 pour lancer le simulateur