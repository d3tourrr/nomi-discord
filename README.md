# Nomi-Discord Integration

[Nomi](https://nomi.ai) is a platform that offers AI companions for human users to chat with. They have opened v1 of their [API](https://api.nomi.ai/docs/) which enables Nomi chatting that occurs outside of the Nomi app or website. This Discord bot allows you to invite a Nomi to Discord to chat with people there.

# Setup

1. Make a Discord Application and Bot
   1. Go to the [Discord Developer Portal](https://discord.com/developers/applications)
   1. Create a new application and then a bot under that application
   1. Copy the bot's token
   1. Add the bot to a server with the required permissions (at least "Read Messages" and "Send Messages")
1. Clone this repo: `git clone https://github.com/d3tourrr/nomi-discord.git`
1. Build the Docker image: `docker build -t nomi-discord .`
1. Get your Nomi API token
   1. Go to the [Integration section](https://beta.nomi.ai/profile/integrations) of the Profile tab
   1. Copy your API key
1. Get the Nomi ID (see [Nomi API Doc: Listing your Nomis](https://api.nomi.ai/docs/#listing-your-nomis))
1. Run the Docker container: `docker run -e DISCORD_BOT_TOKEN=$DISCORD_BOT_TOKEN -e NOMI_TOKEN=$NOMI_TOKEN -e NOMI_ID=$NOMI_ID nomi-discord`
1. Interact with your Nomi in Discord!

# Interacting in Discord with your Nomi

This integration is setup so that your Nomi will see messages where they are pinged (including replies to messages your Nomi posts). Discord messages sent to Nomis are sent with a prefix to help your Nomi tell the difference between messages you send them in the Nomi app and messages that are sent to them from Discord. They look something like this.

> `*Discord Message from Bealy:* Hi @Vicky I'm one of the trolls that @.d3tour warned you about.`

In this message, a Discord user named `Bealy` sent a message to a Nomi named `Vicky` and also mentioned a Discord user named `.d3tour`.

Mentions of other users show that user's username Discord property, rather than their server-specific nickname. This was just the easiest thing to do and may change in the future (maybe with a feature flag you can set).

Nomis don't have context of what server or channel they are talking in, and don't see messages where they aren't mentioned in or being replied to.

## Suggested Nomi Configurations

It's a good idea to put something like this in your Nomi's "Backstory" shared note.

> `NomiName sometimes chats on Discord. Messages that come from Discord are prefixed with "*Discord Message from X:*" while messages that are private between HumanName and NomiName in the Nomi app have no prefix. Replies to Discord messages are automatically sent to Discord. NomiName doesn't have to narrate that she is replying to a Discord user.`

You may also wish to change your Nomi's Communication Style to `Texting`.

