# Le chat et les souris

## Énoncé

Des souris et un chat peuplent un monde. Le but des souris est de trouver du fromage et de se reproduire tandis que le but du chat est de manger des souris.

Chaque bête, à sa naissance, est dotée d’une certaine vélocité qui définit la distance qu’elle parcourt en un instant et d’une acuité qui lui permet de détecter les autres jusqu’à une certaine distance. Si un animal ne mange pas pendant un certain temps il devient affamé, puis il meurt. En revanche, un animal affamé ira plus vite qu’un animal repu. À intervalles donnés, du fromage apparaît en différents points du monde. Il existe des trous de souris au bord du monde où les souris peuvent se cacher.

Les règles sont les suivantes :

- le chat poursuit la souris la plus proche et la mange s’il l’attrape ;
- une souris mâle adulte poursuit une souris femelle adulte non fécondée la plus proche et la féconde s’il l’attrape ;
- une souris mâle adulte fuit le chat s’il est plus proche que la souris femelle adulte non fécondée la plus proche ;
- une souris femelle adulte fuit le chat si elle le détecte et recherche du fromage si elle est affamée. Sinon elle reste sur place ;
- après un temps de gestation, une souris femelle adulte fécondée met bas 4 petits qu’elle envoie dans quatre directions différentes ;
- les souriceaux s’écartent de leur lieu de naissance jusqu’à ce qu’ils deviennent adultes ;
- deux animaux peuvent être à la même position en même temps ;
- lorsqu’elles sont affamées, les souris sont attirées par le fromage plus que par les autres souris ;
- le chat peut passer sur du fromage, il ne se passe rien ;
- dès qu’une souris est sur du fromage, elle (et elle seule) le mange entièrement ou en partie. Elle est plus ou moins repue ;
- les trous de souris ne peuvent contenir qu'une seule souris.

## Questions

1. Qu'est-ce qu'un système multi-agent selon la proposition de [Ferber 1995] ? En particulier, que sont les agents, les objets et l'environnement ?
2. Proposer une modélisation du problème sous forme d'agents, d'objets et d'environnement. Donner les attributs et les méthodes associées à chacun des agents. En particulier, comment gérer les souriceaux ?
3. Proposer une architecture pour simuler cet écosystème.
4. À quelle question cette simulation peut-elle répondre ? Quels indicateurs définir pour y répondre ? 
