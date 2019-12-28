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
Cloud Functionsの[Console を使用したクイックスタート](https://cloud.google.com/functions/docs/quickstart-console?hl=ja)の **始める前に** を参照し、Cloud Functionsを作成する。

### Cloud Functionsへデプロイ
### デプロイ時の注意点
#### ①go.modが必要
Goではcloud Functionsへのdeployには対象のソースコード以外にも `go.mod` （と `go.sum` ）ファイルも必要となっている。
参考：[Go での依存関係の指定](https://cloud.google.com/functions/docs/writing/specifying-dependencies-go?hl=ja)

今回のGitHubには既に入っているため、 **作成する必要はない** が作成する場合は[こちらを参照](https://cloud.google.com/functions/docs/writing/specifying-dependencies-go?hl=ja)

#### ②goの関数名は大文字始まり
私が躓いたところでもありますが、 **関数名は大文字始まり** にし他パッケージからの参照を可能とします。

#### ③基本的にfunction名は関数名と同じ
[ローカルマシンからのデプロイ](https://cloud.google.com/functions/docs/deploying/filesystem?hl=ja)に記載があるように `-entry-point フラグを指定しない限り、コードに同じ名前の関数を含める必要があります。` 

### デプロイ
以下のコマンドでデプロイし、cloud console上に `Omubot` というfunctionが作成されたことを確認する。
※gcloudコマンドがセットアップされていない場合は[こちら](https://cloud.google.com/sdk/docs/?hl=ja)を参照

```
gcloud functions deploy Omubot --runtime go111 --trigger-http
```

--runtime : goの実行バージョン
--trigger : トリガーの種類

### Webhookの設定
cloud functionsに登録された関数をLINE Developer内の１で作成したチャネルに設定する。
設定は[こちら](https://developers.line.biz/ja/docs/messaging-api/building-bot/)を参照。

設定するURLはデプロイ時に表示される `httpsTrigger: url:` 
