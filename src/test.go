// Example usage
package main

import (
	"fmt"
	"github.com/jackcsheehan/re-go/rego"
)

func main() {
	t := `To be, or not to be, that is the question:
	Whether 'tis nobler in the mind to suffer
	The slings and arrows of outrageous fortune,
	Or to take arms against a sea of troubles
	And by opposing end them. To die—to sleep,
	No more; and by a sleep to say we end
	The heart-ache and the thousand natural shocks
	That flesh is heir to: 'tis a consummation
	Devoutly to be wish'd. To die, to sleep;
	To sleep, perchance to dream—ay, there's the rub:
	For in that sleep of death what dreams may come,
	When we have shuffled off this mortal coil,
	Must give us pause—there's the respect
	That makes calamity of so long life.
	For who would bear the whips and scorns of time,
	Th'oppressor's wrong, the proud man's contumely,
	The pangs of dispriz'd love, the law's delay,
	The insolence of office, and the spurns
	That patient merit of th'unworthy takes,
	When he himself might his quietus make
	With a bare bodkin? Who would fardels bear,
	To grunt and sweat under a weary life,`

	r := rego.Compile(`(?P<word>'\w+)(?P<fivechars>\w{5})`)
	//fmt.Println(r.IsMatch(t))
	//fmt.Println(r.FindAll(t))
	//fmt.Println(r.FindAllGroups(t))
	fmt.Println(r.FindAllNamedGroups(t))
}