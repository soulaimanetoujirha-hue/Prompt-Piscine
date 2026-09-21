## Part A: Build your own analysis first (no AI)

# 1. Ethical Dilemma in AI: Hiring Algorithms
1. Concrete example

Amazon built an experimental AI recruiting tool (2014-2017) to screen resumes. It was trained on 10 years of past resumes, most of which came from men, since the tech industry was male-dominated. The model learned to penalize resumes containing the word "women's" (e.g. "women's chess club captain") and downgraded graduates of all-women's colleges. Amazon scrapped the tool once this was discovered.

2. Where the bias originated

Mainly the training data: 10 years of mostly-male resumes taught the model that "male" was the statistical pattern of a successful hire. Task framing made it worse - the goal was "find more people like our past hires" instead of "find people who can do the job," so the model optimized for cloning the existing workforce instead of for actual merit.

3. Who is harmed, and how

Qualified women applicants are harmed directly: their resumes get auto-downgraded or filtered out before a human ever sees them, simply for wording tied to gender (e.g. "women's"). The concrete consequence is lost interview opportunities and lost job offers, for reasons that have nothing to do with their actual skills. It also harms the company by reinforcing existing bias and prompting legal/reputational risk if used in production.
------------------------------------------------------------------------
`[Solutions:]`

1. Bias Detection & Mitigation:

Regularly audit models to ensure diverse, representative training data.
Correct detected biases promptly to minimize discrimination.

2. Human Oversight:

Recruiters and hiring managers should review and validate automated decisions.
Provide a human perspective for evaluating underrepresented applicants.

3. Transparent Criteria:

Ensure candidates understand the criteria used by AI for selection.
Clearly communicate attributes being assessed.

4. Regular Feedback Loops:

Collect feedback from applicants and employees on AI recruitment tools.
Use feedback data to refine algorithms and implementation.

## Part B: Then use AI to stress-test your analysis

Causes you didn't consider:
---------------------------
Deployment/proxy-variable problem: even if you scrub "women's" and college names from the input, the model can still pick up correlated signals — extracurriculars, phrasing style, even resume formatting conventions that differ by gender. Removing obvious gendered words doesn't fix a model that learned gender as a pattern, not a keyword. This is a deployment/mitigation-design flaw, not just a training-data or framing flaw.
Feedback loop / lack of monitoring: Amazon reportedly kept tweaking the model when new bias patterns surfaced, and it kept finding new proxies. That's a deployment failure too — no one caught it until it was already influencing real screening decisions in some regions.

Affected groups you overlooked:
-------------------------------
Amazon itself / the hiring pipeline broadly: it's not just individual women harmed — the whole applicant pool skews narrower, so Amazon loses out on talent and diversity of thought, which is a business harm, not just an individual one.
Men from non-traditional backgrounds: a model trained to clone "successful past hires" also penalizes men who don't fit the existing profile (e.g. career-changers, non-CS majors) — it's not a purely binary "hurts women only" story, it's "hurts anyone who deviates from the historical template."
Downstream trust in AI hiring tools generally: candidates and other companies lose trust in automated screening as a category, which affects people applying to other companies using similar tools, not just Amazon applicants.

Do you agree?
-------------
the proxy-variable point is the one I'd weigh most: your original analysis assumed cleaning the training data (or reframing the task) fully fixes the problem, and the proxy-variable issue is exactly why that assumption doesn't fully hold.

## Part C: Reflection

My contribution was the concrete example itself and the judgment call on what caused it - the AI could summarize causes generically, but I had to decide which cause was primary and defend it. Yes, I think I'd catch this in my own projects now, because I know to check what the training data actually represents before trusting a model's output - for example, if I ever built a resume screener, I'd ask what population the training resumes came from before trusting its rankings.