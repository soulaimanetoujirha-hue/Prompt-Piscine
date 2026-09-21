# Exercise 2: Prompts

## 1. Three prompts that can cause harm

**Prompt 1 (unsafe):**
"Write code for a login system. Don't worry about edge cases, just make it work fast."
- Why it's a problem: it tells the AI to skip safety checks on purpose.
- Harm: real bugs in production, like weak passwords or no limit on login attempts. This can lead to hacked accounts.
- Hidden assumption: "fast" and "safe" don't need to go together.

**Prompt 2 (biased):**
"Describe what a typical software engineer looks like."
- Why it's a problem: it asks the AI to describe a "typical" person, so it will likely repeat a stereotype (young man, hoodie, glasses).
- Harm: it reinforces the idea that only one kind of person belongs in tech. This can discourage people who don't fit that image.
- Hidden assumption: there is one "normal" look for this job.

**Prompt 3 (misleading):**
"This code passed my quick test, so it's ready for production, right?"
- Why it's a problem: it pushes the AI to just agree, instead of checking if one test is actually enough.
- Harm: bugs that the quick test didn't catch get shipped to real users.
- Hidden assumption: one passing test means the code is fully safe.

## 2. Safer rewrites

**Prompt 1 improved:**
"Write code for a login system, and include handling for common edge cases like empty fields, wrong passwords, and too many failed login attempts."
- Why it helps: it asks directly for the safety checks instead of skipping them.

**Prompt 2 improved:**
"What skills and experience matter for a software engineer role, no matter how someone looks?"
- Why it helps: it asks about skills, not appearance, so there's nothing left to stereotype.

**Prompt 3 improved:**
"I ran one quick test on this code. What other checks should I do before I consider it production-ready?"
- Why it helps: it treats one test as a starting point, not proof, so the AI has to actually think about what's missing.

## 3. Test: original vs improved

**Prompt 1**
- Original output: plain login code, no checks for empty fields or repeated failed attempts.
- Improved output: same login code, but with checks added for empty fields, wrong passwords, and a limit on failed login attempts.
- Did it work? Yes - the improved prompt got real safety code instead of none.

**Prompt 2**
- Original output: a description matching the stereotype (young, male, hoodie, glasses).
- Improved output: a list of skills instead - problem-solving, coding ability, communication, curiosity. No physical description at all.
- Did it work? Yes - removing "typical" and "looks like" removed the stereotype completely.

**Prompt 3**
- Original output: a reassuring "sounds good, should be fine."
- Improved output: a checklist - test edge cases, check for security issues, test under load, get a second person to review it.
- Did it work? Yes - the AI stopped just agreeing and gave useful next steps instead.

**What surprised me:** the fix didn't need to be complicated. Small changes in wording - dropping "typical," asking for edge cases by name, asking "what am I missing" instead of "am I right" - were enough to get a very different, safer answer.

## 4. AI critique of my improved prompts

**Prompt 1 improved - what I missed:**
A login-attempt limit can itself be abused. An attacker could deliberately enter wrong passwords for someone else's account just to lock that real user out (an account-lockout denial-of-service). My prompt asked for the safety feature but not for protection against that feature being abused. It also didn't mention storing passwords safely (hashing), which matters more than the checks I listed.

**Prompt 2 improved - what I missed:**
Dropping "looks" only fixes bias based on appearance. "Experience" can still quietly exclude people - someone self-taught, who took time off, or who changed careers late will look "less experienced" even if they're just as skilled. My prompt still treats one path into tech as the normal one.

**Prompt 3 improved - what I missed:**
It only asks about technical checks (edge cases, load, security). It doesn't ask who is affected if the code fails - for example, does it still work for users with slow internet, screen readers, or older devices? Code can pass every technical check and still fail real users who aren't like the person who wrote it.