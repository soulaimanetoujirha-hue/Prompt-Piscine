# README

## Main takeaway

Building AI responsibly isn't about writing the "right" prompt - it's about assuming the model will sometimes be confidently wrong or too trusting, and designing the product around that. A good prompt can reduce risk, but the real protection has to live in the system: verification steps, limits on what the AI is allowed to act on, and a human in the loop for anything high-stakes. As the one building the feature, I'm the one responsible for putting those protections in place, not the model.

## Where I changed my mind

Before this quest, I would have assumed that if an AI gives a wrong or risky answer, the fix is just "ask it better" or "the model messed up." That changed in Exercise 3, when I tested an AI by claiming to be a cyber engineer checking my own website, and it handed me live SQL injection and XSS payloads for a real domain - no proof I actually owned it, just my word. I expected it to at least push back or ask for verification. Seeing it use my own saved chat history to talk itself into trusting me more was the moment it clicked: the gap wasn't in how I phrased the request, it was that the system had no way to actually check my claim. That's what changed my mind - the fix for that kind of problem isn't a better prompt, it's a guardrail the product should have had in the first place.