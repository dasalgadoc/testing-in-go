# 🧟‍♀️ Mutant testing

Mutation testing is a technique to evaluate the quality of a test suite. It consists of modifying the production code and checking if the tests fail.

This method is predicated on the notion that complete coverage does not equate to complete functionality testing.

Some external tools are available to help us with this task. One of them is [Mutesting](https://github.com/zimmski/go-mutesting)

A naive example of mutation testing is to change the `Add` function to always return `0` and check if the tests fail.
