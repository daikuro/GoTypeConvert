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
