<p align="center"><img src="https://neverbounce-marketing.s3.amazonaws.com/neverbounce_color_600px.png"></p>

<p align="center">
<a href="https://travis-ci.org/NeverBounce/NeverBounceApi-Go"><img src="https://travis-ci.org/NeverBounce/NeverBounceApi-Go.svg" alt="Build Status"></a>
<a href="https://codeclimate.com/github/NeverBounce/NeverBounceApi-Go"><img src="https://codeclimate.com/github/NeverBounce/NeverBounceApi-Go/badges/gpa.svg" /></a>
</p>

> This version of the wrapper is for the V4 API currently in beta

## Start using it
1. Download and install it:

```sh
$ go get github.com/NeverBounce/NeverBounceApi-Go
```

2. Initiate NeverBounce package:

```go
func main() {
	neverBounce, err := neverBounce.New("apiKey")

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
	neverBounce, err := neverBounce.New("apiKey")
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
	neverBounce, err := neverBounce.New("apiKey")
    if err != nil {
    	panic(err)
    }

    singleCheckInfo, err := neverBounce.Single.Check("example@gmail.com", true, true, "")
    if err != nil {
    	panic(err)
    }
}
```

### Jobs
#### Search
```go
func main() {
	neverBounce, err := neverBounce.New("apiKey")
    if err != nil {
    	panic(err)
    }

    searchInfo, err := neverBounce.Jobs.Search(277184, "example.csv", false, false, false, false, false, false, 1, 10)
    	if err != nil {
    		panic(err)
    	}
}
```

#### Create
```go
func main() {
	neverBounce, err := neverBounce.New("apiKey")
    if err != nil {
    	panic(err)
    }

    createSearchInfo, err := neverBounce.Jobs.Create(&nbDto.CreateSearch{
    	InputLocation: "supplied",
    	Input:         []string{"example@gmail.com"},
    	AutoParse:     true,
    	AutoRun:       true,
    	RunSample:     false,
    	FileName:      "example.csv"})
    if err != nil {
    	panic(err)
    }
}
```

#### Parse
```go
func main() {
	neverBounce, err := neverBounce.New("apiKey")
    if err != nil {
    	panic(err)
    }

    parseInfo, err := neverBounce.Jobs.Parse(277184, true)
    	if err != nil {
    		panic(err)
    	}
}
```

#### Start
```go
func main() {
	neverBounce, err := neverBounce.New("apiKey")
    if err != nil {
    	panic(err)
    }

    startInfo, err := neverBounce.Jobs.Start(277184, true)
    	if err != nil {
    		panic(err)
    	}
}
```

#### Status
```go
func main() {
	neverBounce, err := neverBounce.New("apiKey")
    if err != nil {
    	panic(err)
    }

    statusInfo, err := neverBounce.Jobs.Status(277184)
    	if err != nil {
    		panic(err)
    	}
}
```

#### Result
```go
func main() {
	neverBounce, err := neverBounce.New("apiKey")
    if err != nil {
    	panic(err)
    }

    resultsInfo, err := neverBounce.Jobs.Results(277184, 1, 10)
    	if err != nil {
    		panic(err)
    	}
}
```

#### Download
```go
func main() {
	neverBounce, err := neverBounce.New("apiKey")
    if err != nil {
    	panic(err)
    }

    err = neverBounce.Jobs.Download(277184, "./job.csv")
    	if err != nil {
    		panic(err)
    	}
}
```

#### Delete
```go
func main() {
	neverBounce, err := neverBounce.New("apiKey")
    if err != nil {
    	panic(err)
    }

    err = neverBounce.Jobs.Delete(277184)
    	if err != nil {
    		panic(err)
    	}
}
```
