import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Stock Prices',
      home: StockPricesScreen(),
    );
  }
}

class StockPricesScreen extends StatefulWidget {
  @override
  _StockPricesScreenState createState() => _StockPricesScreenState();
}

class _StockPricesScreenState extends State<StockPricesScreen> {
  String apiKey = '66KDWJP2EI10612E';
  List<String> symbols = [
    'AAPL',
    'GOOGL',
    'MSFT',
    'AMZN',
    'TSLA',
    'FB',
    'NVDA'
  ];
  List<Map<String, dynamic>> stockData = [];

  @override
  void initState() {
    super.initState();
    fetchData();
  }

  Future<void> fetchData() async {
    for (String symbol in symbols) {
      await getStockData(apiKey, symbol);
    }
  }

  Future<void> getStockData(String apiKey, String symbol) async {
    String endpoint = 'https://www.alphavantage.co/query';
    String function = 'TIME_SERIES_INTRADAY';
    String interval = '1min';

    String apiUrl =
        '$endpoint?function=$function&symbol=$symbol&interval=$interval&apikey=$apiKey';

    try {
      var response = await http.get(Uri.parse(apiUrl));

      if (response.statusCode == 200) {
        Map<String, dynamic> data = json.decode(response.body);
        Map<String, dynamic> timeSeries = data['Time Series (1min)'];

        if (timeSeries != null) {
          String latestDate = timeSeries.keys.first;
          double latestPrice =
              double.parse(timeSeries[latestDate]['1. open'].toString());
          double previousClose = double.parse(
              timeSeries[timeSeries.keys.elementAt(1)]['4. close'].toString());

          double variation = latestPrice - previousClose;

          Map<String, dynamic> stockInfo = {
            'symbol': symbol,
            'latestPrice': latestPrice,
            'variation': variation,
          };

          setState(() {
            stockData.add(stockInfo);
          });
        } else {
          print('Erro ao obter dados para $symbol: ${data['Error Message']}');
        }
      } else {
        print('Erro na solicitação para $symbol: ${response.statusCode}');
      }
    } catch (e) {
      print('Erro ao processar a solicitação para $symbol: $e');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Stock Prices'),
      ),
      body: ListView.builder(
        itemCount: stockData.length,
        itemBuilder: (context, index) {
          return ListTile(
            title: Text('Ação: ${stockData[index]['symbol']}'),
            subtitle: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                    'Último preço: \$${stockData[index]['latestPrice'].toStringAsFixed(2)}'),
                Row(
                  children: [
                    Text(
                        'Variação: \$${stockData[index]['variation'].toStringAsFixed(2)}  '),
                    Icon(
                      stockData[index]['variation'] >= 0
                          ? Icons.arrow_upward
                          : Icons.arrow_downward,
                      color: stockData[index]['variation'] >= 0
                          ? Colors.green
                          : Colors.red,
                    ),
                  ],
                ),
              ],
            ),
          );
        },
      ),
    );
  }
}
