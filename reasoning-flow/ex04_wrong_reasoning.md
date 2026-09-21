## 1. Was every step actually correct?

For the specific inputs used in ex02 and ex03 - yes, every step checked out. But that's exactly the trap: "correct on the input I picked" isn't the same as "correct in general.".
so "looked right" was doing a lot of unearned work.

## 2. One step I'd double-check

Ex03, Step 2 (the algorithm design): the else if x > secondLargest and x < largest guard. It's correct for arrays that have two distinct values, but it was never checked against an array where no valid second-largest exists (all elements equal). I'd double-check this because it's the kind of gap that doesn't show up unless you go looking for it - the trace on a normal input will always look clean.

## 3. What this teaches me

Confident, step-by-step reasoning can be internally consistent and still reach a wrong or meaningless answer - the steps following logically from each other isn't proof the steps are right. Trust the trace only as far as you've actually verified the edge cases, not just the happy path.

