## 1. Weak prompt

Summarize this.

## 2. Improved prompt

Summarize the following function in 3 bullet points: what it does, its inputs, and one edge case to watch. [code]

## 3. Test runs

### Input A - recursive function

```python
def factorial(n):
    if n <= 1:
        return 1
    return n * factorial(n - 1)
```

**Output A:**
- What it does: Recursively computes the factorial of `n` (n! = n × (n-1) × ... × 1) by multiplying `n` by the result of calling itself on `n - 1`.
- Inputs: A single integer `n`.
- Edge case: Negative input causes infinite recursion (and eventual stack overflow), since `n <= 1` is never true for negative numbers — no negative check exists.

### Input B - commit message

fix: handle negative input in factorial() to avoid infinite recursion

**Output B:**
- What it does: Fixes the `factorial()` function so it no longer recurses infinitely.
- Inputs: The fix specifically targets negative integer inputs.
- Edge case: Negative input was the original bug — the commit adds a guard so negative values are handled instead of triggering unbounded recursion.

## 4. Consistency check

Yes. both outputs came back in the same 3-bullet shape (what it does / inputs / edge case), even though one input was code and the other was a one-line commit message, which is exactly the repeatable structure the weak prompt failed to guarantee.