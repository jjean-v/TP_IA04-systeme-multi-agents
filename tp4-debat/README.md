# TP4: Débats d'IA


Le but de ce TP est d'avoir deux agents **Debatteur** qui débattent entre elles sur un sujet. Chaque Agent a un prompt de départ afin d'affiner son personnage et ses réponses.


Le débat sera géré par un agent **Modérateur**, qui s'occupera de donner la parole à l'un puis à l'autre.

Au bout d'un nombre de tour donné, le Modérateur arrête le débat et appelle un agent **Juge**, qui aura l'intégrité de la conversation et qui devra décider d'un vainqueur.

Chaque agent Débatteur, peut à tous moment, lorsque leur réponse est trop grande, appeller un agent **Synthetiseur** qui devra réduire la réponse de l'agent, dans un souci d'économie des tokens.
