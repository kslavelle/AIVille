# AI-Ville

A simulation game where the player has to write the logic / intellegence required to manage a city. The city has some basic need's that have to be satisified such as meeting it's energy demands, meeting it's water demands, meeting the demand of raw materials and food supply. The challenge of the game is to write code that is able to manage to complex interplay beetween these resources.

## Security Considerations
1. We should limit the number of requests a user can make per unit time
2. The client should be capable of refreshing the user's access token so they can have uninterupted play

## Game States
1. The game cant be won, the objective is to survive as long as possible.
2. The game can be lost by failed to meet any of the _worlds_ requirements.
    1. Chances in temperature (caused by lack of or excessive use of other resources) exceeding a defined max
    2. Unable to provide enough food for the people of the city (due to the climate, i.e water, heat etc)

## Game Logic -> Time
1. The game and consequently the logic need to have a temporal component. Things need to have effects over time. In order to prevent having to continually run calculations on a server to keep a consistent view of each players game the calculation on state should be factored in when they do a call to the API. See the example below:
    1. @3pm a player add's a new coal power station to their world (This will add CO2 to the environment over time)
    2. @6pm the player makes anothe API request. We calculate the duration since the last world action (6-3 = 3) and consequently add 3hours worth of CO2 to the environment from the power plant. (This is a simplified example, there may be no CO2 added if there was also a carbon sink)
2. Time will run at an accelerated pace in the game i.e 1hr IRL = 1 month in game.
3. Players should be able to pause there game if they wish, this is so they can take time off without it destrying their world.

## Machine Interface
1. The main bulk of the game should consist of a REST API to manipulate and get info on the state of the world and a client written in something like python to serve as a wrapper around the API to make it easy for a machine to use.
3. In order for a machine to interact with the client we should have methods that are suitable for machines to interact with... see below

```python
# this would be hard to interact with programatically
def add_coal_plant(...):
  ...

# this would be easy to interact with programatically
def add_resource(type='coal_plant'):
  ...
```

## POC Starting Point
1. A database capable of holding:
    1. game [game_id, paused, owner, last_operation]
    2. game_state [state_id, game_id, CO2/hr, food/hr, energy/hr, money]
    3. game_resources [game_id, ] 
    3. resource_type [resouce_id, resource_name] -> _lookup table_
    4. energy_resource
    5. food_resource
    6. human_resource
    7. environment_resource
