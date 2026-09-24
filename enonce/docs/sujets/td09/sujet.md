---
tags: TD
---

# [IA04 / AI30] Sujet de TD n°9 : Négociation et enchères

## Exercice 1 : Jeu séquentiel, équilibre de Nash, équilibre parfait

On considère le jeu séquentiel à deux joueurs (I et II) représenté par l'arbre suivant :
![](https://i.imgur.com/4qcwtyX.png)

1. Énumérer toutes les **stratégies** pures de chacun des joueurs
2. Représenter ce jeu sous **forme normale** à l'aide d'une bimatrice de gain
3. Déterminer les correspondances de **meilleure réponse** de chacun des joueurs, puis tous les **équilibres de Nash** en stratégies pures et la valeur de chacun
4. Déterminer tous les **équilibres parfaits** en sous-jeu en stratégies pures et la valeur de chacun. Commenter.

## Exercice 2 : Analyse stratégique de la négociation

On se place dans le cadre vu en cours :
* Les agents A et B doivent se partager un gâteau de taille 1.
* L'agent A joue en premier
* L'agent dont c'est le tour propose une offre de partage, de la forme $o_1 = (x,1-x)$ avec $x\in[0,1]$
* L'autre agent peut **accepter**, et le jeu se termine, ou bien **refuser** et prendre la main au tour suivant
* *tie-breaking* : lorsqu'un agent reçoit une offre, s'il est indifférent entre acceptation et refus, alors il accepte cette offre

### A. Modèle à coûts fixes

Dans ce modèle, si l'offre $o_t=(x,1-x)$ est acceptée, les agents reçoivent les gains $g_A = x-(t-1) \ c_A$, $g_B = 1-x-(t-1) \ c_B$.
On considère le jeu à **horizon fini**. 
1. Analyser les jeux en 1, 2, 3 tours en précisant les équilibres parfaits.

### B. Modèle à taux d'escompte

Dans ce modèle, si l'offre $o_t=(x,1-x)$ est acceptée, les agents reçoivent les gains $g_A = x\cdot\delta^{t-1}$, $g_2 = (1-x)\cdot\delta^{t-1}$, où $\delta$ est un paramètre de l'intervalle ]0,1[
> **Remarque :** les taux d'escompte des 2 joueurs sont supposés égaux, ce qui simplifie l'analyse. Ce modèle revient à considérer que la valeur du gâteau est multipliée par $\delta$ à chaque tour.

On considère le jeu à **horizon fini**. 

2. Analyser les jeux en 1, 2, 3 tours en précisant les équilibres parfaits.

### C. Pour aller plus loin

Emettre une conjecture quant aux équilibres parfaits et aux valeurs de ces jeux lorsque le nombre de tours est un entier $N$ quelconque. Prouver cette conjecture par récurrence.

## Exercice 3 : Enchère de Vickrey

Il s'agit du mécanisme d'enchères vu en cours : les acheteurs soumettent des offres privées, et l'objet va au plus offrant, qui paie au vendeur le prix correspondant à la seconde offre la plus élevée. Les autres acheteurs ne paient rien.

Montrer que le mécanisme est sincère : pour chaque acheteur, faire une offre dont le montant correspond à sa propre évaluation de l'objet est une *stratégie faiblement dominante* (quel que soit le comportement des autres acheteurs --sincères ou non-- il n'existe pas de stratégie strictement plus profitable).
