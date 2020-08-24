# Worktest for networkgaming

Simulate a betting game where users can play or observe a game.

To start game run "docker-compose up" in main folder
visit localhost:3000 in your browser

How to start a game:
1. open localhost:3000
2. write a name and change the switch depending if you want to be observer or player.
3. If switch is on player select two guess values. 
4. Click join.
5. When game starts the left scoreboard will show liveresult after each round.
6. When game ends a text will show below the scoreboards with last winner and their score. 

Services run on nginx proxy to avoid cors issues.

Backend Packages:
- scoring (Contains a function that takes in the value from the game and the users guess and return a score)
- scoreboard (A map structure that keeps track of scores and who won)
- player (Struct with name,websocket and guess) A player is considered a observer if no guess 
- gameservice (binds all the stuff neccesary to play a game together)
- game (Runs for a given amount of round , returns a randomized value in the interval every round through a channel. Randomized on given seed value) 
- config (Reads configuration for game from environment variables or default value) 
- api (Contain the neccesitaty to setup the endpoints for the api ) 



The game is configurable at start using enviorments varaibles
The variables are:
- Rounds int ( Set the number of rounds on each game)
- Sleep time.Duration('example: 1s')  ( Set the sleep time between each new draw in a game)
- SleepBetween  time.Duration('example: 1s')   ( Set the sleep time between each game)
- MinimumPlayer int (Set the minimum number of player that has to join before a game can be started)
- BeginInterval int (Set the lowest number that a user can guess on in the game) 
- EndInterval int (Set the highest number that a user can guess on in the game) 


Testing:
 The backend test can be run with go test ./...
 The frontend components styling is tested using storybook, run command npm run storybook in the web folder.

#TODO:
  1.Two player with the same name could potentialy join the same game and there scores would be joined on the scoreboard.
  it is unneccesary an observer have to input a name before being able to view games.

  2. The way of handling if a user has reached score 21 is a bit embarrasing and hard to grasp. 

  3. The game(name and guess) could be sent as a message through the websocket rather than at the setup of the game. 
     That way the websocket wouldn't need to be close to run a new game with new guess. 


