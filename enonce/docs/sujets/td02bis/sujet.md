# IA04 TD02 bis – synchronisation et channels

## Exercice 1 – La routine quotidienne d'Alice et Bob

Chaque matin, Alice et Bob se promènent, et étant pétris d'habitude, ils suivent la même routine tous les jours.

Tout d'abord, ils se préparent tous les deux. Ils saisissent leurs lunettes de soleil, éventuellement une ceinture, ils ferment les fenêtres ouvertes, éteignent les ventilateurs de plafond, et mettent dans leur poche leur téléphone et leurs clés.

Une fois qu'ils sont tous les deux prêts, ce qui prend généralement pour chacun d'eux entre 60 et 90 secondes, ils arment l'alarme, qui a un délai de 60 secondes.

Alors que l'alarme compte à rebours, ils mettent chacun leurs chaussures, un processus qui tend à prendre pour chacun d'eux entre 35 et 45 secondes.

Ensuite, ils quittent la maison ensemble et verrouillent la porte, avant même que l'alarme n'ait terminé son compte à rebours.

### Questions

1. Comment procéder pour simuler (en Go) la matinée d'Alice et Bob. En particulier, comment coordonner la fermeture des fenêtres, des ventilateurs, de la porte et la mise en route de l'alarme ? Qui, d'Alice ou de Bob, doit armer l'alarme une fois que les deux sont prêts ?
2. Écrire un programme pour simuler la routine matinale d'Alice et Bob. Alice et Bob doivent être chacun simulés par une goroutine. Le compte à rebours de l'alarme, une fois armée, ne doit pas être attendu par le reste du programme.
3. Quelle est la différence entre utiliser un `sync.WaitGroup` et un motif `Barrier` pour synchroniser la fin de la préparation ? Quels sont les avantages de chaque approche ?

*Exemple de sortie de l'exécution du programme*

```
Allons nous promener!
Bob a commencé à se préparer
Alice a commencé à se préparer
Alice a passé 72 secondes à se préparer
Bob a passé 76 secondes à se préparer
Alarme armée.
Bob a commencé à mettre ses chaussures
Alarme compte à rebours.
Alice a commencé à mettre ses chaussures
Alice a passé 37 secondes à mettre ses chaussures
Bob a passé 39 secondes à mettre ses chaussures
Fermeture et verrouillage de la porte.
L'alarme a terminé son compte à rebours.
```

## Exercice 2 – La machine à café

Le Phil est fermé et il ne reste qu'une seule machine à café dans tout BF. Il faut entre 15s et 30s pour se faire couler un café. Le cours d'IA04 étant fini, 24 étudiants se précipitent sur la pauvre machine.

### Questions

1. Simuler l'exécution du scénario en Go (NB : utiliser `time.Sleep()` pour simuler les attentes et `math/rand` pour la durée aléatoire). Chaque étudiant sera simulé par une goroutine.
2. Idem avec 2 puis n machines à café.
3. Afficher le temps total nécessaire pour que tous les étudiants aient eu leur café, afin de comparer les cas à 1, 2 et n machines.

## Exercice 3 – Un ping pong en TCP

2 programmes en Go communiquent via TCP. Les *pingers* envoient un message `ping id` et le serveur renvoie `pong id`.

Exemple d'échange :

```
> ping 0
< pong 0
> ping 12
< pong 12
> ping 1
< pong 1
```

### Questions

1. Créer un serveur *pong* capable de gérer plusieurs connexions simultanées.
2. Créer des clients *ping* dans un programme séparé.
3. Modifier le client pour qu'il puisse envoyer plusieurs `ping` en parallèle (une goroutine par ping, ou plusieurs pings en vol sur une même connexion) sans attendre séquentiellement chaque réponse. L'identifiant transmis dans `ping id` / `pong id` permet d'associer chaque réponse à sa requête.
4. (optionnel) Ajouter un timeout côté client avec `context` ou `time.After` pour détecter une absence de réponse.

On pourra utiliser `netcat` (`nc`) afin de tester les clients et les serveurs.