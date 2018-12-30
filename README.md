# pow
replace golang math.Pow to optimal

*benchmark*

```
go test -bench=. -count=5 > bench.txt
benchstat bench.txt 
```
```
name          time/op
/math.Pow2-4   357ns ± 0%
/pow.E2-4     5.69ns ± 0%

/math.Pow3-4   359ns ± 1%
/pow.E3-4     8.86ns ± 1%
```

So, that package is more optimal at more 30 times.
