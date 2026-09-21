# Exercise 4: Accountability

**Scenario:** I ship a feature that uses an LLM. It gives a user a confidently wrong answer, and that causes real harm.

## 1. Who is responsible?

I think responsibility is shared, but it starts with me as the developer, not the model provider and not the user.

- **Me (the developer):** I chose to put this model in front of users for this specific task, without necessarily adding checks, limits, or warnings. I decided what the feature does and how much trust a user is invited to place in it. That decision is mine, so the consequences of that decision are mine too.
- **My company:** shares responsibility for what it approved to ship, and for whether it gave me the time/resources to build in safety checks, or pushed to ship fast without them.
- **The model provider:** responsible for being honest about the model's known limits (e.g. that it can be confidently wrong), but not responsible for how I chose to use it in my product.
- **The user:** only responsible if they ignored a clear warning I gave them. If I didn't warn them the answer could be wrong, that's on me, not them.

"The model did it" is not acceptable from an engineer because the model didn't choose to ship itself into a product - I did. An engineer is responsible for knowing the tool's limits and designing around them, the same way a civil engineer can't blame the concrete for a bridge collapse if they used it outside its rated limits. Using an LLM means accepting that it will sometimes be confidently wrong, and building the feature as if that will happen, not hoping it won't.

## 2. Two things I would put in place before shipping

1. **A visible confidence/limits disclaimer at the point of use** - telling the user this is an AI answer that can be wrong, especially for anything with real consequences (medical, legal, financial, security). This stops users from placing more trust in the answer than it deserves.
2. **A human-in-the-loop or verification step for high-stakes answers** - for example, routing anything flagged as medical/legal/financial/security-related to a human reviewer, or requiring the AI to cite a verifiable source before the answer is shown, instead of letting a raw LLM answer go straight to the user unchecked.

## 3. Going further: for and against each party

**Me (the developer)**
- For: I made the design choices - what the feature does, what checks exist, what warnings are shown.
- Against: I may not control company deadlines, budget for safety work, or the underlying model's behavior.

**My company**
- For: it approved the launch and set the priorities (speed vs. safety) I had to work within.
- Against: it may have trusted my technical judgment on how safe the feature was, in good faith.

**The model provider**
- For: they built and trained the model, so its tendency to be confidently wrong originates with them.
- Against: they typically publish known limitations and terms of use - it's on the product built on top to design around those, not on the provider to anticipate every possible use case.

**The user**
- For: if I gave a clear warning and they ignored it, some responsibility shifts to them.
- Against: most users don't have the background to know when an AI answer might be wrong, especially if the product doesn't say so.