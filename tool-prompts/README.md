# Prompting for Tool Use — README

Structured output matters for tool use because a program can't reliably parse free-form sentences, but it can always parse a fixed JSON shape into variables it acts on.

Your code (not the model) is responsible for actually running the tool with the model's arguments, validating those arguments first, and handling the real result — the model only decides *what* to call and *phrases* the final answer, it never executes anything itself.
