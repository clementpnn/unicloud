# Documentation Développeur Backend

## Table des matières
1. [Installation](#installation)
2. [Configuration](#configuration)
3. [Stack de monitoring](#stack-de-monitoring)
4. [Démarrage](#démarrage)
5. [Métriques](#métriques)
6. [Maintenance](#maintenance)
7. [Troubleshooting](#troubleshooting)

## Installation

### Prérequis
- Go 1.21+
- PostgreSQL 14+
- Docker & Docker Compose
- Node.js 18+ (pour le frontend)

### Installation des dépendances
```bash
# Installation de Prometheus
brew install prometheus

# Installation de postgres_exporter
brew install postgres_exporter

# Installation de Grafana
brew install grafana

# Installation des dépendances Go
go mod download
```

## Configuration

### Variables d'environnement
Créer un fichier `.env` dans `apps/backend/` :
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=abde-bigdev
DB_PASSWORD=
DB_NAME=postgres
```

### Configuration de PostgreSQL
```bash
# Création de l'utilisateur et de la base
createuser -P abde-bigdev
createdb -O abde-bigdev postgres

# Test de connexion
psql -U abde-bigdev -d postgres
```

### Configuration des exporters

#### postgres_exporter
```bash
# Configuration de la connexion
export DATA_SOURCE_NAME="postgresql://abde-bigdev:@localhost:5432/postgres?sslmode=disable"

# Test de l'exporter
curl http://localhost:9187/metrics
```

#### Prometheus
Fichier `prometheus.yml` :
```yaml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'postgres'
    static_configs:
      - targets: ['localhost:9187']
```

## Stack de monitoring

### Architecture
```
Backend (Go) → PostgreSQL ← postgres_exporter → Prometheus → Grafana
```

### Ports utilisés
- Backend : 3000 (auto-incrémenté si occupé)
- PostgreSQL : 5432
- postgres_exporter : 9187
- Prometheus : 9090
- Grafana : 3000

## Démarrage

### Lancement du backend
```bash
# Dans apps/frontend/
Pnpm dev
```

### Lancement du backend
```bash
# Dans apps/backend/
go run main.go
```

### Lancement de la stack monitoring
```bash
# Démarrage de PostgreSQL (si non démarré)
brew services start postgresql

# Démarrage de postgres_exporter
postgres_exporter

# Démarrage de Prometheus
prometheus --config.file=prometheus.yml

# Démarrage de Grafana
brew services start grafana
```

## Métriques

### Métriques clés PostgreSQL

#### Santé de la base
```promql
# État de la connexion
pg_up

# Nombre de connexions actives
pg_stat_database_numbackends{datname="postgres"}
```

#### Performance
```promql
# Transactions par seconde
rate(pg_stat_database_xact_commit{datname="postgres"}[5m])

# Ratio de cache hit
sum(pg_stat_database_blks_hit) / 
(sum(pg_stat_database_blks_hit) + sum(pg_stat_database_blks_read)) * 100
```

#### Erreurs
```promql
# Transactions échouées
rate(pg_stat_database_xact_rollback{datname="postgres"}[5m])

# Deadlocks
pg_stat_database_deadlocks{datname="postgres"}
```

### Configuration Grafana

#### Ajout de la source de données
1. Accéder à http://localhost:3000 (admin/admin)
2. Configuration → Data Sources → Add data source
3. Sélectionner Prometheus
4. URL : http://localhost:9090
5. Sauvegarder & tester

#### Dashboards essentiels
1. Vue générale PostgreSQL
   - Statut de connexion
   - Nombre de connexions
   - Taille de la base
   ```promql
   pg_database_size_bytes{datname="postgres"} / 1024 / 1024
   ```

2. Performance
   - Taux de transactions
   - Ratio de cache hit
   - Temps de requête moyen

3. Erreurs et alertes
   - Taux d'erreurs
   - Deadlocks
   - Conflits de réplication

## Maintenance

### Logs
```bash
# Logs PostgreSQL
tail -f /usr/local/var/log/postgres.log

# Logs Prometheus
tail -f /usr/local/var/log/prometheus.log

# Logs Grafana
tail -f /usr/local/var/log/grafana/grafana.log
```

### Backup
```bash
# Backup PostgreSQL
pg_dump -U abde-bigdev postgres > backup.sql

# Export dashboards Grafana
curl -H "Authorization: Bearer ${GRAFANA_API_KEY}" \
     http://localhost:3000/api/dashboards/uid/${DASHBOARD_UID} \
     > dashboard-backup.json
```

### Vérification de l'état
```bash
# État des services
brew services list

# Vérification des métriques
curl http://localhost:9187/metrics | grep pg_up
curl http://localhost:9090/api/v1/query?query=pg_up
```

## Troubleshooting

### Problèmes courants

#### Port déjà utilisé
Le backend tentera automatiquement d'utiliser le port suivant disponible.
Logs à vérifier :
```bash
tail -f apps/backend/logs/app.log
```

#### Connexion PostgreSQL échouée
1. Vérifier les variables d'environnement
2. Tester la connexion :
```bash
psql -U abde-bigdev -d postgres
```

#### Métriques manquantes
1. Vérifier postgres_exporter :
```bash
curl http://localhost:9187/metrics
```

2. Vérifier Prometheus :
```bash
curl http://localhost:9090/api/v1/query?query=pg_up
```

### Commandes utiles
```bash
# Restart des services
brew services restart postgresql
brew services restart grafana

# Nettoyage des logs
truncate -s 0 /usr/local/var/log/postgres.log
```