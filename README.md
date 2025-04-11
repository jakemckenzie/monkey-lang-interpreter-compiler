# Notes

# Chapter 1:

NextToken() within lexer.go got fairly verbose. There's clearly a refactoring 
and way to simplify the logic within the case statement that I'm not seeing at
the moment. For now I've verified that it works and will work towards reducing 
the complexity later.

EDIT: Went back to the chapter. Messed up my case statement fall through logic.