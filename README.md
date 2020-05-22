# Overview
Spot is a CLI tool used to talk to the Spotify API.

# Features
Spot currently has the following functionality built in
* Authenticate with Spotify to use their public API
* View currently playing track
* Pause currently playing track
* Play a paused track
* Restart the current track
* Go to the next track
* Go to the previous track
* Check to see if you are following the artist
* Follow the currently playing artist
* Unfollow the currently playing artist
* View the existing OAuth token

# Setup
You can follow the instructions 
[here](https://developer.spotify.com/documentation/general/guides/app-settings/#register-your-app) 
to create a client ID.

Once you create your clientID, you will need to keep the clientID 
and the clientSecret to put in the [configuration](#Configuration) file.


# Configuration
Spot uses a yaml configuration file to store the oauth token, 
the Spotify clientID and client secret. This file is typically
stored at `~/.spot/config.yaml`. Here is an example file
```
spotify_access_token: xyz
spotify_access_token_expiration: 2020-05-22T12:00:00Z
spotify_client: xyz
spotify_refresh_token: xyz
spotify_secret: xyz
spotify_token_type: Bearer
```
The only required values to get started are `spotify_client` and
`spotify_secret`. You will get those when you create your Client ID
on Spotify's developer page.

# First Run
The first time you run, you will need to run `spot auth` to get
an access token and refresh token. They will be saved in your
configuration file and stored for later use.