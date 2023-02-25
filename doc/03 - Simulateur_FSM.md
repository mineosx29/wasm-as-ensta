# Simulateur FSM(Finite State Machine)


## Qu'est une FSM ?
Une machine à état finis est une machine qui est susceptible d'être dans un nombre fini d'état mais étant un moment donné que dans un état à la fois. L'état dans lequel il se trouve se nomme l'état courant. Le passage d'un état à un autre est activé par un évenement ou une condition.

On rencontre de nombreuses FSM dans la vie quotidienne : un distributeur automatique, un ascenseur, un digicode etc...

Voici un exemple de machine à état finis : 


<img src="images/FSM.png">

# Application : Simulateur de FSM

Dans le cadre du module Application Système, j'ai programmé un simulateur de machine à État Finis dans le but de montrer les possibilités du WebAssembly. En effet, avec l'accord de mon professeur Superviseur, j'ai décidé d'inventer un langage automate qui décrira une FSM. Ensuite, un compilateur que j'ai construit s'occupera de le compiler. Enfin à la sortie du compilateur, un diagramme automate sera crée ainsi qu'un programme en GO qui décrit la FSM.

Voici la chaîne de fonctionnement : 

<img src="images/compil.png">

## Le langage DSL_Auto : Syntaxe
J'ai décidé de rapprocher la grammaire et la syntaxe du langage des langages de haut niveau pour faciliter l'apprentissage et la prise en main du langage.

La syntaxe du langage DSL_Auto est la suivante : 
```
begin <nom_automate> :
    input = <entrées1>;<entrées2>
    state = <etat_initial>,<etat_1>,<etat_2>

    case <etat>:
        <condition>
    end case
end

```

le bloc condition est le suivant : 
```
if <input> then 
    next_state=etat_suivant
```

## Grammaire du langage

Auto -> "begin" identifier : {Declaration}   
Declaration -> Type "input" "=" identifier {"; identifier"} Type "state" "=" identifier {"," identifier}   
Type -> input | state  
Etats -> {Etats}  
Etats -> "case" identifier ":" {conditions} 
conditions -> if "identifier" then next_state=identifier   

Nous avons vu la syntaxe et la grammaire du langage DSL_Auto. Le langage DSL_Auto va être compilé et à partir de ce langage va être généré le diagramme automate ainsi que le programme GO décrivant la FSM.

## Fonctionnement du compilateur
Le compilateur DSL_Auto se décompose en trois partie : 
- Un lexer : cette partie va s’occuper de découper le code source en lexèmes.  
- Un parser : elle va vérifier que la suite de lexèmes obtenu grâce au lexer correspond à la grammaire de notre langage source. Cette partie va également concevoir l’arbre de syntaxe abstraite.
- Un visiteur : C’est notre design pattern. Il va principalement visiter l’arbre de syntaxe abstraite. Le visiteur est directement intégré dans le générateur de code.

Tout d'abord, il y a l'analyse lexical :    
En effet, une expression régulière est une chaine de caractère qui décrit, selon une syntaxe précise, un ensemble de chaîne de caractère possible.
L’analyse lexical sert à découper le code source en mots ou lexème et ensuite vérifier avec les expressions régulières que ces mots correspondent au dictionnaire du langage source.    
Voici ci-dessous le dictionnaire de mots avec leurs expressions régulière correspondantes :

```Python
reg_expression = [
    (r'[ \n\t]+', None),
    (r'#[^\n]*', None),
    (r'case\b', 'CASE'),
    (r'if\b', 'IF'),
    (r'begin\b', 'BEGIN'),
    (r'state\b', 'STATE'),
    (r'not\b', 'NOT'),
    (r'next_state\b', 'NEXTST'),
    (r'input\b', 'INPUT'),
    (r'state\b', 'STATE'),
    (r'then\b', 'THEN'),
    (r'end\b', 'END'),
    (r'\(', 'LPAREN'),
    (r'\)', 'RPAREN'),
    (r'\{', 'LBRACE'),
    (r'\}', 'RBRACE'),
    (r'\,', 'COMMA'),
    (r'\:', 'COLON'),
    (r'\;', 'SEMICOLON'),
    (r'\=', 'ASSIGN'),
    (r'and', 'AND'),
    (r'or', 'OR'),
    (r'[a-zA-Z][a-zA-Z0-9]*', 'IDENTIFIER')
    
]
```
Ci-dessus, nous définissons une expression régulière pour chaque mot clé du langage de programmation.

Ensuite le lexer va séparer le langage source en lexèmes et voir si chaque lexème correspond à une des expressions régulières définit ci-dessus. C’est qu’on appelle l’analyse lexicale.

Voici le diagramme UML du lexer : 
<img src="images/lexerClass.png">

Ensuite viens l'étape de l'analyse syntaxique : 

L’analyse syntaxique dans un compilateur est une étape où le compilateur va vérifier que l’enchaînement des lexèmes correspond bien à la grammaire du langage. Ce sera le rôle de notre analyseur syntaxique ou parser.

Voici le diagramme UML de l'analyseur syntaxique : 

<img src="images/parserClass.png">
L’analyseur syntaxique va viser deux objectifs dans notre compilateur :
    • Vérifier que le code source respecte la grammaire préétablie du langage.
    • Construire un arbre de syntaxe abstraite.
Ici, nous utilisons dans notre analyseur syntaxique une analyse récursive descendante et nous utilisons généralement trois méthodes dans notre analyseur syntaxique et ces méthodes sont implémenté dans les méthodes parse_X également :    

- La méthode expect() : elle va servir à consommer un lexème de type attendu.
- La méthode acceptit() : elle va consommer le lexème courant.
- La méthode showNext() : elle va retourner le lexème courant.   

La méthode parse définit dans le diagramme de classe sont des méthodes où on définit la grammaire du langage soit ses règles.
Ainsi pour chaque déclaration de variable, pour chaque boucle conditionnel et conditions, le parser va vérifier via les méthodes parse_X() que chaque enchaînements de lexèmes du code source respectent bien la grammaire du langage.
