# Exercise 2: Make the Model Trace Your Code

## The function

Iterative singly-linked-list reversal in Go

```go
type Node struct {
	Val  int
	Next *Node
}

func reverseLinkedList(head *Node) *Node {
	var prev *Node
	curr := head
	for curr != nil {
		nextNode := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextNode
	}
	return prev
}
```

Input for the trace: `1 -> 2 -> 3`

## Prompt given to the model

> Trace `reverse_linked_list` step by step for the input list `1 -> 2 -> 3`.
> At each loop iteration show the values of `prev`, `curr`, and `next_node`,
> and what `curr.next` gets reassigned to. Give the final returned list.

## Model's trace

**Initial:** `prev = None`, `curr = 1 -> 2 -> 3`

**Iteration 1:**
- `next_node = 2 -> 3`
- `curr.next = prev` → node 1 now points to `None`
- `prev = 1 -> None`
- `curr = 2 -> 3`

**Iteration 2:**
- `next_node = 3`
- `curr.next = prev` → node 2 now points to `1 -> None`
- `prev = 2 -> 1 -> None`
- `curr = 3`

**Iteration 3:**
- `next_node = None`
- `curr.next = prev` → node 3 now points to `2 -> 1 -> None`
- `prev = 3 -> 2 -> 1 -> None`
- `curr = None` → loop ends

**Return value:** `3 -> 2 -> 1`

## Actual result (ran the code)


```
[3, 2, 1]
```


## Verdict

**Correct — no divergence.** The model's trace matches the actual run at every
step and the final output. All three pointer reassignments (`prev`, `curr`,
`next_node`) line up exactly with what the interpreter did.