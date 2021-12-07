# discordfile
Uploads local files to discord channels. Modified version of (https://github.com/Not-Cyrus/discord-file-webhook-upload) to support arguments.

## Installation Instructions

Go 1.6:
```bash
▶ go get github.com/vermsec/discordfile
```

Go 1.7:
```bash
▶ go install github.com/vermsec/discordfile@latest
```

## Help Menu
```bash
  -file string
        Path to file which is to be sent (default "nofile")
  -help
        usage info
  -version
        current version
  -webhook string
        Webhook URL to send the data to (default "whook")
```

## Usage
```bash
▶ discordfile -webhook https://discord.com/api/webhooks/213123 -file /path/to/file
```
