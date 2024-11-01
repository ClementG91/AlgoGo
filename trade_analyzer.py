import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
import seaborn as sns
from datetime import datetime

class TradeAnalyzer:
    def __init__(self, csv_path='trades.csv'):
        self.df = pd.read_csv(csv_path)
        self.df['Timestamp'] = pd.to_datetime(self.df['Timestamp'])
        self.prepare_data()

    def prepare_data(self):
        """Prépare les données pour l'analyse"""
        self.df['Date'] = self.df['Timestamp'].dt.date
        self.df['Hour'] = self.df['Timestamp'].dt.hour
        
        # Calcul des rendements quotidiens
        daily_returns = self.df.groupby('Date').agg({
            'PnL': 'sum',
            'PnLPercentage': 'sum',
            'Price': 'last',  # Prix de clôture du dernier trade
            'Quantity': 'sum'  # Volume total tradé
        }).reset_index()
        
        # Calcul du rendement cumulatif
        daily_returns['Cumulative_PnL'] = daily_returns['PnL'].cumsum()
        daily_returns['Cumulative_Return'] = (1 + daily_returns['PnLPercentage']/100).cumprod() - 1
        
        self.daily_returns = daily_returns
        
    def calculate_metrics(self):
        """Calcule les métriques principales"""
        # Métriques par trade
        total_trades = len(self.df)
        winning_trades = len(self.df[self.df['PnL'] > 0])
        losing_trades = len(self.df[self.df['PnL'] < 0])
        win_rate = (winning_trades / total_trades) * 100 if total_trades > 0 else 0
        
        # Métriques quotidiennes
        daily_returns = self.daily_returns['PnLPercentage']
        profitable_days = len(daily_returns[daily_returns > 0])
        total_days = len(daily_returns)
        daily_win_rate = (profitable_days / total_days) * 100 if total_days > 0 else 0
        
        # Calcul du Ratio de Sharpe quotidien (supposant un taux sans risque de 0%)
        daily_sharpe = np.sqrt(252) * (daily_returns.mean() / daily_returns.std()) if len(daily_returns) > 0 else 0
        
        # Maximum Drawdown sur les rendements cumulatifs
        cumulative_returns = self.daily_returns['Cumulative_Return']
        rolling_max = cumulative_returns.expanding(min_periods=1).max()
        drawdowns = (cumulative_returns - rolling_max) / rolling_max
        max_drawdown = drawdowns.min() * 100

        metrics = {
            'Total Trades': total_trades,
            'Winning Trades': winning_trades,
            'Losing Trades': losing_trades,
            'Trade Win Rate (%)': win_rate,
            'Total Trading Days': total_days,
            'Profitable Days': profitable_days,
            'Daily Win Rate (%)': daily_win_rate,
            'Average Daily Return (%)': daily_returns.mean(),
            'Daily Sharpe Ratio': daily_sharpe,
            'Max Drawdown (%)': max_drawdown,
            'Total Return (%)': cumulative_returns.iloc[-1] * 100 if len(cumulative_returns) > 0 else 0
        }
        
        return metrics

    def plot_performance(self):
        """Génère des graphiques de performance"""
        plt.style.use('seaborn')
        fig, axes = plt.subplots(2, 2, figsize=(15, 10))
        
        # PnL Cumulatif quotidien
        axes[0, 0].plot(self.daily_returns['Date'], self.daily_returns['Cumulative_PnL'])
        axes[0, 0].set_title('PnL Cumulatif Quotidien')
        axes[0, 0].set_xlabel('Date')
        axes[0, 0].set_ylabel('PnL ($)')
        plt.xticks(rotation=45)

        # Distribution des returns quotidiens
        sns.histplot(data=self.daily_returns['PnLPercentage'], bins=50, ax=axes[0, 1])
        axes[0, 1].set_title('Distribution des Returns Quotidiens')
        axes[0, 1].set_xlabel('Return Quotidien (%)')
        
        # Performance moyenne par heure (sur tous les trades)
        hourly_perf = self.df.groupby('Hour')['PnLPercentage'].mean()
        axes[1, 0].bar(hourly_perf.index, hourly_perf.values)
        axes[1, 0].set_title('Performance Moyenne par Heure')
        axes[1, 0].set_xlabel('Heure')
        axes[1, 0].set_ylabel('Return Moyen (%)')

        # Drawdown quotidien
        cumulative_returns = self.daily_returns['Cumulative_Return']
        rolling_max = cumulative_returns.expanding(min_periods=1).max()
        drawdowns = (cumulative_returns - rolling_max) * 100
        axes[1, 1].fill_between(self.daily_returns['Date'], drawdowns.values, 0)
        axes[1, 1].set_title('Drawdown Quotidien')
        axes[1, 1].set_xlabel('Date')
        axes[1, 1].set_ylabel('Drawdown (%)')
        plt.xticks(rotation=45)

        plt.tight_layout()
        plt.savefig('performance_analysis.png')
        plt.close()

        # Graphique supplémentaire pour la distribution des trades par jour
        plt.figure(figsize=(10, 6))
        trades_per_day = self.df.groupby('Date').size()
        plt.hist(trades_per_day, bins=20)
        plt.title('Distribution du Nombre de Trades par Jour')
        plt.xlabel('Nombre de Trades')
        plt.ylabel('Nombre de Jours')
        plt.savefig('trades_distribution.png')
        plt.close()

def main():
    analyzer = TradeAnalyzer()
    metrics = analyzer.calculate_metrics()
    
    print("\n=== Métriques de Performance ===")
    print("\nMétriques Quotidiennes:")
    daily_metrics = ['Total Trading Days', 'Profitable Days', 'Daily Win Rate (%)', 
                    'Average Daily Return (%)', 'Daily Sharpe Ratio', 'Total Return (%)']
    for metric in daily_metrics:
        print(f"{metric}: {metrics[metric]:.2f}")
    
    print("\nMétriques par Trade:")
    trade_metrics = ['Total Trades', 'Winning Trades', 'Losing Trades', 'Trade Win Rate (%)']
    for metric in trade_metrics:
        print(f"{metric}: {metrics[metric]:.2f}")
    
    print(f"\nMax Drawdown (%): {metrics['Max Drawdown (%)']:.2f}")
    
    analyzer.plot_performance()
    print("\nGraphiques générés et sauvegardés dans 'performance_analysis.png' et 'trades_distribution.png'")

if __name__ == "__main__":
    main() 