## 1. Chosen unstructured text type

Log lines 

## 2. Reusable template

Extract these fields from the log line and return JSON:
{ "level": ..., "timestamp": ..., "message": ... }
Log line: {INPUT}

## 3. Test runs

### Input A

```
2026-09-19 14:32:07 ERROR Failed to connect to database: connection timeout after 30s
```

**Output A:**
```json
{
  "level": "ERROR",
  "timestamp": "2026-09-19 14:32:07",
  "message": "Failed to connect to database: connection timeout after 30s"
}
```

### Input B

```
[WARN] 2026-09-18T09:15:42Z Cache miss ratio exceeded threshold (82%)
```

**Output B:**
```json
{
  "level": "WARN",
  "timestamp": "2026-09-18T09:15:42Z",
  "message": "Cache miss ratio exceeded threshold (82%)"
}
```

## 4. Consistency check

Yes, despite the two log lines using different formats (space-separated vs. bracketed level, different timestamp styles), both outputs came back as the same three-key JSON shape.

## Going further: missing-field rule

Updated template:

```
Extract these fields from the log line and return JSON:
{ "level": ..., "timestamp": ..., "message": ... }
If a field is missing from the log line, set its value to null instead of guessing.
Log line: {INPUT}
```

### Input C - missing timestamp

```
DEBUG Retrying request to payment-service (attempt 2/5)
```

**Output C:**
```json
{
  "level": "DEBUG",
  "timestamp": null,
  "message": "Retrying request to payment-service (attempt 2/5)"
}
```

The rule held: instead of inventing a timestamp, the model correctly returned `null` for the missing field.