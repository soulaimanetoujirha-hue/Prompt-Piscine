## Problem :
Three shelves each start with the same number of books, x.

Monday: Shelf 1 gets a shipment of 20 books. Shelf 2 gets a shipment of 5 books. Shelf 3 sells 5 books, then gets a shipment of 15 books.
Tuesday: Shelf 2 sells 15 books. Shelf 1 and Shelf 3 have no Tuesday activity.
After Tuesday's sale, Shelf 2 has exactly half as many books as Shelf 1, and Shelf 3 has exactly 20 more books than Shelf 2.

How many books did each shelf start with (x), and how many books does each shelf have now?
Answer with just the result!!

## 1:
x = 40

Shelf 1: 60 | Shelf 2: 30 | Shelf 3: 50

## 2: 
Step 1: Track each shelf after Monday's activity

Shelf 1: x + 20
Shelf 2: x + 5
Shelf 3: x − 5 + 15 = x + 10

Step 2: Apply Tuesday's sale

Shelf 1: x + 20 (no change)
Shelf 2: (x + 5) − 15 = x − 10
Shelf 3: x + 10 (no change)

Step 3: Use the condition "Shelf 2 = half of Shelf 1"

x − 10 = (x + 20) / 2
2x − 20 = x + 20
x = 40

Step 4: Check "Shelf 3 = Shelf 2 + 20"

Shelf 2 = 40 − 10 = 30
Shelf 3 = 40 + 10 = 50
30 + 20 = 50 ✓ (confirms consistency)

Result:

Starting number, x = 40
Shelf 1 now: 60
Shelf 2 now: 30
Shelf 3 now: 50

## 3:
Same answer, only the presentation changed. This problem was simple enough (one linear equation, few variables) that I could solve it correctly in a single pass; showing steps just made the reasoning visible, it didn't alter the result.

CoT (step-by-step reasoning) actually matters most on problems where one wrong move early on messes up everything after it, like word problems with several conditions to juggle, puzzles where you have to track different cases, "if A then B then C" type reasoning, or tracing code through multiple changing states. Basically: the more stuff the model has to keep track of in its head at once, the more likely it messes something up if it rushes. Writing out the steps forces it to slow down and check each one before moving on, so mistakes get caught instead of snowballing.