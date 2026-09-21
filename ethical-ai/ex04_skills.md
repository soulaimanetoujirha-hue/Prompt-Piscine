Skill	rate
Problem decomposition	4
Systems thinking	4
Critical evaluation	3
Debugging mindset	3
Conceptual understanding	4
---------- Critical evaluation && Debugging mindset ---------
One habit that prepares me is coding and learning on paper alongside the laptop. It may seem like a waste of time at first, but it's actually a great way to keep the fundamentals fresh in memory.
--------------Problem decomposition-------------
When fixing RemoveValue, I eventually separated the head-node case from the rest-of-list case instead of trying to handle both in one loop.
--------------Systems thinking---------------
I didn't immediately see how changing curr inside the loop was disconnected from the actual Next pointers on the list — I had to trace it by hand to understand why reassigning a local variable doesn't mutate the structure.
--------------Conceptual understanding----------------
I understood pointers well enough to eventually rewire curr.Next, but I initially confused reassigning a local pointer variable (curr = curr.Next) with mutating the list itself — a gap in why Go pointers work the way they do, not just how to use them.