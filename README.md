# GameOfLife
This is implementation of Conway's Game of Life in three languages: Kotlin, Scala and Go

## Methodology 

We want to have a fair comparison of the three languages in terms of debuggability and hence, limited one test subject to one language for a maximum time of 30 minutes. We have created all three codes with the same logic and added the same logical and syntactical bugs in all three codes so that we can evaluate the three languages equivalently. The syntactical bugs are introduced in the data types in the data declarations, to explore the complex varieties of data types in that language. The logical bugs are introduced to estimate the test subject's ability to look at the remainder of the code and fill in the missing lines of code, which further signifies the language's ease of use.

We did not test the code on all three languages for the following reasons:
- If we would have created the same code logic and the similar bugs, the language solved by the test subject first will be at a disadvantage and the languages debugged later would be easier to debug. This leads to a biased assessment of the languages.
- If we would have created the code in the three languages with different bugs or different logic, it also leads to a biased comparison of languages.


## Evaluation

We will ask our test subjects to debug one of the codes. For providing a runtime environment, we will ask our subjects to connect on online IDEs where we have setup our library dependencies. This way we avoid having to setup runtime environments for each language on the subjects' computers.

We have implemented a Node.js server that will keep communicating with the codes as they debug them. Since we have included logical bugs as well, the server will check when the output of the code matches the desired output and record each subject's debugging time. With three subjects per language, we can use the average debugging time of a language for evaluation. The tokens are assigned randomly from the queue to the test subject by the server.

We will also provide a google form asking questions about the support and documentation provided by each language to evaluate the ease of debugging them.
