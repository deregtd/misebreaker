# Mise breaker example

[Discussion on mise repo](https://github.com/jdx/mise/discussions/4489)

## Run

Make sure you don't have node globally installed/set via mise.  The test app will check for this.

Out of the box, if you run:

```
mise run run
```

You should get the following error:
```
➜  misebreaker git:(main) mise run run
[run] $ go run runner.go 
Got error on node version #1

2025/02/22 15:25:43 exec: "node": executable file not found in $PATH
exit status 1
[run] ERROR task failed
```

If you go into the `runner.go` file and change line 16 from:

```
	pool := pond.NewPool(numDirs)
```

to

```
	pool := pond.NewPool(1)
```

... then that will remove the concurrency and it will run the two tasks sequentially, and it should pass just fine:

```
➜  misebreaker git:(main) ✗ mise run run
[run] $ go run runner.go 
node: v22.7.0

node: v22.7.0
```
