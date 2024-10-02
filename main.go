package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"
    "net/http"
    "os"
    "time"
    "sync"

    "github.com/bwmarrin/discordgo"
)

type QueuedMessage struct {
    Message   *discordgo.MessageCreate
    Session   *discordgo.Session
}

type MessageQueue struct {
    messages []QueuedMessage
    mu       sync.Mutex
}

func (q *MessageQueue) Enqueue(message QueuedMessage) {
    q.mu.Lock()
    defer q.mu.Unlock()
    q.messages = append(q.messages, message)
}

func (q *MessageQueue) Dequeue() (QueuedMessage, bool) {
    q.mu.Lock()
    defer q.mu.Unlock()

    if len(q.messages) == 0 {
        return QueuedMessage{}, false
    }

    message := q.messages[0]
    q.messages = q.messages[1:]
    return message, true
}

func (q *MessageQueue) ProcessMessages() {
    for {
        queuedMessage, ok := q.Dequeue()
        if !ok {
            time.Sleep(1 * time.Second) // No messages in queue, sleep for a while
            continue
        }

        err := sendMessageToAPI(queuedMessage.Session, queuedMessage.Message)
        if err != nil {
            log.Printf("Failed to send message to Nomi API: %v", err)
            q.Enqueue(queuedMessage) // Requeue the message if failed
        }

        time.Sleep(5 * time.Second) // Try to keep from sending messages toooo quickly
    }
}

func sendMessageToAPI(s *discordgo.Session, m *discordgo.MessageCreate) error {
    // Ignore messages from the bot itself - this should be filtered out already but you never know
    if m.Author.ID == s.State.User.ID {
        return nil
    }

    // Check if the message mentions the bot
    for _, user := range m.Mentions {
        if user.ID == s.State.User.ID {
            nomiToken := os.Getenv("NOMI_TOKEN")
            if nomiToken == "" {
                fmt.Println("No Nomi API Key provided. Set NOMI_TOKEN environment variable.")
                return nil
            }

            nomiId := os.Getenv("NOMI_ID")
            if nomiToken == "" {
                fmt.Println("No Nomi ID provided. Set NOMI_ID environment variable.")
                return nil
            }
            url := "https://api.nomi.ai/v1/nomis/" + nomiId + "/chat"

            // Replacing mentions makes it so the Nomi sees the usernames instead of <@userID> syntax
            updatedMessage, err := m.ContentWithMoreMentionsReplaced(s)
            if err != nil {
                log.Printf("Error replacing Discord mentions with usernames: %v", err)
            }

            // Prefix messages sent to the Nomi so they know who they're from and that it's Discord
            // and not a normal Nomi app message
            updatedMessage = "*Discord Message from " + m.Author.Username + ":* " + updatedMessage

            headers := map[string]string{
                "Authorization": nomiToken,
                "Content-Type": "application/json",
            }

            bodyMap := map[string]string{
                "messageText": updatedMessage,
            }
            jsonBody, err := json.Marshal(bodyMap)
            jsonString := string(jsonBody)
            fmt.Printf("Sending message to Nomi API: %v", jsonString)

            req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
            if err != nil {
                log.Fatalf("Error reading HTTP request: %v", err)
            }

            req.Header.Set("Authorization", headers["Authorization"])
            req.Header.Set("Content-Type", headers["Content-Type"])

            client := &http.Client{}
            resp, err := client.Do(req)
            if err != nil {
                log.Fatalf("Error making HTTP request: %v", err)
            }

            defer resp.Body.Close()

            body, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                log.Fatalf("Error reading HTTP response: %v", err)
            }


            if resp.StatusCode != http.StatusOK {
                log.Printf("Error response from Nomi API: %v", string(body))

                // Sometimes Nomi responds with an error: {"error":{"type":"NoReply"}}
                // The Nomi got the message and replied, but the reply wasn't sent back
                // This is only one kind of error, but it seems to be common enough that we
                // Should send it back to the Discord server and let them know something happened
                var errorResult map[string]interface{}
                if err := json.Unmarshal(body, &errorResult); err != nil {
                    log.Fatalf("Error unmarshalling error response: %v", err)
                }

                if errorMessage, ok := errorResult["error"].(map[string]interface{}); ok {
                    if typeValue, ok := errorMessage["type"].(string); ok {
                        if typeValue == "NoReply" {
                            log.Print("'NoReply' error - Sending 'Replied but you did not see it' message to Discord")
                            // Send as a reply to the message that triggered the response, helps keep things orderly
                            _, err := s.ChannelMessageSendReply(m.ChannelID, "❌ ERROR! ❌\nI got your message and I replied to it, but the Nomi API choked and now you don't get to see what I said. I have no idea that you didn't see my response. Try saying 'Sorry, I missed what you said when I said...' and send me your message again.", m.Reference())
                            if err != nil {
                                fmt.Println("Error sending message to Discord: ", err)
                            }
                        }
                    }
                }
            }

            var result map[string]interface{}
            if err := json.Unmarshal(body, &result); err != nil {
                log.Fatalf("Error unmarshalling Nomi API response: %v", err)
            }

            if replyMessage, ok := result["replyMessage"].(map[string]interface{}); ok {
                log.Printf("Received reply message from Nomi API: %v\n", replyMessage)
                if textValue, ok := replyMessage["text"].(string); ok {
                    // Send as a reply to the message that triggered the response, helps keep things orderly
                    _, err := s.ChannelMessageSendReply(m.ChannelID, textValue, m.Reference())
                    if err != nil {
                        fmt.Println("Error sending message to Discord: ", err)
                    }
                }
            }
            return nil
        }
    }
    return nil
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    // Ignore messages from the bot itself
    if m.Author.ID == s.State.User.ID {
        return
    }

    message := QueuedMessage{
        Message: m,
        Session: s,
    }

    queue.Enqueue(message)
}

var queue MessageQueue

func main() {
    botToken := os.Getenv("DISCORD_BOT_TOKEN")
    if botToken == "" {
        fmt.Println("No bot token provided. Set DISCORD_BOT_TOKEN environment variable.")
        return
    }

    dg, err := discordgo.New("Bot " + botToken)
    if err != nil {
        log.Fatalf("Error creating Discord session: %v", err)
    }

    dg.AddHandler(messageCreate)

    err = dg.Open()
    if err != nil {
        log.Fatalf("Error opening Discord connection: %v", err)
    }

    go queue.ProcessMessages()

    fmt.Println("Bot is now running. Press CTRL+C to exit.")
    select {}
}
