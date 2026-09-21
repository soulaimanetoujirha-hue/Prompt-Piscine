what the bug was.
-----------------
The RemoveValue() function didn't check the first element of the list — it jumped directly to the second element, so removing a value at the head never worked
------------------------------
did AI point out anything useful you had not thought of?
--------------------------------------------------------
If the list has duplicate values, e.g. 1 -> 2 -> 2 -> 4, removing 2 will only remove the first match it finds.