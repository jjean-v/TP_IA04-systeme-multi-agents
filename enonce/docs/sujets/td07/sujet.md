---
tags: TD, dynamique des croyances, fusion
---

# [IA04] Sujet TD7


## Exercice 1 – modélisation logique des croyances (examen final A22)

On se place dans le cadre de la logique propositionnelle basée sur l'ensemble de variables propositionnelles $`\{a,b,c\}`$ avec les 2 valeurs de vérité $`\{0,1\}`$. L'ensemble des interprétations induit est noté $`\Omega`$.

On considère $`\varphi`$ une formule logique tel que 

$$\varphi = (a \vee b \vee c)\wedge(\neg a \vee \neg b)\wedge(\neg a \vee \neg c)\wedge (\neg b \vee \neg c)$$

et une nouvelle information $$\mu = a\wedge b$$

### Questions

1. Reproduire et remplir la table de vérité suivante. Pour rappel, $`val(\varphi,\omega)`$ représente la valuation de la formule $`\varphi`$ par l'interprétation $`\omega`$.

| $`\omega_i`$  | $`a`$  | $`b`$  | $`c`$  | $`val(\varphi,\omega_i)`$  | $`val(\mu,\omega_i)`$  |
|---|---|---|---|---|---|
| $`\omega_0`$ | 0 | 0 | 0 |   |  |
| $`\omega_1`$ | 0 | 0 | 1 |   |  |
| $`\omega_2`$ | 0 | 1 | 0 |   |  |
| $`\omega_3`$ | 0 | 1 | 1 |   |  |
| $`\omega_4`$ | 1 | 0 | 0 |   |  |
| $`\omega_5`$ | 1 | 0 | 1 |   |  |
| $`\omega_6`$ | 1 | 1 | 0 |   |  |
| $`\omega_7`$ | 1 | 1 | 1 |   |  |

2. Déduire de la question précédente les modèles de $`\varphi`$ et de $`\mu`$.

On considère l'opérateur de changement de croyance de Dalal $`\circ_D`$. Pour rappel, le résultat de la révision $`\circ_D`$ est une formule dont les modèles de $`\mu`$ minimisent la distance de Hamming $`d_H`$ par rapport à la base $`\varphi`$.

3. Rappeler ce qu'est la distance de Hamming.
4. Donner le résultat de $`\varphi\circ_D\mu`$ sous forme d'un ensemble d'interprétation puis d'une formule logique.
5. Interpréter $\varphi$, $`\mu`$ et  $`\varphi\circ_D\mu`$ en terme d'ananas, de bananes et de noix de coco. Quelle différence avec un opérateur de mise-à-jour ?


## Exercice 2 – fusion et vaccins (examen final A21)

4 experts médicaux discutent de l'efficacité de 3 vaccins. Considérons les variables propositionnelles $`\{A,B,C\}`$ signifiant « le vaccin A (resp. B, C) est efficace. »

- L'expert n°1 et l'expert n°2 pensent que les vaccins A et B sont efficaces (mais ne se prononce pas pour le dernier). 
- L'expert n°3 estime que le vaccin B est efficace, mais pas le vaccin A. 
- L'expert n°4 estime que si le vaccin B est efficace, alors le vaccin C l'est également (N.B. si le vaccin B n'est pas efficace, il n'a aucune idée de l'efficacité du vaccin C).

### Questions

1. Représenter les croyances de chaque agent $`j`$ à l'aide d'une formule propositionnelle $`\varphi_j`$. Quels sont les modèles de chacune de ces formules (cf. tableau ci-après) ?
3. Qu'est-ce la distance de Hamming entre une interprétation et une formule ? Comment la calculer ?
4. Remplir le tableau suivant, sachant que $`d_H^j(\varphi_j, \omega_i)`$ représente la distance de Hamming entre l'interprétation $`ω_i`$ et $`\varphi_j`$.
5. Quel est le résultat de la fusion de ces 4 croyances en prenant comme opérateur d'agrégation la somme, la somme des carrés, le $`\max`$ ou le $`\text{GMax}`$ ?
6. Après une série massive de tests, on est maintenant sûr que le vaccin C n'est pas efficace. Comment modéliser cette contrainte d'intégrité, notée $`\mu`$ à l'aide d'une formule propositionnelle ?
7. Que devient le résultat de la fusion si l'on prend comme contrainte d'intégrité $`\mu`$ ?

| $`\Omega`$ | $`A`$ | $`B`$ | $`C`$ | $`d_H^1(\varphi_1, \omega_i)`$  | $`d_H^2(\varphi_2, \omega_i)`$ | $`d_H^3(\varphi_3, \omega_i)`$  | $`d_H^4(\varphi_4, \omega_i)`$  | $`\displaystyle\sum_{j=1}^4 d_H^j(\varphi_j, \omega_i)`$  | $`\displaystyle\sum_{j=1}^4 d_H^j(\varphi_j, \omega_i)^2`$ | $`\displaystyle\max_{j=1..4} d_H^j(\varphi_j, \omega_i)`$ | $`\displaystyle\underset{j=1..4}{\text{GMax}}\ d_H^j(\varphi_j, \omega_i)`$ |
|---|---|---|---|---|---|---|---|---|---|---|---|
| $`ω_0`$ | 0 | 0 | 0 |   |   |   |   |   |   |   |   |
| $`ω_1`$ | 0 | 0 | 1 |   |   |   |   |   |   |   |   |
| $`ω_2`$ | 0 | 1 | 0 |   |   |   |   |   |   |   |   |
| $`ω_3`$ | 0 | 1 | 1 |   |   |   |   |   |   |   |   |
| $`ω_4`$ | 1 | 0 | 0 |   |   |   |   |   |   |   |   |
| $`ω_5`$ | 1 | 0 | 1 |   |   |   |   |   |   |   |   |
| $`ω_6`$ | 1 | 1 | 0 |   |   |   |   |   |   |   |   |
| $`ω_7`$ | 1 | 1 | 1 |   |   |   |   |   |   |   |   |
