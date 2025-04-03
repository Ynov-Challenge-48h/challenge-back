# Projet48h

## Méthode de travail :
### Flutter 
1. folder by feature

### Commits 
- gitmoji correspondante : https://gitmoji.dev/
 Example ajout de fonctionnalité : :sparkles: added login page 

### Organisation des branches
1. **Branches principales** :
   - **`main`** : Branche principale représentant l'environnement de **préproduction**.
   - **`staging`** : Branche dédiée aux tests et validations.

2. **Branches de travail** :
   - Préfixées selon le domaine :
     - **`feat/`** Exemple : `feat/add-login-page`

### Workflow GitHub

1. **Création de branche** :  
   Travailler sur une branche dédiée (préfixée par `feat/...`).

2. **Passage en staging** :  
   Fusionner la branche dans `staging`. Exécuter des tests pour valider les modifications.

3. **Validation et sauvegarde** :  
   Avant de fusionner dans `main`, effectuer une sauvegarde pour prévenir tout problème.

4. **Mise à jour de main** :  
   Fusionner `staging` dans `main` après validation complète.


### Énnoncée :
application mobile :
- scan de cartes d'identité
- extraction des données afin de simplifier la gestion administrative
-  enregistre le fichier scanné sur son poste en local 
- stockage de : nom, prénom, adresse, CNI (Date de fin de validité de, Numéro de la CNI.)
- sécurité du transfert, du traitement et du stockage de ces données (conformément aux normes en vigueur.)

l'app doit pouvoir :
- Sélectionner un client dans une liste de clients 
- Sélectionner une personne physique présente dans ce dossier client
- Appeler une API pour venir passer en paramètre
- Prendre en photo sa carte d’identité: la détourer, remettre à plat, etc., 
