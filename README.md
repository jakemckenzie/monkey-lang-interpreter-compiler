# Notes

# Chapter 1:

NextToken() within lexer.go got fairly verbose. There's clearly a refactoring 
and way to simplify the logic within the case statement that I'm not seeing at
the moment. For now I've verified that it works and will work towards reducing 
the complexity later.

Update 1: I went back to this problem after struggling to see where I went wrong with 
my logic and I screwed up the fall through logic with what Ball was presenting. I'm trying 
to not just copy the code from the book, read 10 minutes then write it myself but I failed 
with that section

# Chapter 2:

Understanding the [tools](https://ast-grep.github.io/) I use each day are important, abstract syntax trees [aren't likely to go anywhere](https://x.com/tsoding/status/1705377825426125270). On a sidenote, I finished [Hypermedia Systems](https://hypermedia.systems/), a lot to think about. I'll talk more about it here tomorrow and Abstract Syntax Trees. tl;dr it feels like a we had a fork in the road when React happened and we're just getting back to some clarity. Too much engineering effort has been done to make the client the only viable place to serve websites.\

Update 1: Jonathan Blow and Casey had a [really good discussion about parsing an expression](https://youtu.be/fIPO4G42wYE?si=zcVilWngZviayuZR) and how the [dragon book overcomplicated the problem](https://en.wikipedia.org/wiki/Principles_of_Compiler_Design) of parsing. I'm going to be honest, I'm not as experienced of a programmer as jon or casey but I was running into the same problem with parsers they were talking about constructing my markdown to html parser in the boot.dev course. It's the same problem jon was talking about where I wasn't handling the presidence properly once I constructed my abstract tree. With markdown vs monkey (like in the case of this book) or jon's language Jai there is far more ambiguity in the language which made it a hard problem.\

The dragon book was profound when I first read it but the more I am beginning to see its shortcomings. I'm not 
sure if it's such an ivory tower problem like Jon describes but more a contextual problem of the time and few bothering to update it after the fact. We're both spoiled and having our hands held by the tools we have today. Seeing jon step through this problem right as I get to it in the book brought clarity. Ball does a much better job at presenting the ideas than I've run across. I wish I had read the [article he links to from douglas crawford](https://crockford.com/javascript/tdop/tdop.html) when I was writing javascript daily.\

With any of this understanding how to set up order pressidence is a huge problem with this sort of parsing. I think the game programmers writing C coming out of the games industry enums have order presidence baked into them, as far as I know few constructs in the object oriented world of functional world have tools for this. Jon's comment about enums for this in that video was fairly subtle but important I think. I big problem I have is that writing the code almost never matches up 1:1 with the design of the thing. I think a big problem with the Dragon book that Ball doesn't fall short with is that you just have to start writing code. Filling people's head with a dozen pathways and ways to represent the problem falls short to 1 succienct explanation from someone in the trenches then experimenting yourself.


pg. 59 Prefix Operations