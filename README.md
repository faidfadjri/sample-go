### Go Routine

[![go Routine](https://repository-images.githubusercontent.com/257380389/e7d41600-9506-11ea-8837-fefe3a4f2b51 "go Routine")](https://repository-images.githubusercontent.com/257380389/e7d41600-9506-11ea-8837-fefe3a4f2b51 "go Routine")

Go Routine is a small **thread**.
Actually is not a thread is a bunch of task which run inside a thread.
is made by go and managed by Go Scheduler.

> Go Routine is run concurrent is (not parallel)

#### Difference between Concurrency & Parallel

Parallelism involves running two tasks at the same time. For example, when you and your friend are both eating your food simultaneously, that's parallelism.

Concurrency, on the other hand, doesn't necessarily mean running tasks alternately. It means that two tasks are being executed independently, but not necessarily simultaneously. IFor example, when you eat your food and occasionally talk with your friend, it's concurrent. However, you can't eat and talk at the exact same moment, or, as you humorously pointed out, you might end up throwing food in your friend's face.

#### Key Point

1. by default go will sort the go routine according the time execution.
   so the less time take will running and the more time take are eliminated (not ignored) but like running in a background
2. by default in go routine you can't call the function with have a return value.
3. but if you want to get the return value by using go routine you can use the **Channel**

#### What is Channel ?

Channel is a way how you can communicate the go routine synchronously, so you can wait until all the go routine are executed. (e.g Async Await on javascript).

1. Channel has it's own go routine which have a function to send / retrieve value from channel
2. By default, Channel only can save one data value
3. But you can retrive value from multiple go routine
4. When you create a channel if you not call the channel to get the return value the code will blocked inside the go routined which you've been created.
5. if you created a channel but didn't send anything to the channel when you try to retrive it it will run a Deadlock Error.
6. Channel must be "closed" if the channel has executed or has been used to avoid memory leak
