# LINE_Bot_parrot_return
Google Cloud Functionsを使ったおうむ返しをするLineBot
BotということでLINEの[messagingAPI](https://developers.line.biz/ja/reference/messaging-api/)を利用。

LINE Botのサンプルを使いました
2019年12月時点の情報です<br>

## 作成方法

### チャネルの作成
LINE messagingAPIを利用するためにチャネルを作成したければならない。
チャネルとは、[LINE Developers](https://developers.line.biz/ja/docs/messaging-api/getting-started/)によると

>チャネルは、LINEプラットフォームが提供する機能を、プロバイダーが開発するサービスで利用するための通信路です。LINEプラットフォームを利用するには、チャネルを作成し、サービスをチャネルに関連付けます。チャネルを作成するには、名前、説明文、およびアイコン画像が必要です。チャネルを作成すると、固有のチャネルIDが識別用に発行されます。

となっており、プロバイダー（開発組織）とAPIとかの機能を紐付けるもので
チャネルによって紐付けられた機能が利用できるようになる。

チャネルの作成方法はLINE Developersで出ているので[そちら](https://developers.line.biz/ja/docs/messaging-api/getting-started/)を参考にした方がいいです。

今回はBotを作成するので **channel Typeはmessaging API** を選択してください。


### CHANNEL_SECRET＆CHANNEL_TOKENの入力
下記のコマンドで本リポジトリをcloneしてください。

```
$ cd {cloneするディレクトリ}
$ git clone https://github.com/Hiroya3/LINE_Bot_parrot_return.git
```

その後、server.goの

```
bot, err := linebot.New(
  "CHANNEL_SECRET",
  "CHANNEL_TOKEN",
)
```

の `CHANNEL_SECRET` , `CHANNEL_TOKEN` の部分にそれぞれ
CHANNEL_SECRET : チェネル基本設定>チャネルシークレット 
CHANNEL_TOKEN  : Messaging API設定>チャネルアクセストークン（ロングターム）
を入力。

### Cloud Functionsの作成

### Cloud Functionsへデプロイ

### Webhookの設定
