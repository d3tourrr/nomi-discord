#!/bin/bash

# Change the value in quotation marks to change the default name for your companion
# You may want to do this if you're running multiple instances of this bot
# Ex: companionName = "friend_1"
companionName="discord_companion"
read -p "Companion Name (name of the Docker container) is set to $companionName - is this okay? Press Enter to accept this name or enter another one: " inputName

if [ -n "$inputName" ]; then
  companionName=$inputName
fi

read -p "Enter Discord Token: " discordToken
read -p "Enter Nomi API Key: " nomiKey
read -p "Enter Nomi AI ID for your companion: " nomiId

scriptroot="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
docker container rm $companionName -f
docker build -t $companionName $scriptroot
docker run -d --name $companionName -e DISCORD_BOT_TOKEN=$discordToken -e NOMI_TOKEN=$nomiKey -e NOMI_ID=$nomiId $companionName
