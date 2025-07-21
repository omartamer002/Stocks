# 📈 Stock Price Scraper (Go + Colly)

This project is a simple and efficient web scraper built in **Go** using the [Colly](https://github.com/gocolly/colly) library. It scrapes current stock data (company name, price, and percentage change) from **Yahoo Finance** for a list of predefined ticker symbols and saves the results into a CSV file.

---

## 🚀 Features

- Scrapes stock data from Yahoo Finance
- Extracts:
  - Company name
  - Stock price
  - Percentage change
- Writes output to `stocks.csv`
- Uses custom HTTP headers and delay rules to avoid blocking

---

## 🔧 Tech Stack

- **Go (Golang)**
- **Colly** - Web scraping framework
- **CSV** - Built-in Go support

---

## 🧪 Example Output

Sample CSV output:
Company,Price,Change
Microsoft Corporation, $350.12, +0.55%
Apple Inc., $189.98, -0.45%
---

## 📁 How to Run

1. **Install dependencies** (if not already):

   ```bash
   go get github.com/gocolly/colly/v2
   ```
2.	**Run the scraper:
   go run main.go
3.	**Check the output:
You’ll find stocks.csv in your project directory.
    ##📝 Ticker List

The script currently scrapes the following tickers:
MSFT, IBM, GE, UNP, COST, MCD, V, WMT,
DIS, MMM, INTC, AXP, AAPL, BA, CSCO, GS,
JPM, CRM, VZ
