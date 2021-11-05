# slack-kintai

slackへの勤怠を自動化するツールです。

- できること
  - 複数チャンネルへの出勤連絡
  - 退勤連絡
  - 退勤連絡を出勤したpostへのリプライ
  - 上記をチャンネル全体へも投稿
    - ただし、slack-apiの制約でユーザーがこの場合だけアプリになってしまう

# 実行方法
- リポジトリをcloneする
- `config.json`を編集する
  - slack tokenの取得などもここで対応
- 出勤
  - `go run kintai.go shukkin`
- 退勤
  - `go run kintai.go taikin`
# configuration

- config.jsonを編集してください。
  - token
    - slack tokenを設定してください。(User Token)
    - 取得方法は下の「slack token取得方法」にて記載
  - channels
    - channelName
      - 勤怠連絡をしたいチャンネル名を入力してください。
    - replyToShukkin
      - 出勤postにリプライをしたい場合、`true`
    - postToChannel
      - リプライをチャンネルにも投稿したい場合、`true`
      - slack-apiの制約で、この場合のみ投稿者がアプリユーザになる
  - shukkin
    - 出勤時に投稿したいコメントを設定してください。
  - taikin
    - 退勤時に投稿したいコメントを設定してください。

# slack token取得方法

- https://api.slack.com/ にアクセスする
- `Create New App`をクリック
- `From Scratch`を選択
- 適当に入力して、`Create App`をクリック
- 左の`OAuth & Permissions`から、`User Token Scopes`で`chat:write`と`search:read`を選択
- `Install to Workspace`をクリック
- `xoxp`からはじまる方がユーザに紐づくtokenになります
