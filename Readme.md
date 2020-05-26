# Action-slack-reporter

## Usage

```
  env:
    SLACK_WEBHOOK: ${{ secrets.SLACK_WEBHOOK }}
    SLACK_CHANNEL: channel_name
    SLACK_COLOR: '#00FF00'
    SLACK_MESSAGE: DEPLOYED!
  uses: maxwellhealth/action-slack-reporter@v1.0.2
```
