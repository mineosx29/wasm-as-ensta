# WebAssembly - Généralité

## Introduction
Dans le cadre du module Application Système du semestre 5 à l'ENSTA Bretagne, plusieurs sujets ont été proposé. Un sujet m'a tout de suite intéressé : Le WebAssembly. J'ai pris ce sujet car c'était un sujet orienté recherche/découverte et la deuxième raison est que je ne connaissais pas du tout le WebAssembly et ma curiosité a pris le dessus. 

Dans ce rapport, je vais vous présenter le WebAssembly. En première partie, je vais présenter les généralité, puis en seconde partie, la proximité du WebAssembly avec les langages de haut niveau puis en troisième partie, vous présenter mon simulateur de machine à état finis et vous expliquer comment il fonctionne et pour finir, une conclusion.

## Qu'est ce que le WebAssembly ?
Le WebAssembly(abrégé WASM) est un nouveau type de code pouvant s'executer sur les navigateur web et peut donner des nouvelles fonctionnalités.
Ce type est principalement développé par le [W3C Community Group](https://www.w3.org/community/webassembly/). Le WebAssembly est principalement intégré dans les navigateur connu d'aujourd'hui : Firefox, Chrome, Opera, Safari.


## A quoi cela sert t-il ?

Le WebAssembly n'est généralement pas utilisé comme langage. Il est principalement utilisé comme cible de compilation à partir des langages de haut niveau choisi par les développeurs. 

Ainsi, les développeurs peuvent développer des programme dans ces langages, les compiler en wasm et les executer sur le navigateur Web.

Cependant, il faut noter que le WebAssembly possède une representation textuelle. Sa représentation textuelle est le WAT(WebAssembly Text).

## Sa représentation Textuelle : le WAT

Pour que le WebAssembly soit lu et edité par les hommes, le WAT, le format textuelle a été développé. 

Sa s-expression: 

La principale unité de code dans le WebAssembly est un module.

```
(module (memory 1) (func))
```
On peut illustrer cela comme un arbre avec des noeuds. Chaque expression entre parenthèse représente un noeud. Dans ce cas, le noeud est le module et les noeuds enfants sont memory avec 1 comme attribut et un autre noeud : func.

Nous pouvons commencer un programme WebAssembly dans le plus simple des cas : 
```
(module)
```

Ensuite, nous pouvons complexifier le module en ajoutant un noeud func : 
```
( func <signature> <locals> <body> )
```
- La signature prend les paramètres de la fonction ainsi que ses types.

- La partie locals prends les variables et ses types.

- La partie body est une suite linéaire d'instruction de bas niveau.

Signature et paramètres : 

| Paramètre | types         |
|-----------|----------------
| i32       | 32bit integer |
| i64       | 64bit integer |
| f32       | 32bit float   |
| f64       | 64bit float   |

## Son format binaire

A la compilation, nous pouvons lire le fichier wasm et on pourra distinguer des données binaire:

<img <img src="images/binaire.jpg" width="600">  >
