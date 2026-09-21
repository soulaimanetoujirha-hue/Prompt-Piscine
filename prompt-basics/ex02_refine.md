## Vague ---Prompt: "Review my code"
func ListSize(l *List) int {
	countre := 0
	current := l.Head
	for current != nil {
		current = current.Next
		countre += 1
	}
	return countre
}
## Model output: 
This function looks fine — it counts the number of nodes in a linked list by walking from the head to the end. Consider renaming countre to count for clarity
-------------------------------------------------------------------
## Better ---Prompt: "Review this Go function for correctness and edge cases:

## Model output:
The logic is correct for a standard singly-linked list — it returns the right count. Edge cases: if l.Head is nil, the loop body never runs and it correctly returns 0. However, if l itself is nil, l.Head will panic. Also, if the list contains a cycle, this function will loop forever
-----------------------------------------------------------------------------
## Precise ---Prompt: "Review this Go function. List any off-by-one or nil-pointer bugs, and for each, show the input that triggers it:
## Model output:
Bug	Triggering                          || input
Nil-pointer dereference on              || l = nil
l.Head	                                ||
Infinite loop (no cycle guard)          ||	A list where the last node's Next points back to an earlier node instead of nil
No off-by-one bug found	                ||  the increment-after-advance pattern still counts  
                                        ||  correctly starting from the head
-----------------------------------------------------------------------------------------------
## what specifically improved the answer
The vague one just got a "looks fine, maybe rename a variable" — no scope, so the model just went with whatever it felt like checking. Once I told it to focus on correctness and edge cases, it actually started thinking about nil and cycles instead of skimming the surface. Then locking it into a table format with "show the triggering input" made it stop rambling and actually give me something I could turn straight into a test case.