import 'package:flutter/material.dart';
import 'crypto_data_fetcher.dart';
import 'crypto_tile.dart';

class CryptoPricesScreen extends StatefulWidget {
  @override
  _CryptoPricesScreenState createState() => _CryptoPricesScreenState();
}

class _CryptoPricesScreenState extends State<CryptoPricesScreen> {
  List<String> symbols = [
    'bitcoin',
    'ethereum',
    'ripple',
    'litecoin',
    'cardano'
  ];
  List<Map<String, dynamic>> cryptoData = [];
  CryptoDataFetcher dataFetcher = CryptoDataFetcher();

  @override
  void initState() {
    super.initState();
    fetchData();
  }

  Future<void> fetchData() async {
    for (String symbol in symbols) {
      await dataFetcher.getCryptoData(symbol);
      setState(() {
        cryptoData.add(dataFetcher.cryptoInfo);
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Crypto Prices'),
      ),
      body: ListView.builder(
        itemCount: cryptoData.length,
        itemBuilder: (context, index) {
          return CryptoTile(cryptoData: cryptoData[index]);
        },
      ),
    );
  }
}
