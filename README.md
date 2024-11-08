# Demonstration Project for Lean TECHniques

## Summary
From the help message:
```
currency2change is a quick program for running the cashier's algorithm on arbitrary amounts of money.


Usage:

        currency2change <dollar/cent amount>

        i.e.
        currency2change $123.45
        currency2change $123.4
        currency2change $123.
        currency2change $123
        currency2change 123
        currency2change .2
        curremcy2change $999,999,999.99


Example output:

        > ./currency2change $25,709.82
        257 - $100 bill
        1 - $5 bill
        4 - $1 bill
        3 - Quarter
        1 - Nickle
        2 - Penny


NOTE:
        In most shells, dollar signs must be escaped, or they will cause
        problems.
```


## Building
Being Go, building currency2change is rather simple. Assuming Go and Git are installed:

```
git clone github.com/11Spades/lt-cashiers
cd lt-cashiers
go build currency2change.go
```
