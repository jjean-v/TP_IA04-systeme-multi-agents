---
tags: TD
---

# [IA04] - TD - Répartition de ressources : échanges et indices de pouvoir


## Jeux coopératifs : le traité de Rome

- 6 pays partenaires ayant chacun un nombre de votes : Allemagne (4), France (4), Italie (4), Belgique (2), Pays-Bas (2), Luxembourg (1)
- une proposition est acceptée si et seulement si elle obtient au moins 12 voix

1. Donner une modélisation pour le jeu sous-jacent et expliciter la fonction d'utilité $u$ associée.
2. Calculer les valeurs de Shapley et Banzhaf de chacun des joueurs.
3. Les valeurs de Shapley et Banzhaf sont elles **efficaces** pour ce jeu ?
4. Montrer que les valeurs de Banzhaf et de Shapley vérifient la propriété de **joueur nul**.


## Partage de ressources : *Top Trading Cycle*

7 philosophes, Camus, Descartes, Foucault, Laplace, Morin, Pascal et Voltaire se retrouvent autour d'un somptueux dîner afin de philosopher. L'hôte des lieux prépare un dessert spécial à chacun d'entre eux et les répartit de la manière suivante :
- Crumble -> Camus
- Éclair -> Descartes
- Flan -> Foucault
- Mille-feuille -> Laplace 
- Opéra -> Morin
- Paris-Brest -> Pascal 
- Trianon -> Voltaire

Chacun n'ayant pas son dessert préféré dans son assiette, les convives décident de se lancer dans des échanges afin de mieux finir leur repas. Leurs préférences individuelles sont représentées par le tableau :


| ≥<sub>Camus</sub>| ≥<sub>Descartes</sub> | ≥<sub>Foucault</sub>  | ≥<sub>Laplace</sub>| ≥<sub>Morin</sub> | ≥<sub>Pascal</sub>  |≥<sub>Voltaire</sub>  |
|---|---|---|---|---|---|---|
| Crumble | Flan | Mille-feuille | Opéra | Flan | Crumble | Éclair |
| Flan | Opéra | Opéra | Crumble | Trianon | Trianon | Paris-Brest |
| Mille-feuille | Paris-Brest | Flan | Flan | Crumble | Flan | Flan |
| Paris-Brest | Éclair | Crumble | Mille-feuille | Opéra | Mille-feuille | Crumble |
| Opéra | Crumble | Éclair | Éclair | Paris-Brest | Éclair | Opéra |
| Éclair | Trianon | Paris-Brest | Trianon | Mille-feuille | Paris-Brest | Mille-feuille |
| Trianon | Mille-feuille | Trianon | Paris-Brest | Éclair | Opéra | Trianon |

1. Appliquer l'algorithme *Top Trading Cycle* et donner étape par étape les cycles d'échanges et l'affectation finale optimale.
2. Montrer qu'à chaque étape de *TTC*, le **bien-être utilitariste** augmente strictement.