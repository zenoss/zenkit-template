{{ $pkg := print ((print (env "GOPATH") "/src/") | trimPrefix (env "PWD")) "/" Name -}}
package main

import "{{$pkg}}/{{Name}}/cmd"

func main() {
	cmd.Execute()
}
