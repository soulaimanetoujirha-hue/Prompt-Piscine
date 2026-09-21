# ex03_transform.md

## 1. Vague prompt

Explain this code.

## 2. Patterned prompt (Instructional pattern)

Explain what this function does in 4 steps, for a teammate reading it for the first time:

[code]
**Code used for both runs:**
```python
def binary_search(arr, target):
    low, high = 0, len(arr) - 1
    while low <= high:
        mid = (low + high) // 2
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            low = mid + 1
        else:
            high = mid - 1
    return -1
```

## 3. Outputs

**Weak output (from "Explain this code."):**
This function performs a binary search on a sorted array `arr` to find `target`. It uses two pointers, `low` and `high`, and repeatedly checks the middle element, narrowing the search range until it finds the target or the range is empty, returning the index or -1 if not found.

**Patterned output (from the instructional prompt):**
1. Set up two pointers, `low` at the start of the array and `high` at the end, to define the current search range.
2. Loop while `low` is still less than or equal to `high`: calculate the middle index `mid` of the current range.
3. Compare `arr[mid]` to `target` — if it matches, return `mid` immediately; if `arr[mid]` is too small, move `low` past `mid`; if too large, move `high` before `mid`.
4. If the loop ends without finding a match (the range has shrunk to nothing), return `-1` to signal the target isn't in the array.

## 4. What improved

The weak output gave a single dense paragraph that's fine as a summary but hard to scan, while the patterned output forced a numbered, step-by-step breakdown that's easier for a teammate to follow line-by-line while reading the actual code for the first time.