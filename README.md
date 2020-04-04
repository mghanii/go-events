#### go-events
[![GoDoc](https://godoc.org/github.com/mghanii/go-events?status.svg)](https://pkg.go.dev/github.com/mghanii/go-events?tab=overview)

a simple event emitter/listener implementation.

#### Install

```bash
go get github.com/mghanii/go-events
```

#### Example

```go
package main

import (
	"fmt"

	"github.com/mghanii/go-events"
)

type Person struct {
	username string
	email    string
	emitter  events.Emitter
}

func (p *Person) UpdateEmail(email string) bool {
	if p.email != email {
		event := fmt.Sprintf("%v updated his email from '%v' to '%v'", p.username, p.email, email)
		p.email = email
		p.emitter.Emit("email-updated", event)
		return true
	}
	return false
}

func main() {
	p := Person{
		username: "Adam",
		email:    "first@first.com",
		emitter:  events.NewEmitter(),
	}

	ids := []string{
		p.emitter.AddListener("email-updated", func(id string, event string) {
			fmt.Printf("listener '%v': %v\n", id, event)
		}),
		p.emitter.AddListener("email-updated", func(id string, event string) {
			fmt.Printf("listener '%v': %v\n", id, event)
		})}

	p.UpdateEmail("second@second.com")

	// Removes listener
	p.emitter.RemoveListener(ids[0])

	p.UpdateEmail("third@third.com")

	fmt.Scanln()
}
```

output

```bash
2020/04/04 06:19:48 Listner added {id:'8065329c-a651-4210-a0b0-26ea1742e67e', event:'email-updated'}
2020/04/04 06:19:48 Listner added {id:'2ca7e0f1-7929-46e8-8e27-cebbaa0709f4', event:'email-updated'}
2020/04/04 06:19:48 Emiting event: 'email-updated'
2020/04/04 06:19:48 Listner removed {id:'8065329c-a651-4210-a0b0-26ea1742e67e', event:'email-updated'}
2020/04/04 06:19:48 Emiting event: 'email-updated'
listener '8065329c-a651-4210-a0b0-26ea1742e67e': Adam updated his email from 'first@first.com' to 'second@second.com'
listener '2ca7e0f1-7929-46e8-8e27-cebbaa0709f4': Adam updated his email from 'first@first.com' to 'second@second.com'
listener '2ca7e0f1-7929-46e8-8e27-cebbaa0709f4': Adam updated his email from 'second@second.com' to 'third@third.com'
```
