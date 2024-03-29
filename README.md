# Spotiyou

Spotiyou is a simple application that can notify you when your friends create a new playlist on Spotify.
It uses the Spotify API to find new playlists and send you a notification.

You may wonder why this project came into existence.
One of my friends, who probably spends half of his life listening to music, would always immediately notice
when I created a new playlist on Spotify and ask me about it in the next conversation.
I, on the other hand, would never even notice when he created a new playlist. So, technology to the rescue!

The notification is implemented via a Notifier Interface, so you can easily implement your own notification system.
A simple example notifier to a Telegram chat is included.

## Usage

First of all you need access to the Spotify API. You can get a client id and client secret from the
[Spotify Developer Dashboard](https://developer.spotify.com/dashboard/).
Moreover, if you want to use the provided Telegram notifier, you need to create a Telegram bot and get the bot token.
Then you need to create a Telegram chat and add the bot to the chat.

By setting the environment variables `SPOTIFY_CLIENT_ID`, `SPOTIFY_CLIENT_SECRET`, and the Telegram `TELEGRAM_TOKEN`
and `TELEGRAM_CHAT_ID` you can run the sample application in the `cmd` folder.
