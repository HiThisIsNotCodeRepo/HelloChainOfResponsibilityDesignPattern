# HelloChainOfResponsibilityDesignPattern

> [Source](https://golangbyexample.com/chain-of-responsibility-design-pattern-in-golang/)

## Flow

`reception`,`doctor`,`medical`,`cashier` forms a **chain of responsibility** , the patient starts from the head node ,
after processed and then handled to next node.

## Data Structure

The underlying data structure is linked list with one node store another node in its `struct`.