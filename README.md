# Code Gen '24

## Why should I be interested?

**You probably shouldn't.** I just took this simple idea of generating Go struct code based on JSON files and then polished it quite a bit, trying to see where it would end up. 

I must say I'm happy with the results :) Architectural d0pe right here man. Use it wisely ;)

## So...?

Well, basically you can make the data source for generating the code an interface, defining a method inside of it to generate the actual code, right next to the data itself where it should be. 

Then you make the FileReader work with that interface, being able to plug-n-play a new data source with its own code generation in a matter of minutes.