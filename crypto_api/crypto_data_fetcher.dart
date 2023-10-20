import 'dart:convert';
import 'package:http/http.dart' as http;

class CryptoDataFetcher {
  Map<String, dynamic> cryptoInfo = {};

  Future<void> getCryptoData(String symbol) async {
    String apiUrl =
        'https://api.coingecko.com/api/v3/simple/price?ids=$symbol&vs_currencies=usd';
    String apiUrl24h =
        'https://api.coingecko.com/api/v3/coins/$symbol/market_chart';

    try {
      var response = await http.get(Uri.parse(apiUrl));
      var response24h =
          await http.get(Uri.parse('$apiUrl24h?vs_currency=usd&days=1'));

      if (response.statusCode == 200 && response24h.statusCode == 200) {
        Map<String, dynamic> data = json.decode(response.body);
        Map<String, dynamic> data24h = json.decode(response24h.body);

        double latestPrice = data[symbol]['usd'];
        List<dynamic> prices24h = data24h['prices'];
        double price24hAgo =
            prices24h.isNotEmpty ? prices24h.first[1].toDouble() : 0.0;

        double variation24h = latestPrice - price24hAgo;

        cryptoInfo = {
          'symbol': symbol,
          'latestPrice': latestPrice,
          'variation24h': variation24h,
        };
      } else {
        print('Erro na solicitação para $symbol: ${response.statusCode}');
      }
    } catch (e) {
      print('Erro ao processar a solicitação para $symbol: $e');
    }
  }
}
