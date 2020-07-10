# Introduction 
A Partial List of Commands and How To Use Them
#
DROP

Usage: drop <item>
       drop all.<item>
       drop all
       drop <number> coins

If you lose interest in a carried item, you can drop it on the ground.

    drop bottle
    drop all.bread
    drop all
    drop 1000 coins

See also: DONATE, GET, JUNK
#
EQUIPMENT EQUIP EQ

Usage: equipment

Gives you a list of all the equipment you're wearing, holding, and wielding.

See also: INVENTORY, REMOVE, WEAR
#
EXITS EXIT EX

Usage: exits

Gives you a list of the obvious exits from your location. Of course, the less
obvious ones won't show up here - you have to THINK in order to locate those...
#
GET TAKE

"Get" and "Take" are exactly the same and can be used interchangeably.

Usage: get | take <object>
    get | take all <object>
    get | take all all.<object>
    get | take all.<object> all.<object>

If you find something interesting on the ground, you may use 'get' to pick
it up. 'get' can also be used to extract items from containers.

Examples:

    get sword corpse
    get all corpse
    get all all.bag
    get all.bread all.bag
    
See also: DROP, PUT        
#
INVENTORY INV

Usage: inventory

Show's what you're carrying around.

See also: EQUIPMENT, GRAB, HOLD, REMOVE, WEAR
#
KILL ATTACK HIT

Usage: kill <victim>
    hit <victim>

How you start a fight. Hitting other players is not recommended.

See also: FLEE, WIMPY
#
LOOK L

Usage: look 
    look [in | at] [the] <item>
    look <direction>

Examine your surroundings.

Examples:

    look
    look at the angel
    look in the bag
    look south
    
See also: EXAMINE, GET, READ, TAKE
#
NORTH SOUTH EAST WEST UP DOWN

Usage: north
    south
    east
    west
    up
    down
    
To travel in a given direction.
#
RECALL

Usage: recall

Puts you back to the recall room in the starting zone. For use when you get
hopelessly lost.

Note: might change to a command that isn't available to all once development
gets further along ... might want to make a map.
#
ROOMSTATUS

usage: roomstatus

Shows much information about the room that you are in currently; details
about the room itself, the exits and where they go; who else is in the
room with you.

#
SAY TELL TELLALL WHISPER SHOUT ASK

Usage: say | shout <message>
    tell | whisper | ask <player> <message>
    
Communicate with those who are around you.

Examples:

    say Hey, has anybody seen my cat?
    tell draal hey, how's it going?

You can use ' as a shorthand for say:

    ' hi everybody
    
SHOUT sends your message to everybody around you, not just in the same room.    
You must be level 2 before you can SHOUT.
#
STAT STATS

Usage: stats

Learn something about yourself and your condition.

#
WIELD

Usage: wield <weapon>

Weapons are much more useful for hitting monsters with.

Example:
    wield sword
    wield 3.sword  (the 3rd sword in your inventory)

See also: EQUIPMENT, REMOVE, WEAR    
#
WEAR

Usage: wear <item> [location]

For putting on clothes, armor, that sort of thing.

Also, to try to wear everything in your inventory, use 'wear all'

Examples:
    wear boots
    wear all
    wear ring finger
    
See also: EQUIPMENT, REMOVE
#
WHO

Usage: who

Lists the people currently in the game. Some might be invisible.