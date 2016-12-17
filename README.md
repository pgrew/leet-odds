# leetodds
Poker engine and odds calculator for wildcard and silly games


We need to define things such as the following:

*objects*
- Deck
- Suit
- Card (including wildcards)
- Game (needs to be interface???)
- Dealer
- Player
- Seat
- Table
- Pot
- Hand
- Chip (denominations)
- Muck
- Board
- Round (interface??)

*actions*
- Betting round
- Showdown
- Action (call, raise, fold, all-in, min raise)
- Hand ranker (maybe under dealer or showdown?)
- Bet
- Emote ( "Dirk!", "I just got completely screwed!")
- Exchange chips with fellow players

*Variants*
- Hold 'em
    - You have two cards
- No-peak
    - Players can't see their own hands
- Stud
    - Some player cards are hidden and some are exposed
- Chicago
    - High (or low) card of a suit wins half the pot
- Waterfall
    - Third card on the flop is wild for everyone, turn determines suit of Chicago, and the river determines if it is high or low card (8 means none!)
- Treasure Chest
    - The cards in the muck are revealed after the game is over. Another showdown is done and the winner gets half the pot.
- Cincinatti
    - Each player gets three cards
- Mice and men
    - Not worth explaining right now
- Mali
    - Ditto
- Omaha
    - Each player gets four cards, and they have to play EXACTLY two
- Hi-low
    - The game is won by the high and low hands (pot is split)
- Tight
    - Players must use all their cards
