# Exercise 1: Generate, Then Add Constraints

## Prompt 1 (simple)
> Write a short story about overcoming a hard bug.

**Output 1:**

Maya stared at the screen for the fourth hour straight. The login form worked for everyone except users named "O'Brien" — the apostrophe was breaking the SQL query somewhere deep in the stack. She'd tried three fixes already, each one patching a symptom instead of the cause.

At 2am she finally found it: a raw string concatenation buried in a legacy helper function nobody had touched in two years. One parameterized query later, it worked. She didn't celebrate. She just saved, closed the laptop, and went to sleep.

## Prompt 2 (constrained: under 150 words, first person, hopeful tone)
> Write a short story about overcoming a hard bug, in under 150 words, first person, hopeful tone.

**Output 2:**

I'd been staring at the same failing test for three hours, convinced the framework itself was broken. Every fix I tried made it worse. I almost gave up and filed it as "known issue, revisit later."

Instead I did something simple — I read the error message slowly, word by word, instead of skimming it like I had ten times before. There it was: a typo in an environment variable name I'd copy-pasted early on.

One character. Three hours. I fixed it in five seconds.

I used to think debugging was about being smart. Now I think it's about being patient enough to actually look. Every bug I've ever "solved by luck" was really just me finally slowing down. That's the part that keeps me going back to the keyboard.

## What changed
The constrained version is punchier and more personal — first-person and a word limit forced it to cut filler and land on one clear insight instead of just narrating events.
