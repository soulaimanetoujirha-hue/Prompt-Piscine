## step 1: 
unemployment rate in morocco 2026!

## step 2:
Reply with only a JSON object naming the tool and its arguments, like {"tool": ..., "args": {...}}.
{"tool": "get_statistic", "args": {"country": "morocco", "metric": "unemployment rate"}}
## step 3:
Fake tool result (made up): {"unemployment_rate": "9.5%", "year": 2026}

Prompt: The tool returned {"unemployment_rate": "9.5%", "year": 2026}. Answer the user's original question in plain language.

Output: Morocco's unemployment rate in 2026 is 9.5%.

##

The weakest link is Step 2 -> Step 3: the model can pick the right tool but pass wrong or malformed arguments (e.g. misspelled country, wrong metric name), so I'd check that the args match what the real tool function actually expects before calling it.