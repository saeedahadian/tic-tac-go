# Tic Tac Go

This is a simple implementation of the famous tic-tac-toe game in the Go programming language. This is done only for learning purposes.

# The Definition of the Game

The game implemented here adheres to the following design.

This is a **two-player** game. At the beginning of the game, each player has 3 pieces that could put in a starting position at their turn. Once the pieces were positioned, each player could move one of their pieces at their own turn. The only possible moves are only one cell at each of the 8 directions providing no other pieces are at that cell: left, right, up, down, upper right, upper left, lower right, and lower left.

The game ends once the total number of moves in the game passes 30 or when all pieces of a single player are positioned in a straight line at each one of the 8 aforementioned directions.

A history of the moves should be kept somewhere and once the player clicked on a certain point in the history, the game should be changed into that exact moment in the game history and players could continue playing from that point on.