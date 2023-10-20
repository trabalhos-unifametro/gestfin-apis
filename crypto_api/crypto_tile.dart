import 'package:flutter/material.dart';

class CryptoTile extends StatelessWidget {
  final Map<String, dynamic> cryptoData;

  CryptoTile({required this.cryptoData});

  @override
  Widget build(BuildContext context) {
    return ListTile(
      title: Text('Criptomoeda: ${cryptoData['symbol']}'),
      subtitle: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
              'Último preço: \$${cryptoData['latestPrice'].toStringAsFixed(2)}'),
          Row(
            children: [
              Text(
                  'Variação 24h: \$${cryptoData['variation24h'].toStringAsFixed(2)}  '),
              Icon(
                cryptoData['variation24h'] >= 0
                    ? Icons.arrow_upward
                    : Icons.arrow_downward,
                color:
                    cryptoData['variation24h'] >= 0 ? Colors.green : Colors.red,
              ),
            ],
          ),
        ],
      ),
    );
  }
}
