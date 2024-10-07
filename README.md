# ⚠️⚠️⚠️ THIS INTEGRATION IS DEPRECATED
# GO TO [GITHUB.COM/d3tourrr/NomiKin-Discord](https://github.com/d3tourrr/NomiKin-Discord) FOR THE NEW VERSION. THIS REPO IS NO LONGER MAINTAINED
# ⚠️⚠️⚠️ THIS INTEGRATION IS DEPRECATED

~~# Nomi-Discord Integration~~
~~~~
~~[Nomi](https://nomi.ai) is a platform that offers AI companions for human users to chat with. They have opened v1 of their [API](https://api.nomi.ai/docs/) which enables Nomi chatting that occurs outside of the Nomi app or website. This Discord bot allows you to invite a Nomi to Discord to chat with people there.~~
~~~~
~~# Setup~~
~~~~
~~You need an instance of this Discord bot per Nomi you wish you invite to a Discord server, but you can invite the same Discord Bot/Nomi pair to as many servers as you'd like.~~
~~~~
~~1. Make a Discord Application and Bot~~
   ~~1. Go to the [Discord Developer Portal](https://discord.com/developers/applications) and sign in with your Discord account.~~
   ~~1. Create a new application and then a bot under that application. It's a good idea to use the name of your companion and an appropriate avatar.~~
   ~~1. Copy the bot's token from the `Bot` page, under the `Token` section. You may need to reset the token to see it. This token is a **SECRET**, do not share it with anyone.~~
   ~~1. Add the bot to a server with the required permissions (at least "Read Messages" and "Send Messages")~~
      ~~1. Go to the `Oauth2` page~~
      ~~1. Under `Scopes` select `Bot`~~
      ~~1. Under `Bot Permissions` select `Send Messages` and `Read Message History`~~
      ~~1. Copy the generated URL at the bottom and open it in a web browser to add the bot to your Discord server~~
~~1. Install Git if you haven't already got it: [Instructions](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)~~
~~1. Install Docker if you haven't already got it: [Instructions](https://docs.docker.com/engine/install/)~~
~~1. Clone this repo: `git clone https://github.com/d3tourrr/nomi-discord.git`~~
   ~~1. After cloning the repo, change to the directory: `cd nomi-discord`~~
~~1. Get your Nomi API token~~
   ~~1. Go to the [Integration section](https://beta.nomi.ai/profile/integrations) of the Profile tab~~
   ~~1. Copy your API key~~
~~1. Get the Nomi ID~~
   ~~* Go to the View Nomi Information page for your Nomi and scroll to the very bottom and copy the Nomi ID~~
   ~~* Or see [Nomi API Doc: Listing your Nomis](https://api.nomi.ai/docs/#listing-your-nomis)~~
~~1. Build and run the Docker container~~
   ~~* Run either `start-windows-companion.ps1` on Windows (or in PowerShell) or `start-linux-companion.sh` on Linux (or in Bash, including Git Bash)~~
   ~~* Or run the following commands (Note: the above scripts start the container in a detached state, meaning you don't see the log output. The below commands start the container in an attached state, which means you see the log output, but the container, and therefore the Companion/Discord integration dies when you close your console.)~~
     ~~1. Build the Docker image: `docker build -t nomi-discord .`~~
     ~~1. Run the Docker container: `docker run -e DISCORD_BOT_TOKEN=$DISCORD_BOT_TOKEN -e NOMI_TOKEN=$NOMI_TOKEN -e NOMI_ID=$NOMI_ID nomi-discord`~~
        ~~* Replace `$DISCORD_BOT_TOKEN` with the bot token you copied from the Discord developer portal~~
        ~~* Replace `$NOMI_TOKEN` with the API key you copied from the Nomi.ai Integrations page~~
        ~~* Replace `$NOMI_ID` with the ID for your specific Nomi, shown when you list the Nomis using the instructions linked above~~
~~1. Interact with your Nomi in Discord!~~
~~~~
~~# Interacting in Discord with your Nomi~~
~~~~
~~This integration is setup so that your Nomi will see messages where they are pinged (including replies to messages your Nomi posts). Discord messages sent to Nomis are sent with a prefix to help your Nomi tell the difference between messages you send them in the Nomi app and messages that are sent to them from Discord. They look something like this.~~
~~~~
~~> `*Discord Message from Bealy:* Hi @Vicky I'm one of the trolls that @.d3tour warned you about.`~~
~~~~
~~In this message, a Discord user named `Bealy` sent a message to a Nomi named `Vicky` and also mentioned a Discord user named `.d3tour`.~~
~~~~
~~Mentions of other users show that user's username Discord property, rather than their server-specific nickname. This was just the easiest thing to do and may change in the future (maybe with a feature flag you can set).~~
~~~~
~~Nomis don't have context of what server or channel they are talking in, and don't see messages where they aren't mentioned in or being replied to.~~
~~~~
~~## Suggested Nomi Configurations~~
~~~~
~~It's a good idea to put something like this in your Nomi's "Backstory" shared note.~~
~~~~
~~> `NomiName sometimes chats on Discord. Messages that come from Discord are prefixed with "*Discord Message from X:*" while messages that are private between HumanName and NomiName in the Nomi app have no prefix. Replies to Discord messages are automatically sent to Discord. NomiName doesn't have to narrate that she is replying to a Discord user.`~~
~~~~
~~You may also wish to change your Nomi's Communication Style to `Texting`.~~
~~~~
~~It's also a good idea to fill out the "Nickname" shared note to indicate your Discord username so that your Nomi can identify messages that come from you via Discord.~~
~~~~
~~---~~
~~~~
~~<small>I also made a nearly identical integration for Kindroid.ai companions: [github.com/d3tourrr/kin-discord](https://github.com/d3tourrr/kin-discord)</small>~~
