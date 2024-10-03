# Change the value in quotation marks to change the default name for your companion
# You may want to do this if you're running multiple instances of this bot
# Ex: $companionName = "friend_1"
$companionName = "discord_companion"

### DO NOT EDIT BELOW THIS LINE ###
$inputName = Read-Host "Companion Name (name of the Docker container) is set to $companionName - is this okay? Press Enter to accept this name or enter another one"

if ([string]::IsNullOrWhiteSpace($inputName)) {$companionName = $defaultName}

$discordToken = Read-Host "Enter Discord Token"
$nomiKey = Read-Host "Enter Nomi API Key"
$nomiId = Read-Host "Enter Nomi AI ID for your companion"

docker container rm $companionName -f
docker build -t $companionName $psscriptroot
docker run -d --name $companionName -e DISCORD_BOT_TOKEN=$discordToken -e NOMI_TOKEN=$nomiKey -e NOMI_ID=$nomiId $companionName
