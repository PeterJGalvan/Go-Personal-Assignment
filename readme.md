# Cognizant Take Home Go Assignment (Entry Level)

This assignment is designed to test your ability to quickly learn basic Go concepts. Go is a procedural language with elements of both functional and object oriented languages sprinkled in.

We don't expect you to learn the entire language over night, just try your best and have fun with it!

## Assignment

### Computing Averages in a Matrix

</br>

Given a `m x n` matrix grid find the the average of each inner slice `n`, and return the values in a single dimensional slice.

You will be given the inputs to use in your function as a json file. It is up to you to write a function to parse the file, format the values into a `m x n` matrix of float32 values.

There will be a good.json and a bad.json, you need to write a function capable of properly handling both.

See the folder `algo` for the function signatures you must satisfy.

**Example 1:**

_Input:_ grid = [[10,23,56,23], [88,23,24,23,12]]
_Output:_ [28, 34]
_Explanation:_ The average of (10 + 23 + 56 + 23) = 28, and the average of (88 + 23 + 24 + 23 + 12) = 34

</br>

**Example 2:**

_Input:_ grid = [[45, 34.5], [], [23,  532]]
_Output:_ [39.75, 0, 277.5]

</br>

**Testing:**

In our backend we try to unit test every function as it makes sense. One of the patterns we use to unit test is called the go test table pattern. Dave Cheney, one of the go founders, has a blog post on how to write test tables. I included it in the developer resources below as well as a youtube tutorial referencing the same thing.

Once you solve the solution I would like you to try to write a test table that covers various inputs to your functions.

## Developer Resources

- Intro to Go: https://go.dev/tour/welcome/1
- A good book: https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing/dp/0134190440
- How to write test tables: https://dave.cheney.net/2019/05/07/prefer-table-driven-tests
- A nice video on go unit tests: https://www.youtube.com/watch?v=0Ov3JAYDyhQ

## Submission

Please zip this folder and send it to farah.fatany@equifax.com and brice.aldrich@equifax.com. Alternatively, you can commit the code to your public github and send the link.
