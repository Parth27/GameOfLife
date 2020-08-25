# GameOfLife
This is implementation of Conway's Game of Life in three languages: Kotlin, Scala and Go

# Evaluation
We have added the same bugs in all 3 codes so that we can evaluate the 3 languages based on how easy/difficult it is to debug them.

We will ask our test subjects to debug one of the codes. For providing a runtime environment, we will ask our subjects to use ssh to connect to one of our computers where they can edit the code from command line using nano or vim. This way we avoid having to setup runtime environments for each language on the subjects' computers.

We have implemented a Node.js server that will keep communicating with the codes as they debug them. The server will record each subject's start time and time taken for them to debug the code.

We will also provide a google form asking questions about the support and documentation provided by each language to evaluate the ease of debugging them.
