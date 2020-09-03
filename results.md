## Goal : To test the hypothesis that "better(Li, Lj)" by tracking the debug time.
## Methods 
The 3 languages we decided to compare are Scala, Kotlin and Go. To do this we implemented Conway's Game of Life in all three languages and introduced the same bugs in each code: 1 syntactical and 1 logical. We gave each subject the code for one language to debug and kept track of the amount of time they took to debug the code. 

We limited one test subject to one language for a maximum time of 30 minutes. The syntactical bug was introduced in the data types in the function declarations, to explore the complex varieties of data types in that language. The logical bug was introduced to estimate the test subject's ability to look at the remainder of the code and fill in the missing lines of code, which further signifies the language's ease of use.

After the test we asked each subject to fill out a survey that gave us an idea about their experience with the language.
## Materials
### Online IDE
We provided test subjects with online IDEs for debugging the codes. Here is the list of IDEs used:

Scala: https://scastie.scala-lang.org/

Kotlin: https://try.kotlinlang.org/

Go: https://play.golang.org/

### Server
We developed a server that kept track of each test subject's debug time and checked whether or not the subject's output grid matches the correct output. This way we were able to automate the process of evaluating our subjects even when multiple subjects were giving tests simultaneously. The server was also responsible for assigning a token to each test subject by fetching it from our stack of 10 tokens and storing their information in a file.

### Survey
We created the survey using Google Forms. This is the [survey](https://forms.gle/MmuK5oDpU2bK6dTy8) we created.

## Observations

![alt text](https://github.com/Parth27/GameOfLife/blob/master/raw/complete.png?raw=true)
![alt text](https://github.com/Parth27/GameOfLife/blob/master/raw/mean.png?raw=true)
![alt text](https://github.com/Parth27/GameOfLife/blob/master/raw/median.png?raw=true)
![alt text](https://github.com/Parth27/GameOfLife/blob/master/raw/rating.jpg?raw=true)
![alt text](https://github.com/Parth27/GameOfLife/blob/master/raw/doc_rating.jpg?raw=true)


## Conclusions
Based on our observations, we conclude that scala was the easiest language to debug amongst the three.

When we asked our subjects which languages they were familiar with in our survey, we found that most of them were familiar with Java. This could be one possible explanation why people found it easier to debug Scala, since it is very similar to Java.

Despite Kotline also being similar to Java, it was the hardest to debug for our subjects. The main reason for this was that our subjects did not find Kotlin's documentation very helpful, causing them to be stuck on our syntactical bug.
## Threats to validity
- Challenge 1: How much of that tracking can you automate?
- Challenge 2: What parts of this analysis should not be automated? How would you do that ?
- Challenge 3: How to control for the background experience of the subject?
