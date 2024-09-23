# Comment faire des commits comme un pro

## 1. Faites des commits atomiques

- Un commit = une unité logique de changement
- Évitez de mélanger plusieurs fonctionnalités ou corrections dans un seul commit

## 2. Écrivez des messages de commit clairs et descriptifs

- Structure : Une ligne de titre courte (50 caractères max) suivie d'une description détaillée si nécessaire
- Format recommandé :
  ```
  <type>(<portée>): <sujet>

  <corps>

  <pied>
  ```

- Types courants : feat, fix, docs, style, refactor, test, chore

  - feat (feature) : Ajout d'une nouvelle fonctionnalité.
  - fix : Correction d'un bug.
  - docs : Modifications de la documentation.
  - style : Changements qui n'affectent pas le sens du code (espaces, formatage, points-virgules manquants, etc.).
  - refactor : Modification du code qui n'apporte ni nouvelle fonctionnalité ni correction de bug.
  - test : Ajout de tests manquants ou correction de tests existants.
  - chore : Modifications de la configuration du build ou d'autres changements qui n'affectent pas le code source.
- Exemple :
  ```
  feat(auth): ajouter la fonctionnalité de connexion par Google

  - Implémenter l'authentification OAuth avec Google
  - Ajouter le bouton de connexion dans l'interface utilisateur
  - Mettre à jour la documentation utilisateur

  Closes #123
  ```

## 3. Utilisez le présent de l'impératif

- "Ajouter la fonctionnalité" plutôt que "Ajouté la fonctionnalité" ou "Ajoute la fonctionnalité"

## 4. Faites des commits régulièrement

- Committez fréquemment pour faciliter le suivi des changements et le retour en arrière si nécessaire

## 5. Vérifiez vos changements avant de committer

- Utilisez `git diff` ou l'interface de votre IDE pour revoir vos modifications
- Assurez-vous de n'inclure que les changements pertinents

## 6. Utilisez des branches

- Créez des branches pour chaque nouvelle fonctionnalité ou correction
- Nommez vos branches de manière descriptive : `feature/login-page`, `bugfix/header-alignment`

## 7. Rebasez et nettoyez l'historique avant de fusionner

- Utilisez `git rebase -i` pour nettoyer et réorganiser vos commits avant de les pousser
- Squashez les petits commits de correction en un seul commit significatif

## 8. Utilisez des outils pour standardiser vos commits

- Configurez un hook de pre-commit pour vérifier le style du code
- Utilisez des outils comme Commitizen pour structurer vos messages de commit

## 9. Référencez les issues dans vos commits

- Mentionnez les numéros d'issues dans vos messages de commit : "Fixes #123"

## 10. N'oubliez pas le .gitignore

- Maintenez un fichier .gitignore à jour pour éviter de committer des fichiers inutiles