# Définition des standards techniques

## Normes RGPD et souveraineté des données

1. **Régions Cloud**

- Choisir les régions France et Allemagne
- Vérifier que les données restent dans l’UE et qu’il n’y a pas de transfert en dehors.

2. **Data residency**

- Configurer les buckets, bases de données et services pour restreindre la localisation aux régions UE.
- Mettre en place des logs d’audit (Admin Activity, Data Access) pour prouver la traçabilité.

3. **Gestion des consentements et protection des données**

- Même si c’est un environnement de test, prévoir comment l’application gère les données utilisateurs.
- Sécuriser les accès (chiffrement en transit et au repos).

## Haute disponibilité et multi-région

1. **Disaster Recovery (DR)**

- Planifier la réplication entre régions, ou au moins des backups réguliers hors site.
- Mettre en place un Load Balancer global pour la répartition du trafic.

2. **Mise en place de SLA**

- Définir des objectifs de temps de récupération (RTO) et de point de récupération (RPO).
- Choisir les zones pour garantir la redondance.

## Sécurité

1. **IAM (Identity and Access Management)**

- Assignation de rôles « least privilege » à chaque membre du projet.
- Audit des comptes de service (qui déploie ? qui administre ?)

2. **Chiffrement**

- Utilisation de KMS ou Secrets Manager pour stocker mots de passe et clés API.
- Assurer le chiffrement des volumes et des bases de données (ex. Transparent Data Encryption).

3. **Conformité aux bonnes pratiques DevSecOps**

- Intégration de tests de sécurité (SAST, DAST) dans le pipeline CI/CD.
- Monitorer les vulnérabilités des images Docker (Artifact Registry, etc.).

## Documentation et traçabilité

1. **Architecture**

- Tenir à jour un diagramme ou un document décrivant chaque composant (service, base, réseau, pipeline).

2. **Procédures**

- Documenter dès maintenant : comment créer une VM, comment se connecter à la base, comment gérer les rôles IAM, etc.
- Sauvegarde et restauration
- Définir comment et où les sauvegardes sont stockées, comment faire la restauration en cas de sinistre.
