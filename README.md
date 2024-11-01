# AlgoGo

AlgoGo est un bot de trading automatique développé en Go qui utilise l'API Binance pour exécuter des stratégies de trading basées sur les moyennes mobiles exponentielles (EMA).

## Fonctionnalités

- Récupération des données de marché en temps réel
- Calcul des EMA courtes et longues
- Génération de signaux d'achat et de vente
- Placement automatique d'ordres sur Binance
- Suivi du solde du compte
- Enregistrement et analyse des performances de trading
- Visualisation des métriques de performance

## Prérequis

- Go 1.23.0 ou supérieur
- Python 3.8 ou supérieur (pour l'analyse des performances)
- Un compte Binance (testnet ou réel)
- Clés API Binance avec les permissions nécessaires

## Installation

1. Clonez le dépôt :
   ```bash
   git clone https://github.com/ClementG91/AlgoGo
   cd AlgoGo
   ```

2. Installez les dépendances Go :
   ```bash
   go mod tidy
   ```

3. Installez les dépendances Python pour l'analyse :
   ```bash
   # Création de l'environnement virtuel
   python -m venv venv
   
   # Activation de l'environnement virtuel
   # Sur Windows :
   venv\Scripts\activate
   # Sur Linux/Mac :
   source venv/bin/activate
   
   # Installation des dépendances
   pip install -r requirements.txt
   ```

4. Configurez l'application :
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

1. Lancez le bot de trading :
   ```bash
   go run .
   ```

2. Analysez les performances (dans un terminal séparé) :
   ```bash
   python trade_analyzer.py
   ```

## Analyse des Performances

Le système d'analyse des performances génère plusieurs métriques et visualisations :

### Métriques Calculées
- Nombre total de trades
- Trades gagnants et perdants
- Taux de réussite (%)
- Profit moyen par trade (%)
- Ratio de Sharpe
- Drawdown maximum (%)

### Visualisations
- PnL cumulatif
- Distribution des returns
- Performance moyenne par heure
- Analyse du drawdown

Les résultats sont sauvegardés dans :
- `trades.csv` : Données brutes des trades
- `performance_analysis.png` : Graphiques de performance

## Structure des Fichiers

AlgoGo/
├── main.go                 # Point d'entrée principal
├── config.go              # Gestion de la configuration
├── market_data.go         # Récupération des données de marché
├── ema_calculation.go     # Calcul des EMAs
├── trading_signal.go      # Génération des signaux
├── order_placement.go     # Placement des ordres
├── account_balance.go     # Gestion du compte
├── trade_logger.go        # Enregistrement des trades
├── error_handling.go      # Gestion des erreurs
├── trade_analyzer.py      # Analyse des performances
├── requirements.txt       # Dépendances Python
├── config.json           # Configuration du bot
├── secret.json          # Clés API (non versionné)
└── README.md            # Documentation