# Wordle solver, written in GO

This is a basic project whose goal was to help me learn the go programming language.

This program will help you to solve any Wordle related problems you have! It works by 
using all the given information to make a guess at a feasible word that could be the 
answer, this is done through the tracking of all letters that aren't in the word, letters
that we don't know the location of but know are in the word, and letters that we know
the location of.

Once you've begun execution of the program, it will prompt you for the word you have 
guessed, this word can be any word of your choice, even if you guess a terrible word, 
most of the time the program will still get the correct word in the end. After inputting 
the starting word, you must input yet another string, this time defining the state of the
word, this is defined by three parts: a '.' if the letter at the previously inputted word
is grey (isn't in the word), a '?' if the letter at the previously inputted word is 
yellow (in the word but not in the right location) or simply the given letter if it is 
in the right location. Note that you should be using the official 'Wordle' game or some 
other version of it, and this program is simply given all the information, and outputs 
the most likely option.

Example:
```python
input: where
input: ..?r.
// in this case, the word is 'berry'
```

Note that after the first word where you can enter any word you like, the program will
automatically enter the previously generated word as the word you guessed e.g. if the
previous best word was "abhor", then the program will automatically use that as the next
guessed word for you to guess in your Wordle game.
