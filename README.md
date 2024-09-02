# AlgoGo

AlgoGo est un bot de trading automatique développé en Go qui utilise l'API Binance pour exécuter des stratégies de trading basées sur les moyennes mobiles exponentielles (EMA).

## Fonctionnalités

- Récupération des données de marché en temps réel
- Calcul des EMA courtes et longues
- Génération de signaux d'achat et de vente
- Placement automatique d'ordres sur Binance
- Suivi du solde du compte

## Prérequis

- Go 1.23.0 ou supérieur
- Un compte Binance (testnet ou réel)
- Clés API Binance avec les permissions nécessaires

## Installation

1. Clonez le dépôt :
   ```bash
   git clone https://github.com/ClementG91/AlgoGo
   cd AlgoGo
   ```

2. Installez les dépendances :
   ```bash
   go mod tidy
   ```

3. Configurez l'application :
   - Ajustez les paramètres dans `config.json` selon vos besoins
   - Créez un fichier `secret.json` avec vos clés API Binance :
     ```json
     {
       "apiKey": "votre-clé-api",
       "apiSecret": "votre-secret-api"
     }
     ```

## Configuration

Modifiez `config.json` pour ajuster les paramètres de trading :

- `symbol` : Paire de trading (ex: "BTCUSDT")
- `interval` : Intervalle de temps pour les données (ex: "1m", "5m", "1h")
- `quantity` : Quantité à trader
- `shortEMA` : Période pour l'EMA courte
- `longEMA` : Période pour l'EMA longue
- `sleepTime` : Temps d'attente entre chaque cycle (en secondes)
- `assets` : Actifs à surveiller dans le solde du compte

## Utilisation

Lancez le bot avec :

```
go run .
```
