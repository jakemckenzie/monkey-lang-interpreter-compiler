# Notes

# Chapter 1:

NextToken() within lexer.go got fairly verbose. There's clearly a refactoring 
and way to simplify the logic within the case statement that I'm not seeing at
the moment. For now I've verified that it works and will work towards reducing 
the complexity later.

EDIT: Went back to the chapter. Messed up my case statement fall through logic.

# Chapter 2:

Understanding the [tools](https://ast-grep.github.io/) I use each day are important, abstract syntax trees [aren't likely to go anywhere](https://x.com/tsoding/status/1705377825426125270). On a sidenote, I finished [Hypermedia Systems](https://hypermedia.systems/), a lot to think about. I'll talk more about it here tomorrow and Abstract Syntax Trees. tl;dr it feels like a we had a fork in the road when React happened and we're just getting back to some clarity. Too much engineering effort has been done to make the client the only viable place to serve websites.

stopping point: page 41 "We can parse let statements!"