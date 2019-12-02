[![CircleCI](https://circleci.com/gh/daikuro/GoTypeConvert.svg?style=svg)](https://circleci.com/gh/daikuro/GoTypeConvert)
[![codecov](https://codecov.io/gh/daikuro/GoTypeConvert/branch/master/graph/badge.svg)](https://codecov.io/gh/daikuro/GoTypeConvert)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# GoTypeConvert


## Install

```
go get github.com/daikuro/GoTypeConvert
```

or 

```
dep init
```

or

```
dep ensure
```

## Update

```
go get -u github.com/daikuro/GoTypeConvert
```

or

```
dep ensure -update github.com/daikuro/GoTypeConvert
```

## Usage

```
import (
  typeconv "github.com/daikuro/GoTypeConvert"
)

func main() {
	var v interface{}
	v = "test"
	s := typeconv.ToString(v).A
	fmt.Println(s)
}
```
