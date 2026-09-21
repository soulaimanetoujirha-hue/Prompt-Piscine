## 1. Prediction (before running anything)

I expected low temperature to give more predictable, conventional names, and high temperature to give more varied, creative, maybe weirder names

## 2. Outputs

### Temperature 0.2, Top-p 1.0
```
1. Standard & Idiomatic (Most Common)
mergeTwoLists
mergeSortedLists
mergeTwoSortedLists
mergeLists

2. Explicit & Highly Descriptive
mergeSortedLinkedLists
mergeTwoSortedLinkedLists
mergeSortedNodes
combineSortedLinkedLists

3. Short & Contextual (Clean / Functional Style)
merge
mergeSorted
combine
zipSorted

4. Instance Method (Object-Oriented Style)
mergeWith
mergeFrom
sortedMergeWith
mergeInPlace

Recommendation: mergeTwoSortedLists (standalone helper) or merge/mergeWith (inside a LinkedList class)
```

### Temperature 0.9, Top-p 1.0
```
1. Metaphorical & Visual
weaveSorted / weaveSortedLists
braidChains / braidLists
stitchSorted / stitchNodes
interlaceSorted
dovetailLists
zipperMerge

2. Functional & Elegant
collate / collateSorted
fuseOrdered
unifyOrdered
conjoinSorted
orderPreservingUnion
zipSorted / zipByOrder

3. Low-Level, Pointer-Centric & Mechanical
spliceSorted / spliceInPlace
relinkSorted
threadNodes
fuseChains
sewLists
destructiveMerge

4. Method / Receiver-Oriented
absorbSorted(...)
ingestSorted(...)
blendWith(...)
assimilate(...)
interthreadWith(...)

5. Architectural & Formal
coalesceOrderedSequences
consolidateSortedStreams
reconcileOrderedLists
blendSortedAscending
```

### Temperature 0.2, Top-p 0.5
```
mergeTwoLists
mergeSortedLists
mergeTwoSortedLists
mergeSortedLinkedLists
merge
mergeWith
```

### System instruction: "You are a terse senior engineer" (Temp 0.2, Top-p 0.5)
```
Don't overthink it. Don't encode types into function names if your language has a type system.

Free / static function: mergeSortedLists(a, b) or mergeTwoLists(a, b)
Method on LinkedList: listA.merge(listB) or listA.mergeWith(listB)
Mutating / in-place: listA.mergeInPlace(listB)

Pick: mergeSortedLists. Clear invariant, no fluff.
```

## 3. Did it match the prediction?

Low temperature (0.2) produced tight, conventional, industry-standard names (mergeTwoLists, mergeSortedLists) clustered around a handful of safe options — low variety, low risk. High temperature (0.9) produced wildly more creative and varied names (zipperMerge, braidChains, absorbSorted), spanning metaphorical, functional, and low-level naming styles — high variety, higher risk of impractical suggestions. Lowering top-p from 1.0 to 0.5 at temp 0.2 barely changed the output, since low temperature already narrows the distribution hard, leaving top-p little room to matter. Low temperature is preferable when you want predictable, reproducible, convention-following output (e.g. production code, docs); high temperature is better for brainstorming or exploring unconventional options.

## 4. System instruction effect

Setting the system instruction to "You are a terse senior engineer" cut out all the categorized lists and flavor commentary, replacing them with a single direct recommendation and a one-line justification — showing the system message controls tone/persona/response style independent of the user prompt itself.