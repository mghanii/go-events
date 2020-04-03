#### go-events

a simple event emitter/listener implementation.

### Example

```go
package main

import (
	"events"
	"fmt"
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
