# NeverBounceApi-Go
[![Build Status](https://travis-ci.org/NeverBounce/NeverBounceApi-Go.svg?branch=master)](https://travis-ci.org/NeverBounce/NeverBounceApi-Go) [![Code Climate](https://codeclimate.com/github/NeverBounce/NeverBounceApi-Go/badges/gpa.svg)](https://codeclimate.com/github/NeverBounce/NeverBounceApi-Go)
## Start using it
1. Download and install it:

```sh
$ go get github.com/NeverBounce/NeverBounceApi-Go
```

2. Initiate NeverBounce package:

```go
func main() {
	neverBounce, err := neverBounce.New("secret_nvrbnc_golang")
    
    	if err != nil {
    		panic(err)
    	}
}
```

## API Examples

### Account
#### Info
```go
func main() {
	neverBounce, err := neverBounce.New("secret_nvrbnc_golang")
	if err != nil {
    	panic(err)
    }
    
    info, err := neverBounce.Info()
    if err != nil {
    	panic(err)
    }
}
```

### Single
#### Check
```go
func main() {
	neverBounce, err := neverBounce.New("secret_nvrbnc_golang")
    if err != nil {
    	panic(err)
    }
    
    singleCheckInfo, err := neverBounce.Single.Check("example@gmail.com", true, true, "")
    if err != nil {
    	panic(err)
    }
}
```