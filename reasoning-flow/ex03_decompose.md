## The problem

> Given an array of integers, find the second largest value without sorting
> the array, and state the time complexity.

Input used: `[7, 2, 9, 9, 4, 1]`

## Prompt given to the model

Solve this problem in four explicit parts, labeled:
1. **Restate** the problem in your own words.
2. **Steps** break the solution into an ordered list of steps before writing any code.
3. **Solve** walk through each step on the input `[7, 2, 9, 9, 4, 1]`, showing the running values of any variables you track.
4. **Final answer** state the second largest value and the time complexity, and confirm it follows from step 3.
Do not skip ahead to the answer before finishing steps 1-3.

## Model's output

**1. Restate:**
Find the second-highest distinct-position value in the array in a single
pass, without using a sort, and report how the runtime scales with input size.

**2. Steps:**
1. Initialize two trackers, `largest` and `secondLargest`, to negative infinity.
2. Walk through the array once, element by element.
3. For each element `x`:
   - If `x > largest`: move `largest` into `secondLargest`, then set `largest = x`.
   - Else if `x > secondLargest` and `x < largest`: set `secondLargest = x`.
4. After the loop, `secondLargest` holds the answer.
5. Report time complexity based on the loop structure.

**3. Solve** (`largest = -∞`, `secondLargest = -∞`):
| x | condition | largest | secondLargest |
|---|---|---|---|
| 7 | 7 > -∞ → new largest | 7 | -∞ |
| 2 | 2 < 7, 2 > -∞ → new second | 7 | 2 |
| 9 | 9 > 7 → new largest | 9 | 7 |
| 9 | 9 > 9? no. 9 > secondLargest(7) and 9 < largest(9)? no, 9 is not < 9 → no change | 9 | 7 |
| 4 | 4 < 9, 4 < 7 → no change | 9 | 7 |
| 1 | 1 < 9, 1 < 7 → no change | 9 | 7 |

**4. Final answer:**
Second largest = **7**. Time complexity = **O(n)**, one pass, constant extra
space (O(1)), and it follows directly from the trace in step 3 - the only
values that ever updated `secondLargest` were 2 and then 7.

Judgment
Step 1: fine, just tacked on "single pass" which nobody asked for.
Step 2: correct. The x < largest check is what saves it from the duplicate-max trap.
Step 3: every row checks out. On the second 9, it correctly doesn't overwrite secondLargest - that's the whole trap, and it passed.
Step 4: 7 and O(n) - both right, and pulled straight from step 3, not just asserted.