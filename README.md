# GameOfLife
This is implementation of Conway's Game of Life in three languages: Kotlin, Scala and Go

## Evaluation
We have added the same logical and syntactical bugs in all three codes so that we can evaluate the three languages based on how easy/difficult it is to debug them.

We will ask our test subjects to debug one of the codes. For providing a runtime environment, we will ask our subjects to use ssh to connect to one of our computers where they can edit the code from command line using nano or vim. This way we avoid having to setup runtime environments for each language on the subjects' computers.

We have implemented a Node.js server that will keep communicating with the codes as they debug them. Since we have included logical bugs as well, the server will check when the output of the code matches the desired output and record each subject's debugging time. With three subjects per language, we can use the average debugging time of a language for evaluation.

We will also provide a google form asking questions about the support and documentation provided by each language to evaluate the ease of debugging them.
